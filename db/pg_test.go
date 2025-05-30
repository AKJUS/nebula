package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	tnoop "go.opentelemetry.io/otel/trace/noop"

	pgmodels "github.com/dennis-tra/nebula-crawler/db/models/pg"
	lp2ptest "github.com/libp2p/go-libp2p/core/test"
	"github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	mnoop "go.opentelemetry.io/otel/metric/noop"
)

func clearDatabase(ctx context.Context, db *sql.DB) error {
	if _, err := pgmodels.Sessions().DeleteAll(ctx, db); err != nil {
		return err
	}
	if _, err := pgmodels.PeerLogs().DeleteAll(ctx, db); err != nil {
		return err
	}
	if _, err := pgmodels.Peers().DeleteAll(ctx, db); err != nil {
		return err
	}
	if _, err := pgmodels.Crawls().DeleteAll(ctx, db); err != nil {
		return err
	}
	if _, err := pgmodels.AgentVersions().DeleteAll(ctx, db); err != nil {
		return err
	}
	if _, err := pgmodels.Protocols().DeleteAll(ctx, db); err != nil {
		return err
	}
	if _, err := pgmodels.ProtocolsSets().DeleteAll(ctx, db); err != nil {
		return err
	}
	if _, err := pgmodels.Visits().DeleteAll(ctx, db); err != nil {
		return err
	}
	if _, err := pgmodels.CrawlProperties().DeleteAll(ctx, db); err != nil {
		return err
	}
	return nil
}

func setup(t *testing.T) (context.Context, *PostgresClient, func(t *testing.T)) {
	ctx := context.Background()

	c := PostgresClientConfig{
		DatabaseHost:           "localhost",
		DatabasePort:           5433,
		DatabaseName:           "nebula_test",
		DatabasePassword:       "password_test",
		DatabaseUser:           "nebula_test",
		DatabaseSSL:            "disable",
		ApplyMigrations:        true,
		AgentVersionsCacheSize: 100,
		ProtocolsCacheSize:     100,
		ProtocolsSetCacheSize:  100,
		MeterProvider:          mnoop.NewMeterProvider(),
		TracerProvider:         tnoop.NewTracerProvider(),
	}

	client, err := NewPostgresClient(ctx, &c)
	require.NoError(t, err)

	return ctx, client, func(t *testing.T) {
		err = clearDatabase(context.Background(), client.dbh)
		require.NoError(t, err)
		err = client.Close()
		require.NoError(t, err)
	}
}

func TestClient_InitCrawl(t *testing.T) {
	ctx, client, teardown := setup(t)
	defer teardown(t)

	err := client.InitCrawl(ctx, "test")
	require.NoError(t, err)

	assert.NotZero(t, client.crawl.ID)
	assert.Equal(t, pgmodels.CrawlStateStarted, client.crawl.State)
	assert.NotZero(t, client.crawl.StartedAt)
	assert.False(t, client.crawl.CrawledPeers.Valid)
	assert.False(t, client.crawl.DialablePeers.Valid)
	assert.False(t, client.crawl.UndialablePeers.Valid)
	assert.False(t, client.crawl.FinishedAt.Valid)

	dbCrawl, err := pgmodels.Crawls(pgmodels.CrawlWhere.ID.EQ(client.crawl.ID)).One(ctx, client.dbh)
	require.NoError(t, err)

	assert.NotZero(t, dbCrawl.ID)
	assert.Equal(t, pgmodels.CrawlStateStarted, dbCrawl.State)
	assert.NotZero(t, dbCrawl.StartedAt)
	assert.False(t, dbCrawl.CrawledPeers.Valid)
	assert.False(t, dbCrawl.DialablePeers.Valid)
	assert.False(t, dbCrawl.UndialablePeers.Valid)
	assert.False(t, dbCrawl.FinishedAt.Valid)
}

