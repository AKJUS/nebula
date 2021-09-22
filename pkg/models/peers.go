// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// Peer is an object representing the database table.
type Peer struct {
	ID                string            `boil:"id" json:"id" toml:"id" yaml:"id"`
	MultiAddresses    types.StringArray `boil:"multi_addresses" json:"multi_addresses,omitempty" toml:"multi_addresses" yaml:"multi_addresses,omitempty"`
	UpdatedAt         time.Time         `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt         time.Time         `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	OldMultiAddresses types.StringArray `boil:"old_multi_addresses" json:"old_multi_addresses,omitempty" toml:"old_multi_addresses" yaml:"old_multi_addresses,omitempty"`
	AgentVersion      null.String       `boil:"agent_version" json:"agent_version,omitempty" toml:"agent_version" yaml:"agent_version,omitempty"`
	Protocol          types.StringArray `boil:"protocol" json:"protocol,omitempty" toml:"protocol" yaml:"protocol,omitempty"`

	R *peerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L peerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PeerColumns = struct {
	ID                string
	MultiAddresses    string
	UpdatedAt         string
	CreatedAt         string
	OldMultiAddresses string
	AgentVersion      string
	Protocol          string
}{
	ID:                "id",
	MultiAddresses:    "multi_addresses",
	UpdatedAt:         "updated_at",
	CreatedAt:         "created_at",
	OldMultiAddresses: "old_multi_addresses",
	AgentVersion:      "agent_version",
	Protocol:          "protocol",
}

var PeerTableColumns = struct {
	ID                string
	MultiAddresses    string
	UpdatedAt         string
	CreatedAt         string
	OldMultiAddresses string
	AgentVersion      string
	Protocol          string
}{
	ID:                "peers.id",
	MultiAddresses:    "peers.multi_addresses",
	UpdatedAt:         "peers.updated_at",
	CreatedAt:         "peers.created_at",
	OldMultiAddresses: "peers.old_multi_addresses",
	AgentVersion:      "peers.agent_version",
	Protocol:          "peers.protocol",
}

// Generated where

type whereHelpertypes_StringArray struct{ field string }

func (w whereHelpertypes_StringArray) EQ(x types.StringArray) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpertypes_StringArray) NEQ(x types.StringArray) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpertypes_StringArray) IsNull() qm.QueryMod { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpertypes_StringArray) IsNotNull() qm.QueryMod {
	return qmhelper.WhereIsNotNull(w.field)
}
func (w whereHelpertypes_StringArray) LT(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_StringArray) LTE(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_StringArray) GT(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_StringArray) GTE(x types.StringArray) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var PeerWhere = struct {
	ID                whereHelperstring
	MultiAddresses    whereHelpertypes_StringArray
	UpdatedAt         whereHelpertime_Time
	CreatedAt         whereHelpertime_Time
	OldMultiAddresses whereHelpertypes_StringArray
	AgentVersion      whereHelpernull_String
	Protocol          whereHelpertypes_StringArray
}{
	ID:                whereHelperstring{field: "\"peers\".\"id\""},
	MultiAddresses:    whereHelpertypes_StringArray{field: "\"peers\".\"multi_addresses\""},
	UpdatedAt:         whereHelpertime_Time{field: "\"peers\".\"updated_at\""},
	CreatedAt:         whereHelpertime_Time{field: "\"peers\".\"created_at\""},
	OldMultiAddresses: whereHelpertypes_StringArray{field: "\"peers\".\"old_multi_addresses\""},
	AgentVersion:      whereHelpernull_String{field: "\"peers\".\"agent_version\""},
	Protocol:          whereHelpertypes_StringArray{field: "\"peers\".\"protocol\""},
}

// PeerRels is where relationship names are stored.
var PeerRels = struct {
	Latencies string
	Sessions  string
}{
	Latencies: "Latencies",
	Sessions:  "Sessions",
}

// peerR is where relationships are stored.
type peerR struct {
	Latencies LatencySlice `boil:"Latencies" json:"Latencies" toml:"Latencies" yaml:"Latencies"`
	Sessions  SessionSlice `boil:"Sessions" json:"Sessions" toml:"Sessions" yaml:"Sessions"`
}

// NewStruct creates a new relationship struct
func (*peerR) NewStruct() *peerR {
	return &peerR{}
}

// peerL is where Load methods for each relationship are stored.
type peerL struct{}

var (
	peerAllColumns            = []string{"id", "multi_addresses", "updated_at", "created_at", "old_multi_addresses", "agent_version", "protocol"}
	peerColumnsWithoutDefault = []string{"id", "multi_addresses", "updated_at", "created_at", "old_multi_addresses", "agent_version", "protocol"}
	peerColumnsWithDefault    = []string{}
	peerPrimaryKeyColumns     = []string{"id"}
)

