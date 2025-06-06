package discv5

import (
	"context"
	"crypto/ecdsa"
	crand "crypto/rand"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"time"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/p2p/enode"
	log "github.com/sirupsen/logrus"

	"github.com/dennis-tra/nebula-crawler/core"
	"github.com/dennis-tra/nebula-crawler/db"
	"github.com/dennis-tra/nebula-crawler/utils"
)

type DialDriverConfig struct {
	Version string
}

type DialDriver struct {
	cfg         *DialDriverConfig
	dbc         db.Client
	peerstore   *enode.DB
	taskQueue   chan PeerInfo
	start       chan struct{}
	shutdown    chan struct{}
	done        chan struct{}
	dialerCount int
	writerCount int
}

var _ core.Driver[PeerInfo, core.DialResult[PeerInfo]] = (*DialDriver)(nil)

func NewDialDriver(dbc db.Client, cfg *DialDriverConfig) (*DialDriver, error) {
	peerstore, err := enode.OpenDB("") // in memory db
	if err != nil {
		return nil, fmt.Errorf("open in-memory peerstore: %w", err)
	}

	d := &DialDriver{
		cfg:       cfg,
		dbc:       dbc,
		peerstore: peerstore,
		taskQueue: make(chan PeerInfo),
		start:     make(chan struct{}),
		shutdown:  make(chan struct{}),
		done:      make(chan struct{}),
	}

	go d.monitorDatabase()

	return d, nil
}

func (d *DialDriver) NewWorker() (core.Worker[PeerInfo, core.DialResult[PeerInfo]], error) {
	// If I'm not using the below elliptic curve, some Ethereum clients will reject communication
	priv, err := ecdsa.GenerateKey(ethcrypto.S256(), crand.Reader)
	if err != nil {
		return nil, fmt.Errorf("new ethereum ecdsa key: %w", err)
	}

	ethNode := enode.NewLocalNode(d.peerstore, priv)

	conn, err := net.ListenUDP("udp", nil)
	if err != nil {
		return nil, fmt.Errorf("listen on udp port: %w", err)
	}

	discv5Cfg := discover.Config{
		PrivateKey:   priv,
		ValidSchemes: enode.ValidSchemes,
	}

	listener, err := discover.ListenV5(conn, ethNode, discv5Cfg)
	if err != nil {
		return nil, fmt.Errorf("listen discv5: %w", err)
	}

	dialer := &Dialer{
		id:       fmt.Sprintf("dialer-%02d", d.dialerCount),
		listener: listener,
	}

	d.dialerCount += 1

	return dialer, nil
}

func (d *DialDriver) NewWriter() (core.Worker[core.DialResult[PeerInfo], core.WriteResult], error) {
	id := fmt.Sprintf("writer-%02d", d.writerCount)
	w := core.NewDialWriter[PeerInfo](id, d.dbc)
	d.writerCount += 1
	return w, nil
}

func (d *DialDriver) Tasks() <-chan PeerInfo {
	close(d.start)
	return d.taskQueue
}

func (d *DialDriver) Close() {
	close(d.shutdown)
	<-d.done
	close(d.taskQueue)
}

// monitorDatabase checks every 10 seconds if there are peer sessions that are due to be renewed.
func (d *DialDriver) monitorDatabase() {
	defer close(d.done)

	select {
	case <-d.start:
	case <-d.shutdown:
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-d.shutdown
		cancel()
	}()

	for {
		log.Infof("Looking for peers to probe...")
		addrInfos, err := d.dbc.SelectPeersToProbe(ctx)
		if errors.Is(err, sql.ErrNoRows) || len(addrInfos) == 0 {
			log.Infoln("No peers due to be probed")
		} else if err != nil {
			log.WithError(err).Warnln("Could not fetch sessions")
			goto TICK
		}

		for _, addrInfo := range addrInfos {
			logEntry := log.WithField("peerID", addrInfo.ID.ShortString())

			// use custom identity scheme to not check the signature.
			node, err := utils.ToEnode(addrInfo.ID, addrInfo.Addrs)
			if err != nil {
				logEntry.WithError(err).Warnln("Could not construct new enode.Node struct")
				continue
			}

			pi := PeerInfo{
				Node:   node,
				peerID: addrInfo.ID,
				maddrs: addrInfo.Addrs,
			}

			select {
			case d.taskQueue <- pi:
				continue
			case <-ctx.Done():
				// fallthrough
			}
			break
		}

	TICK:
		select {
		case <-time.Tick(10 * time.Second):
			continue
		case <-ctx.Done():
			return
		}
	}
}