func TestClient_GetOrCreateAgentVersion(t *testing.T) {
	ctx, client, teardown := setup(t)
	defer teardown(t)

	id, err := client.GetOrCreateAgentVersionID(ctx, client.Handle(), "")
	assert.Error(t, err)
	assert.Nil(t, id)
	client.agentVersions.Purge()

	id, err = client.GetOrCreateAgentVersionID(ctx, client.Handle(), "agent-1")
	assert.NoError(t, err)
	assert.Greater(t, *id, 0)
	prevID := id
	client.agentVersions.Purge()

	id, err = client.GetOrCreateAgentVersionID(ctx, client.Handle(), "agent-1")
	assert.NoError(t, err)
	assert.Greater(t, *id, 0)
	assert.Equal(t, *prevID, *id)
	client.agentVersions.Purge()

	id, err = client.GetOrCreateAgentVersionID(ctx, client.Handle(), "agent-2")
	assert.NoError(t, err)
	assert.Greater(t, *id, 0)
	assert.NotEqual(t, *prevID, *id)
}

func TestClient_GetOrCreateProtocol(t *testing.T) {
	ctx, client, teardown := setup(t)
	defer teardown(t)

	id, err := client.GetOrCreateProtocol(ctx, client.Handle(), "")
	assert.Error(t, err)
	assert.Nil(t, id)
	client.protocols.Purge()

	id, err = client.GetOrCreateProtocol(ctx, client.Handle(), "protocol-1")
	assert.NoError(t, err)
	assert.Greater(t, *id, 0)
	prevID := id
	client.protocols.Purge()

	id, err = client.GetOrCreateProtocol(ctx, client.Handle(), "protocol-1")
	assert.NoError(t, err)
	assert.Greater(t, *id, 0)
	assert.Equal(t, *prevID, *id)
	client.protocols.Purge()

	id, err = client.GetOrCreateProtocol(ctx, client.Handle(), "protocol-2")
	assert.NoError(t, err)
	assert.Greater(t, *id, 0)
	assert.NotEqual(t, *prevID, *id)
}

func TestClient_GetOrCreateProtocolsSetID(t *testing.T) {
	ctx, client, teardown := setup(t)
	defer teardown(t)

	id, err := client.GetOrCreateProtocolsSetID(ctx, client.Handle(), []string{})
	assert.Error(t, err)
	assert.Nil(t, id)
	client.protocolsSets.Purge()

	id, err = client.GetOrCreateProtocolsSetID(ctx, client.Handle(), []string{"protocol-1", "protocol-2"})
	assert.NoError(t, err)
	assert.Greater(t, *id, 0)
	prevID := id
	client.protocolsSets.Purge()

	id, err = client.GetOrCreateProtocolsSetID(ctx, client.Handle(), []string{"protocol-2", "protocol-1"})
	assert.NoError(t, err)
	assert.Equal(t, *prevID, *id)
	client.protocolsSets.Purge()

	id, err = client.GetOrCreateProtocolsSetID(ctx, client.Handle(), []string{"protocol-2", "protocol-1", "protocol-3"})
	assert.NoError(t, err)
	assert.Greater(t, *id, 0)
	assert.NotEqual(t, *prevID, *id)
}

func TestClient_PersistCrawlProperties(t *testing.T) {
	ctx, client, teardown := setup(t)
	defer teardown(t)

	err := client.InitCrawl(ctx, "test")
	require.NoError(t, err)

	props := map[string]map[string]int{}
	props["agent_version"] = map[string]int{
		"agent-1": 1,
		"agent-2": 2,
	}
	props["protocol"] = map[string]int{
		"protocols-1": 1,
		"protocols-2": 2,
	}
	props["error"] = map[string]int{
		"unknown":    1,
		"io_timeout": 2,
	}

	err = client.InsertCrawlProperties(ctx, props)
	require.NoError(t, err)

	cps, err := pgmodels.CrawlProperties(pgmodels.CrawlPropertyWhere.CrawlID.EQ(client.crawl.ID)).All(ctx, client.dbh)
	require.NoError(t, err)

	for _, cp := range cps {
		if !cp.ProtocolID.IsZero() {
			protocol, err := pgmodels.Protocols(pgmodels.ProtocolWhere.ID.EQ(cp.ProtocolID.Int)).One(ctx, client.dbh)
			require.NoError(t, err)
			assert.Equal(t, cp.Count, props["protocol"][protocol.Protocol])
		} else if !cp.AgentVersionID.IsZero() {
			agentVersion, err := pgmodels.AgentVersions(pgmodels.AgentVersionWhere.ID.EQ(cp.AgentVersionID.Int)).One(ctx, client.dbh)
			require.NoError(t, err)
			assert.Equal(t, cp.Count, props["agent_version"][agentVersion.AgentVersion])
		} else if !cp.Error.IsZero() {
			assert.Equal(t, cp.Count, props["error"][cp.Error.String])
		}
	}
	assert.Equal(t, 6, len(cps))
}