type (
	// PeerSlice is an alias for a slice of pointers to Peer.
	// This should almost always be used instead of []Peer.
	PeerSlice []*Peer
	// PeerHook is the signature for custom Peer hook methods
	PeerHook func(context.Context, boil.ContextExecutor, *Peer) error

	peerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	peerType                 = reflect.TypeOf(&Peer{})
	peerMapping              = queries.MakeStructMapping(peerType)
	peerPrimaryKeyMapping, _ = queries.BindMapping(peerType, peerMapping, peerPrimaryKeyColumns)
	peerInsertCacheMut       sync.RWMutex
	peerInsertCache          = make(map[string]insertCache)
	peerUpdateCacheMut       sync.RWMutex
	peerUpdateCache          = make(map[string]updateCache)
	peerUpsertCacheMut       sync.RWMutex
	peerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var peerBeforeInsertHooks []PeerHook
var peerBeforeUpdateHooks []PeerHook
var peerBeforeDeleteHooks []PeerHook
var peerBeforeUpsertHooks []PeerHook

var peerAfterInsertHooks []PeerHook
var peerAfterSelectHooks []PeerHook
var peerAfterUpdateHooks []PeerHook
var peerAfterDeleteHooks []PeerHook
var peerAfterUpsertHooks []PeerHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Peer) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peerBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Peer) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peerBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Peer) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peerBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Peer) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peerBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Peer) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peerAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Peer) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peerAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Peer) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peerAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Peer) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peerAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Peer) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range peerAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPeerHook registers your hook function for all future operations.
func AddPeerHook(hookPoint boil.HookPoint, peerHook PeerHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		peerBeforeInsertHooks = append(peerBeforeInsertHooks, peerHook)
	case boil.BeforeUpdateHook:
		peerBeforeUpdateHooks = append(peerBeforeUpdateHooks, peerHook)
	case boil.BeforeDeleteHook:
		peerBeforeDeleteHooks = append(peerBeforeDeleteHooks, peerHook)
	case boil.BeforeUpsertHook:
		peerBeforeUpsertHooks = append(peerBeforeUpsertHooks, peerHook)
	case boil.AfterInsertHook:
		peerAfterInsertHooks = append(peerAfterInsertHooks, peerHook)
	case boil.AfterSelectHook:
		peerAfterSelectHooks = append(peerAfterSelectHooks, peerHook)
	case boil.AfterUpdateHook:
		peerAfterUpdateHooks = append(peerAfterUpdateHooks, peerHook)
	case boil.AfterDeleteHook:
		peerAfterDeleteHooks = append(peerAfterDeleteHooks, peerHook)
	case boil.AfterUpsertHook:
		peerAfterUpsertHooks = append(peerAfterUpsertHooks, peerHook)
	}
}

// One returns a single peer record from the query.
func (q peerQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Peer, error) {
	o := &Peer{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for peers")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Peer records from the query.
func (q peerQuery) All(ctx context.Context, exec boil.ContextExecutor) (PeerSlice, error) {
	var o []*Peer

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Peer slice")
	}

	if len(peerAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Peer records in the query.
func (q peerQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count peers rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q peerQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if peers exists")
	}

	return count > 0, nil
}

// Latencies retrieves all the latency's Latencies with an executor.
func (o *Peer) Latencies(mods ...qm.QueryMod) latencyQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"latencies\".\"peer_id\"=?", o.ID),
	)

	query := Latencies(queryMods...)
	queries.SetFrom(query.Query, "\"latencies\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"latencies\".*"})
	}

	return query
}

// Sessions retrieves all the session's Sessions with an executor.
func (o *Peer) Sessions(mods ...qm.QueryMod) sessionQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"sessions\".\"peer_id\"=?", o.ID),
	)

	query := Sessions(queryMods...)
	queries.SetFrom(query.Query, "\"sessions\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"sessions\".*"})
	}

	return query
}