func TestClient_QueryBootstrapPeers(t *testing.T) {
	ctx, client, teardown := setup(t)
	defer teardown(t)

	peers, err := client.QueryBootstrapPeers(ctx, 10)
	require.NoError(t, err)
	assert.Len(t, peers, 0)
}

func TestClient_PersistCrawlVisit(t *testing.T) {
	ctx, client, teardown := setup(t)
	defer teardown(t)

	err := client.InitCrawl(ctx, "test")
	require.NoError(t, err)

	peerID, err := lp2ptest.RandPeerID()
	require.NoError(t, err)

	ma1, err := multiaddr.NewMultiaddr("/ip4/100.0.0.1/tcp/2000")
	require.NoError(t, err)

	ma2, err := multiaddr.NewMultiaddr("/ip4/100.0.0.2/udp/3000")
	require.NoError(t, err)

	protocols := []string{"protocol-1", "protocol-2"}
	agentVersion := "agent-1"

	visitStart := time.Now().Add(-time.Second)
	visitEnd := time.Now()

	args := &VisitArgs{
		PeerID:          peerID,
		DialMaddrs:      []multiaddr.Multiaddr{ma1, ma2},
		Protocols:       protocols,
		AgentVersion:    agentVersion,
		ConnectDuration: time.Second,
		CrawlDuration:   time.Second,
		VisitStartedAt:  visitStart,
		VisitEndedAt:    visitEnd,
		ConnectErrorStr: pgmodels.NetErrorIoTimeout,
		CrawlErrorStr:   "",
		VisitType:       VisitTypeCrawl,
		Neighbors:       nil,
		ErrorBits:       0,
		Properties:      json.RawMessage(marshalProperties(t, "is_exposed", true)),
	}
	err = client.InsertVisit(ctx, args)
	require.NoError(t, err)
}

func TestClient_SessionScenario_1(t *testing.T) {
	ctx, client, teardown := setup(t)
	defer teardown(t)

	err := client.InitCrawl(ctx, "test")
	require.NoError(t, err)

	peerID, err := lp2ptest.RandPeerID()
	require.NoError(t, err)

	ma1, err := multiaddr.NewMultiaddr("/ip4/100.0.0.1/tcp/2000")
	require.NoError(t, err)

	ma2, err := multiaddr.NewMultiaddr("/ip4/100.0.0.2/udp/3000")
	require.NoError(t, err)

	protocols := []string{"protocol-1", "protocol-2"}
	agentVersion := "agent-1"

	visitStart := time.Now().Add(-time.Second)
	visitEnd := time.Now()

	args := &VisitArgs{
		PeerID:          peerID,
		DialMaddrs:      []multiaddr.Multiaddr{ma1, ma2},
		Protocols:       protocols,
		AgentVersion:    agentVersion,
		ConnectDuration: time.Second,
		CrawlDuration:   time.Second,
		VisitStartedAt:  visitStart,
		VisitEndedAt:    visitEnd,
		ConnectErrorStr: "",
		CrawlErrorStr:   "",
		VisitType:       VisitTypeCrawl,
		Neighbors:       nil,
		ErrorBits:       0,
		Properties:      json.RawMessage(marshalProperties(t, "is_exposed", true)),
	}
	err = client.InsertVisit(ctx, args)
	require.NoError(t, err)

	dbPeer := fetchPeerByMultihash(t, ctx, client.Handle(), peerID.String())

	assert.Equal(t, dbPeer.R.AgentVersion.AgentVersion, agentVersion)
	assert.Equal(t, dbPeer.MultiHash, peerID.String())
	assert.Len(t, dbPeer.R.MultiAddresses, 2)
	assert.True(t, unmarshalProperties(t, dbPeer.Properties.JSON)["is_exposed"].(bool))

	for _, ma := range dbPeer.R.MultiAddresses {
		assert.True(t, ma.Maddr == ma1.String() || ma.Maddr == ma2.String())
	}
	session := dbPeer.R.SessionsOpen
	sessionID1 := session.ID
	assert.Equal(t, session.PeerID, dbPeer.ID)
	assert.Equal(t, session.SuccessfulVisitsCount, 1)
	assert.Equal(t, session.FailedVisitsCount, int16(0))
	assert.Equal(t, session.State, pgmodels.SessionStateOpen)
	assert.InDelta(t, session.FirstSuccessfulVisit.UnixNano(), visitStart.UnixNano(), float64(time.Microsecond))
	assert.InDelta(t, session.LastVisitedAt.UnixNano(), visitEnd.UnixNano(), float64(time.Microsecond))
	assert.True(t, session.FirstFailedVisit.IsZero())
	assert.True(t, session.FinishReason.IsZero())
	assert.True(t, session.LastFailedVisit.IsZero())

	visitStart = time.Now().Add(-time.Second)
	visitEnd = time.Now()

	args = &VisitArgs{
		PeerID:         peerID,
		DialMaddrs:     []multiaddr.Multiaddr{ma1, ma2},
		DialDuration:   time.Second,
		VisitStartedAt: visitStart,
		VisitEndedAt:   visitEnd,
		VisitType:      VisitTypeDial,
	}
	err = client.InsertVisit(ctx, args)
	require.NoError(t, err)

	dbPeer = fetchPeerByMultihash(t, ctx, client.Handle(), peerID.String())

	assert.True(t, dbPeer.Properties.Valid)

	session = dbPeer.R.SessionsOpen
	assert.Equal(t, session.PeerID, dbPeer.ID)
	assert.Equal(t, session.SuccessfulVisitsCount, 2)
	assert.Equal(t, session.FailedVisitsCount, int16(0))
	assert.Equal(t, session.State, pgmodels.SessionStateOpen)
	assert.NotEqual(t, session.FirstSuccessfulVisit.UnixMicro(), visitStart.UnixMicro())
	assert.InDelta(t, session.LastVisitedAt.UnixNano(), visitEnd.UnixNano(), float64(time.Microsecond))
	assert.True(t, session.FirstFailedVisit.IsZero())
	assert.True(t, session.FinishReason.IsZero())
	assert.True(t, session.LastFailedVisit.IsZero())

	visitStart = time.Now().Add(-time.Second)
	visitEnd = time.Now()

	args = &VisitArgs{
		PeerID:          peerID,
		DialMaddrs:      []multiaddr.Multiaddr{ma1, ma2},
		Protocols:       protocols,
		AgentVersion:    agentVersion,
		DialDuration:    time.Second,
		VisitStartedAt:  visitStart,
		VisitEndedAt:    visitEnd,
		ConnectErrorStr: pgmodels.NetErrorConnectionRefused,
		VisitType:       VisitTypeDial,
	}
	err = client.InsertVisit(ctx, args)
	require.NoError(t, err)
	dbPeer = fetchPeerByMultihash(t, ctx, client.Handle(), peerID.String())

	assert.Nil(t, dbPeer.R.SessionsOpen)

	s, err := pgmodels.Sessions(pgmodels.SessionWhere.State.EQ(pgmodels.SessionStateClosed)).One(ctx, client.Handle())
	require.NoError(t, err)

	assert.Equal(t, s.PeerID, dbPeer.ID)
	assert.Equal(t, s.SuccessfulVisitsCount, 2)
	assert.Equal(t, s.FailedVisitsCount, int16(1))
	assert.Equal(t, s.State, pgmodels.SessionStateClosed)
	assert.NotEqual(t, s.FirstSuccessfulVisit.UnixMicro(), visitStart.UnixMicro())
	assert.InDelta(t, s.LastVisitedAt.UnixNano(), visitEnd.UnixNano(), float64(time.Microsecond))
	assert.InDelta(t, s.FirstFailedVisit.Time.UnixNano(), visitStart.UnixNano(), float64(time.Microsecond))
	assert.InDelta(t, s.LastFailedVisit.Time.UnixNano(), visitEnd.UnixNano(), float64(time.Microsecond))
	assert.Equal(t, s.FinishReason.String, pgmodels.NetErrorConnectionRefused)

	// allow another crawl initialization
	client.crawl = nil

	err = client.InitCrawl(ctx, "test")
	require.NoError(t, err)

	visitStart = time.Now().Add(-time.Second)
	visitEnd = time.Now()

	args = &VisitArgs{
		PeerID:          peerID,
		DialMaddrs:      []multiaddr.Multiaddr{ma1, ma2},
		Protocols:       []string{},
		AgentVersion:    "",
		ConnectDuration: time.Second,
		CrawlDuration:   time.Second,
		VisitStartedAt:  visitStart,
		VisitEndedAt:    visitEnd,
		VisitType:       VisitTypeCrawl,
		Properties:      json.RawMessage(marshalProperties(t, "is_exposed", true)),
	}
	err = client.InsertVisit(ctx, args)
	require.NoError(t, err)

	visitStart = time.Now().Add(-time.Second)
	visitEnd = time.Now()

	args = &VisitArgs{
		PeerID:          peerID,
		DialMaddrs:      []multiaddr.Multiaddr{ma1, ma2},
		DialDuration:    time.Second,
		VisitStartedAt:  visitStart,
		VisitEndedAt:    visitEnd,
		ConnectErrorStr: pgmodels.NetErrorNegotiateSecurityProtocol,
		VisitType:       VisitTypeDial,
	}
	err = client.InsertVisit(ctx, args)
	dbPeer = fetchPeerByMultihash(t, ctx, client.Handle(), peerID.String())

	assert.Equal(t, dbPeer.R.AgentVersion.AgentVersion, agentVersion)
	assert.Equal(t, dbPeer.MultiHash, peerID.String())
	assert.Len(t, dbPeer.R.MultiAddresses, 2)

	newSession, err := pgmodels.Sessions(qm.OrderBy(pgmodels.SessionColumns.UpdatedAt+" desc")).One(ctx, client.Handle())
	require.NoError(t, err)

	sessionID2 := newSession.ID
	require.NotEqual(t, sessionID1, sessionID2)

	assert.Equal(t, newSession.PeerID, dbPeer.ID)
	assert.Equal(t, newSession.SuccessfulVisitsCount, 1)
	assert.Equal(t, newSession.FailedVisitsCount, int16(1))
	assert.Equal(t, newSession.State, pgmodels.SessionStateClosed)
	assert.NotEqual(t, newSession.FirstSuccessfulVisit.UnixMicro(), visitStart.UnixMicro())
	assert.InDelta(t, newSession.LastVisitedAt.UnixNano(), visitEnd.UnixNano(), float64(time.Microsecond))
	assert.InDelta(t, newSession.FirstFailedVisit.Time.UnixNano(), visitStart.UnixNano(), float64(time.Microsecond))
	assert.InDelta(t, newSession.LastFailedVisit.Time.UnixNano(), visitEnd.UnixNano(), float64(time.Microsecond))
	assert.Equal(t, newSession.FinishReason.String, pgmodels.NetErrorNegotiateSecurityProtocol)

	err = s.Reload(ctx, client.Handle())
	require.NoError(t, err)

	// untouched:
	assert.Equal(t, s.PeerID, dbPeer.ID)
	assert.Equal(t, s.SuccessfulVisitsCount, 2)
	assert.Equal(t, s.FailedVisitsCount, int16(1))
	assert.Equal(t, s.State, pgmodels.SessionStateClosed)
	assert.NotEqual(t, s.FirstSuccessfulVisit.UnixMicro(), visitStart.UnixMicro())
	assert.NotEqual(t, s.LastVisitedAt.UnixMicro(), visitStart.UnixMicro())
	assert.NotEqual(t, s.FirstFailedVisit.Time.UnixMicro(), visitStart.UnixMicro())
	assert.NotEqual(t, s.LastFailedVisit.Time.UnixMicro(), visitEnd.UnixMicro())
	assert.Equal(t, s.FinishReason.String, pgmodels.NetErrorConnectionRefused)
}