// LoadLatencies allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (peerL) LoadLatencies(ctx context.Context, e boil.ContextExecutor, singular bool, maybePeer interface{}, mods queries.Applicator) error {
	var slice []*Peer
	var object *Peer

	if singular {
		object = maybePeer.(*Peer)
	} else {
		slice = *maybePeer.(*[]*Peer)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &peerR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &peerR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`latencies`),
		qm.WhereIn(`latencies.peer_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load latencies")
	}

	var resultSlice []*Latency
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice latencies")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on latencies")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for latencies")
	}

	if len(latencyAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Latencies = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &latencyR{}
			}
			foreign.R.Peer = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.PeerID {
				local.R.Latencies = append(local.R.Latencies, foreign)
				if foreign.R == nil {
					foreign.R = &latencyR{}
				}
				foreign.R.Peer = local
				break
			}
		}
	}

	return nil
}

// LoadSessions allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (peerL) LoadSessions(ctx context.Context, e boil.ContextExecutor, singular bool, maybePeer interface{}, mods queries.Applicator) error {
	var slice []*Peer
	var object *Peer

	if singular {
		object = maybePeer.(*Peer)
	} else {
		slice = *maybePeer.(*[]*Peer)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &peerR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &peerR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`sessions`),
		qm.WhereIn(`sessions.peer_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load sessions")
	}

	var resultSlice []*Session
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice sessions")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on sessions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for sessions")
	}

	if len(sessionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Sessions = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &sessionR{}
			}
			foreign.R.Peer = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.PeerID {
				local.R.Sessions = append(local.R.Sessions, foreign)
				if foreign.R == nil {
					foreign.R = &sessionR{}
				}
				foreign.R.Peer = local
				break
			}
		}
	}

	return nil
}

// AddLatencies adds the given related objects to the existing relationships
// of the peer, optionally inserting them as new records.
// Appends related to o.R.Latencies.
// Sets related.R.Peer appropriately.
func (o *Peer) AddLatencies(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Latency) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.PeerID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"latencies\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"peer_id"}),
				strmangle.WhereClause("\"", "\"", 2, latencyPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.PeerID = o.ID
		}
	}

	if o.R == nil {
		o.R = &peerR{
			Latencies: related,
		}
	} else {
		o.R.Latencies = append(o.R.Latencies, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &latencyR{
				Peer: o,
			}
		} else {
			rel.R.Peer = o
		}
	}
	return nil
}

// AddSessions adds the given related objects to the existing relationships
// of the peer, optionally inserting them as new records.
// Appends related to o.R.Sessions.
// Sets related.R.Peer appropriately.
func (o *Peer) AddSessions(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Session) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.PeerID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"sessions\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"peer_id"}),
				strmangle.WhereClause("\"", "\"", 2, sessionPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.PeerID = o.ID
		}
	}

	if o.R == nil {
		o.R = &peerR{
			Sessions: related,
		}
	} else {
		o.R.Sessions = append(o.R.Sessions, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &sessionR{
				Peer: o,
			}
		} else {
			rel.R.Peer = o
		}
	}
	return nil
}

// Peers retrieves all the records using an executor.
func Peers(mods ...qm.QueryMod) peerQuery {
	mods = append(mods, qm.From("\"peers\""))
	return peerQuery{NewQuery(mods...)}
}

// FindPeer retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPeer(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Peer, error) {
	peerObj := &Peer{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"peers\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, peerObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from peers")
	}

	if err = peerObj.doAfterSelectHooks(ctx, exec); err != nil {
		return peerObj, err
	}

	return peerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Peer) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no peers provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(peerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	peerInsertCacheMut.RLock()
	cache, cached := peerInsertCache[key]
	peerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			peerAllColumns,
			peerColumnsWithDefault,
			peerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(peerType, peerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(peerType, peerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"peers\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"peers\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into peers")
	}

	if !cached {
		peerInsertCacheMut.Lock()
		peerInsertCache[key] = cache
		peerInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Peer.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Peer) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	peerUpdateCacheMut.RLock()
	cache, cached := peerUpdateCache[key]
	peerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			peerAllColumns,
			peerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update peers, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"peers\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, peerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(peerType, peerMapping, append(wl, peerPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update peers row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for peers")
	}

	if !cached {
		peerUpdateCacheMut.Lock()
		peerUpdateCache[key] = cache
		peerUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q peerQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for peers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for peers")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PeerSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), peerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"peers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, peerPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in peer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all peer")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Peer) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no peers provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(peerColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	peerUpsertCacheMut.RLock()
	cache, cached := peerUpsertCache[key]
	peerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			peerAllColumns,
			peerColumnsWithDefault,
			peerColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			peerAllColumns,
			peerPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert peers, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(peerPrimaryKeyColumns))
			copy(conflict, peerPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"peers\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(peerType, peerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(peerType, peerMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert peers")
	}

	if !cached {
		peerUpsertCacheMut.Lock()
		peerUpsertCache[key] = cache
		peerUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Peer record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Peer) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Peer provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), peerPrimaryKeyMapping)
	sql := "DELETE FROM \"peers\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from peers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for peers")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q peerQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no peerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from peers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for peers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PeerSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(peerBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), peerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"peers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, peerPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from peer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for peers")
	}

	if len(peerAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Peer) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPeer(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PeerSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PeerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), peerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"peers\".* FROM \"peers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, peerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PeerSlice")
	}

	*o = slice

	return nil
}

// PeerExists checks if the Peer row exists.
func PeerExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"peers\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if peers exists")
	}

	return exists, nil
}