func TestClient_SessionScenario_2(t *testing.T) {
	ctx, client, teardown := setup(t)
	defer teardown(t)

	err := client.InitCrawl(ctx, "test")
	require.NoError(t, err)

	peerID, err := lp2ptest.RandPeerID()
	require.NoError(t, err)

	ma1, err := multiaddr.NewMultiaddr("/ip4/100.0.0.3/tcp/2000")
	require.NoError(t, err)

	ma2, err := multiaddr.NewMultiaddr("/ip4/100.0.0.4/udp/3000")
	require.NoError(t, err)

	protocols := []string{"protocol-2", "protocol-3"}
	agentVersion := "agent-1"

	visitStart := time.Now()
	visitEnd := time.Now().Add(time.Second)

	args := &VisitArgs{
		PeerID:          peerID,
		DialMaddrs:      []multiaddr.Multiaddr{ma1, ma2},
		Protocols:       protocols,
		AgentVersion:    agentVersion,
		ConnectDuration: time.Second,
		CrawlDuration:   time.Second,
		VisitStartedAt:  visitStart,
		VisitEndedAt:    visitEnd,
		VisitType:       VisitTypeCrawl,
		Properties:      json.RawMessage(marshalProperties(t, "is_exposed", false)),
	}
	err = client.InsertVisit(ctx, args)
	require.NoError(t, err)
	visitStart = time.Now().Add(22 * time.Hour)
	visitEnd = time.Now().Add(22 * time.Hour).Add(time.Second)

	args = &VisitArgs{
		PeerID:         peerID,
		DialMaddrs:     []multiaddr.Multiaddr{ma1, ma2},
		DialDuration:   time.Second,
		VisitStartedAt: visitStart,
		VisitEndedAt:   visitEnd,
		VisitType:      VisitTypeDial,
	}
	err = client.InsertVisit(ctx, args)

	visitStart = time.Now().Add(23 * time.Hour).Add(time.Second)
	visitEnd = time.Now().Add(23 * time.Hour)

	args = &VisitArgs{
		PeerID:          peerID,
		DialMaddrs:      []multiaddr.Multiaddr{ma1, ma2},
		DialDuration:    time.Second,
		VisitStartedAt:  visitStart,
		VisitEndedAt:    visitEnd,
		VisitType:       VisitTypeDial,
		ConnectErrorStr: pgmodels.NetErrorIoTimeout,
	}
	err = client.InsertVisit(ctx, args)
	require.NoError(t, err)
	dbPeer := fetchPeerByMultihash(t, ctx, client.Handle(), peerID.String())

	session := dbPeer.R.SessionsOpen
	assert.Equal(t, session.PeerID, dbPeer.ID)
	assert.Equal(t, session.SuccessfulVisitsCount, 2)
	assert.Equal(t, session.FailedVisitsCount, int16(1))
	assert.Equal(t, session.RecoveredCount, 0)
	assert.Equal(t, session.State, pgmodels.SessionStatePending)
	assert.NotEqual(t, session.FirstSuccessfulVisit.UnixMicro(), visitStart.UnixMicro())
	assert.InDelta(t, session.LastVisitedAt.UnixNano(), visitEnd.UnixNano(), float64(time.Microsecond))
	assert.InDelta(t, session.FirstFailedVisit.Time.UnixNano(), visitStart.UnixNano(), float64(time.Microsecond))
	assert.Equal(t, session.FinishReason.String, pgmodels.NetErrorIoTimeout)
	assert.InDelta(t, session.LastFailedVisit.Time.UnixNano(), visitEnd.UnixNano(), float64(time.Microsecond))

	visitStart = time.Now().Add(-time.Second)
	visitEnd = time.Now()
	args = &VisitArgs{
		PeerID:          peerID,
		DialMaddrs:      []multiaddr.Multiaddr{ma1, ma2},
		DialDuration:    time.Second,
		VisitStartedAt:  visitStart,
		VisitEndedAt:    visitEnd,
		VisitType:       VisitTypeDial,
		ConnectErrorStr: "",
	}
	err = client.InsertVisit(ctx, args)
	require.NoError(t, err)
	dbPeer = fetchPeerByMultihash(t, ctx, client.Handle(), peerID.String())

	session = dbPeer.R.SessionsOpen
	assert.Equal(t, session.PeerID, dbPeer.ID)
	assert.Equal(t, session.SuccessfulVisitsCount, 3)
	assert.Equal(t, session.FailedVisitsCount, int16(0))
	assert.Equal(t, session.State, pgmodels.SessionStateOpen)
	assert.Equal(t, session.RecoveredCount, 1)
	assert.NotEqual(t, session.FirstSuccessfulVisit.UnixMicro(), visitStart.UnixMicro())
	assert.InDelta(t, session.LastSuccessfulVisit.UnixNano(), visitEnd.UnixNano(), float64(time.Microsecond))
	assert.InDelta(t, session.LastVisitedAt.UnixNano(), visitEnd.UnixNano(), float64(time.Microsecond))
	assert.True(t, session.FirstFailedVisit.IsZero())
	assert.True(t, session.FinishReason.IsZero())
	assert.True(t, session.LastFailedVisit.IsZero())

	count, err := pgmodels.Sessions().Count(ctx, client.Handle())
	require.NoError(t, err)
	assert.Equal(t, int64(1), count)
}

func TestClient_UpsertPeer(t *testing.T) {
	ctx, client, teardown := setup(t)
	defer teardown(t)

	dbAgentVersionID, err := client.GetOrCreateAgentVersionID(ctx, client.Handle(), "agent-1")
	require.NoError(t, err)

	dbProtocolsSetID, err := client.GetOrCreateProtocolsSetID(ctx, client.Handle(), []string{"protocol-1", "protocol-2"})
	require.NoError(t, err)

	peerID, err := lp2ptest.RandPeerID()
	require.NoError(t, err)

	dbPeerID, err := client.UpsertPeer(peerID.String(), null.IntFromPtr(nil), null.IntFromPtr(nil), null.JSONFromPtr(nil))
	require.NoError(t, err)
	assert.NotZero(t, dbPeerID)

	dbPeer := fetchPeer(t, ctx, client.Handle(), dbPeerID)
	assert.True(t, dbPeer.AgentVersionID.IsZero())
	assert.True(t, dbPeer.ProtocolsSetID.IsZero())
	assert.True(t, dbPeer.Properties.IsZero())

	dbPeerID, err = client.UpsertPeer(peerID.String(), null.IntFromPtr(dbAgentVersionID), null.IntFromPtr(nil), null.JSONFromPtr(nil))
	require.NoError(t, err)

	dbPeer = fetchPeer(t, ctx, client.Handle(), dbPeerID)
	assert.Equal(t, dbPeer.AgentVersionID.Int, *dbAgentVersionID)
	assert.True(t, dbPeer.ProtocolsSetID.IsZero())
	assert.True(t, dbPeer.Properties.IsZero())

	dbPeerID, err = client.UpsertPeer(peerID.String(), null.IntFromPtr(nil), null.IntFromPtr(dbProtocolsSetID), null.JSONFromPtr(nil))
	require.NoError(t, err)

	dbPeer = fetchPeer(t, ctx, client.Handle(), dbPeerID)
	assert.Equal(t, dbPeer.AgentVersionID.Int, *dbAgentVersionID)
	assert.Equal(t, dbPeer.ProtocolsSetID.Int, *dbProtocolsSetID)
	assert.True(t, dbPeer.Properties.IsZero())

	dbPeerID, err = client.UpsertPeer(peerID.String(), null.IntFromPtr(nil), null.IntFromPtr(dbProtocolsSetID), null.JSONFrom(marshalProperties(t, "is_exposed", false)))
	require.NoError(t, err)

	dbPeer = fetchPeer(t, ctx, client.Handle(), dbPeerID)
	assert.Equal(t, dbPeer.AgentVersionID.Int, *dbAgentVersionID)
	assert.Equal(t, dbPeer.ProtocolsSetID.Int, *dbProtocolsSetID)
	assert.False(t, unmarshalProperties(t, dbPeer.Properties.JSON)["is_exposed"].(bool))

	dbPeerID, err = client.UpsertPeer(peerID.String(), null.IntFromPtr(nil), null.IntFromPtr(nil), null.JSONFromPtr(nil))
	require.NoError(t, err)

	dbPeer = fetchPeer(t, ctx, client.Handle(), dbPeerID)
	assert.Equal(t, dbPeer.AgentVersionID.Int, *dbAgentVersionID)
	assert.Equal(t, dbPeer.ProtocolsSetID.Int, *dbProtocolsSetID)
	assert.False(t, unmarshalProperties(t, dbPeer.Properties.JSON)["is_exposed"].(bool))

	dbAgentVersionID, err = client.GetOrCreateAgentVersionID(ctx, client.Handle(), "agent-2")
	require.NoError(t, err)

	dbProtocolsSetID, err = client.GetOrCreateProtocolsSetID(ctx, client.Handle(), []string{"protocol-3", "protocol-2"})
	require.NoError(t, err)

	dbPeerID, err = client.UpsertPeer(peerID.String(), null.IntFromPtr(dbAgentVersionID), null.IntFromPtr(dbProtocolsSetID), null.JSONFrom(marshalProperties(t, "is_exposed", true)))
	require.NoError(t, err)

	dbPeer = fetchPeer(t, ctx, client.Handle(), dbPeerID)
	assert.Equal(t, dbPeer.AgentVersionID.Int, *dbAgentVersionID)
	assert.Equal(t, dbPeer.ProtocolsSetID.Int, *dbProtocolsSetID)
	assert.True(t, unmarshalProperties(t, dbPeer.Properties.JSON)["is_exposed"].(bool))
}

func fetchPeer(t *testing.T, ctx context.Context, exec boil.ContextExecutor, dbPeerID int) *pgmodels.Peer {
	dbPeer, err := pgmodels.Peers(
		pgmodels.PeerWhere.ID.EQ(dbPeerID),
		qm.Load(pgmodels.PeerRels.AgentVersion),
		qm.Load(pgmodels.PeerRels.MultiAddresses),
		qm.Load(pgmodels.PeerRels.ProtocolsSet),
		qm.Load(pgmodels.PeerRels.SessionsOpen),
	).One(ctx, exec)
	require.NoError(t, err)
	return dbPeer
}

func fetchPeerByMultihash(t *testing.T, ctx context.Context, exec boil.ContextExecutor, multiHash string) *pgmodels.Peer {
	dbPeer, err := pgmodels.Peers(
		pgmodels.PeerWhere.MultiHash.EQ(multiHash),
		qm.Load(pgmodels.PeerRels.AgentVersion),
		qm.Load(pgmodels.PeerRels.MultiAddresses),
		qm.Load(pgmodels.PeerRels.ProtocolsSet),
		qm.Load(pgmodels.PeerRels.SessionsOpen),
	).One(ctx, exec)
	require.NoError(t, err)
	return dbPeer
}

func marshalProperties(t testing.TB, args ...any) []byte {
	if len(args)%2 != 0 {
		t.Fatal("args must be multiple of 2")
	}

	properties := map[string]any{}
	for i := 0; i < len(args); i += 2 {
		key, ok := args[i].(string)
		require.True(t, ok)
		value := args[i+1]

		properties[key] = value
	}

	data, err := json.Marshal(properties)
	require.NoError(t, err)
	return data
}

func unmarshalProperties(t testing.TB, data []byte) map[string]any {
	properties := map[string]any{}
	err := json.Unmarshal(data, &properties)
	require.NoError(t, err)
	return properties
}
