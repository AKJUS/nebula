// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("AgentVersions", testAgentVersions)
	t.Run("CrawlProperties", testCrawlProperties)
	t.Run("Crawls", testCrawls)
	t.Run("IPAddresses", testIPAddresses)
	t.Run("Latencies", testLatencies)
	t.Run("MultiAddresses", testMultiAddresses)
	t.Run("MultiAddressesSets", testMultiAddressesSets)
	t.Run("MultiAddressesXIPAddresses", testMultiAddressesXIPAddresses)
	t.Run("Neighbors", testNeighbors)
	t.Run("Peers", testPeers)
	t.Run("PegasysConnections", testPegasysConnections)
	t.Run("PegasysNeighbours", testPegasysNeighbours)
	t.Run("Protocols", testProtocols)
	t.Run("ProtocolsSets", testProtocolsSets)
	t.Run("RawVisits", testRawVisits)
	t.Run("Sessions", testSessions)
	t.Run("Visits", testVisits)
}

func TestDelete(t *testing.T) {
	t.Run("AgentVersions", testAgentVersionsDelete)
	t.Run("CrawlProperties", testCrawlPropertiesDelete)
	t.Run("Crawls", testCrawlsDelete)
	t.Run("IPAddresses", testIPAddressesDelete)
	t.Run("Latencies", testLatenciesDelete)
	t.Run("MultiAddresses", testMultiAddressesDelete)
	t.Run("MultiAddressesSets", testMultiAddressesSetsDelete)
	t.Run("MultiAddressesXIPAddresses", testMultiAddressesXIPAddressesDelete)
	t.Run("Neighbors", testNeighborsDelete)
	t.Run("Peers", testPeersDelete)
	t.Run("PegasysConnections", testPegasysConnectionsDelete)
	t.Run("PegasysNeighbours", testPegasysNeighboursDelete)
	t.Run("Protocols", testProtocolsDelete)
	t.Run("ProtocolsSets", testProtocolsSetsDelete)
	t.Run("RawVisits", testRawVisitsDelete)
	t.Run("Sessions", testSessionsDelete)
	t.Run("Visits", testVisitsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("AgentVersions", testAgentVersionsQueryDeleteAll)
	t.Run("CrawlProperties", testCrawlPropertiesQueryDeleteAll)
	t.Run("Crawls", testCrawlsQueryDeleteAll)
	t.Run("IPAddresses", testIPAddressesQueryDeleteAll)
	t.Run("Latencies", testLatenciesQueryDeleteAll)
	t.Run("MultiAddresses", testMultiAddressesQueryDeleteAll)
	t.Run("MultiAddressesSets", testMultiAddressesSetsQueryDeleteAll)
	t.Run("MultiAddressesXIPAddresses", testMultiAddressesXIPAddressesQueryDeleteAll)
	t.Run("Neighbors", testNeighborsQueryDeleteAll)
	t.Run("Peers", testPeersQueryDeleteAll)
	t.Run("PegasysConnections", testPegasysConnectionsQueryDeleteAll)
	t.Run("PegasysNeighbours", testPegasysNeighboursQueryDeleteAll)
	t.Run("Protocols", testProtocolsQueryDeleteAll)
	t.Run("ProtocolsSets", testProtocolsSetsQueryDeleteAll)
	t.Run("RawVisits", testRawVisitsQueryDeleteAll)
	t.Run("Sessions", testSessionsQueryDeleteAll)
	t.Run("Visits", testVisitsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("AgentVersions", testAgentVersionsSliceDeleteAll)
	t.Run("CrawlProperties", testCrawlPropertiesSliceDeleteAll)
	t.Run("Crawls", testCrawlsSliceDeleteAll)
	t.Run("IPAddresses", testIPAddressesSliceDeleteAll)
	t.Run("Latencies", testLatenciesSliceDeleteAll)
	t.Run("MultiAddresses", testMultiAddressesSliceDeleteAll)
	t.Run("MultiAddressesSets", testMultiAddressesSetsSliceDeleteAll)
	t.Run("MultiAddressesXIPAddresses", testMultiAddressesXIPAddressesSliceDeleteAll)
	t.Run("Neighbors", testNeighborsSliceDeleteAll)
	t.Run("Peers", testPeersSliceDeleteAll)
	t.Run("PegasysConnections", testPegasysConnectionsSliceDeleteAll)
	t.Run("PegasysNeighbours", testPegasysNeighboursSliceDeleteAll)
	t.Run("Protocols", testProtocolsSliceDeleteAll)
	t.Run("ProtocolsSets", testProtocolsSetsSliceDeleteAll)
	t.Run("RawVisits", testRawVisitsSliceDeleteAll)
	t.Run("Sessions", testSessionsSliceDeleteAll)
	t.Run("Visits", testVisitsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("AgentVersions", testAgentVersionsExists)
	t.Run("CrawlProperties", testCrawlPropertiesExists)
	t.Run("Crawls", testCrawlsExists)
	t.Run("IPAddresses", testIPAddressesExists)
	t.Run("Latencies", testLatenciesExists)
	t.Run("MultiAddresses", testMultiAddressesExists)
	t.Run("MultiAddressesSets", testMultiAddressesSetsExists)
	t.Run("MultiAddressesXIPAddresses", testMultiAddressesXIPAddressesExists)
	t.Run("Neighbors", testNeighborsExists)
	t.Run("Peers", testPeersExists)
	t.Run("PegasysConnections", testPegasysConnectionsExists)
	t.Run("PegasysNeighbours", testPegasysNeighboursExists)
	t.Run("Protocols", testProtocolsExists)
	t.Run("ProtocolsSets", testProtocolsSetsExists)
	t.Run("RawVisits", testRawVisitsExists)
	t.Run("Sessions", testSessionsExists)
	t.Run("Visits", testVisitsExists)
}

func TestFind(t *testing.T) {
	t.Run("AgentVersions", testAgentVersionsFind)
	t.Run("CrawlProperties", testCrawlPropertiesFind)
	t.Run("Crawls", testCrawlsFind)
	t.Run("IPAddresses", testIPAddressesFind)
	t.Run("Latencies", testLatenciesFind)
	t.Run("MultiAddresses", testMultiAddressesFind)
	t.Run("MultiAddressesSets", testMultiAddressesSetsFind)
	t.Run("MultiAddressesXIPAddresses", testMultiAddressesXIPAddressesFind)
	t.Run("Neighbors", testNeighborsFind)
	t.Run("Peers", testPeersFind)
	t.Run("PegasysConnections", testPegasysConnectionsFind)
	t.Run("PegasysNeighbours", testPegasysNeighboursFind)
	t.Run("Protocols", testProtocolsFind)
	t.Run("ProtocolsSets", testProtocolsSetsFind)
	t.Run("RawVisits", testRawVisitsFind)
	t.Run("Sessions", testSessionsFind)
	t.Run("Visits", testVisitsFind)
}

func TestBind(t *testing.T) {
	t.Run("AgentVersions", testAgentVersionsBind)
	t.Run("CrawlProperties", testCrawlPropertiesBind)
	t.Run("Crawls", testCrawlsBind)
	t.Run("IPAddresses", testIPAddressesBind)
	t.Run("Latencies", testLatenciesBind)
	t.Run("MultiAddresses", testMultiAddressesBind)
	t.Run("MultiAddressesSets", testMultiAddressesSetsBind)
	t.Run("MultiAddressesXIPAddresses", testMultiAddressesXIPAddressesBind)
	t.Run("Neighbors", testNeighborsBind)
	t.Run("Peers", testPeersBind)
	t.Run("PegasysConnections", testPegasysConnectionsBind)
	t.Run("PegasysNeighbours", testPegasysNeighboursBind)
	t.Run("Protocols", testProtocolsBind)
	t.Run("ProtocolsSets", testProtocolsSetsBind)
	t.Run("RawVisits", testRawVisitsBind)
	t.Run("Sessions", testSessionsBind)
	t.Run("Visits", testVisitsBind)
}

func TestOne(t *testing.T) {
	t.Run("AgentVersions", testAgentVersionsOne)
	t.Run("CrawlProperties", testCrawlPropertiesOne)
	t.Run("Crawls", testCrawlsOne)
	t.Run("IPAddresses", testIPAddressesOne)
	t.Run("Latencies", testLatenciesOne)
	t.Run("MultiAddresses", testMultiAddressesOne)
	t.Run("MultiAddressesSets", testMultiAddressesSetsOne)
	t.Run("MultiAddressesXIPAddresses", testMultiAddressesXIPAddressesOne)
	t.Run("Neighbors", testNeighborsOne)
	t.Run("Peers", testPeersOne)
	t.Run("PegasysConnections", testPegasysConnectionsOne)
	t.Run("PegasysNeighbours", testPegasysNeighboursOne)
	t.Run("Protocols", testProtocolsOne)
	t.Run("ProtocolsSets", testProtocolsSetsOne)
	t.Run("RawVisits", testRawVisitsOne)
	t.Run("Sessions", testSessionsOne)
	t.Run("Visits", testVisitsOne)
}

func TestAll(t *testing.T) {
	t.Run("AgentVersions", testAgentVersionsAll)
	t.Run("CrawlProperties", testCrawlPropertiesAll)
	t.Run("Crawls", testCrawlsAll)
	t.Run("IPAddresses", testIPAddressesAll)
	t.Run("Latencies", testLatenciesAll)
	t.Run("MultiAddresses", testMultiAddressesAll)
	t.Run("MultiAddressesSets", testMultiAddressesSetsAll)
	t.Run("MultiAddressesXIPAddresses", testMultiAddressesXIPAddressesAll)
	t.Run("Neighbors", testNeighborsAll)
	t.Run("Peers", testPeersAll)
	t.Run("PegasysConnections", testPegasysConnectionsAll)
	t.Run("PegasysNeighbours", testPegasysNeighboursAll)
	t.Run("Protocols", testProtocolsAll)
	t.Run("ProtocolsSets", testProtocolsSetsAll)
	t.Run("RawVisits", testRawVisitsAll)
	t.Run("Sessions", testSessionsAll)
	t.Run("Visits", testVisitsAll)
}

func TestCount(t *testing.T) {
	t.Run("AgentVersions", testAgentVersionsCount)
	t.Run("CrawlProperties", testCrawlPropertiesCount)
	t.Run("Crawls", testCrawlsCount)
	t.Run("IPAddresses", testIPAddressesCount)
	t.Run("Latencies", testLatenciesCount)
	t.Run("MultiAddresses", testMultiAddressesCount)
	t.Run("MultiAddressesSets", testMultiAddressesSetsCount)
	t.Run("MultiAddressesXIPAddresses", testMultiAddressesXIPAddressesCount)
	t.Run("Neighbors", testNeighborsCount)
	t.Run("Peers", testPeersCount)
	t.Run("PegasysConnections", testPegasysConnectionsCount)
	t.Run("PegasysNeighbours", testPegasysNeighboursCount)
	t.Run("Protocols", testProtocolsCount)
	t.Run("ProtocolsSets", testProtocolsSetsCount)
	t.Run("RawVisits", testRawVisitsCount)
	t.Run("Sessions", testSessionsCount)
	t.Run("Visits", testVisitsCount)
}

func TestHooks(t *testing.T) {
	t.Run("AgentVersions", testAgentVersionsHooks)
	t.Run("CrawlProperties", testCrawlPropertiesHooks)
	t.Run("Crawls", testCrawlsHooks)
	t.Run("IPAddresses", testIPAddressesHooks)
	t.Run("Latencies", testLatenciesHooks)
	t.Run("MultiAddresses", testMultiAddressesHooks)
	t.Run("MultiAddressesSets", testMultiAddressesSetsHooks)
	t.Run("MultiAddressesXIPAddresses", testMultiAddressesXIPAddressesHooks)
	t.Run("Neighbors", testNeighborsHooks)
	t.Run("Peers", testPeersHooks)
	t.Run("PegasysConnections", testPegasysConnectionsHooks)
	t.Run("PegasysNeighbours", testPegasysNeighboursHooks)
	t.Run("Protocols", testProtocolsHooks)
	t.Run("ProtocolsSets", testProtocolsSetsHooks)
	t.Run("RawVisits", testRawVisitsHooks)
	t.Run("Sessions", testSessionsHooks)
	t.Run("Visits", testVisitsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("AgentVersions", testAgentVersionsInsert)
	t.Run("AgentVersions", testAgentVersionsInsertWhitelist)
	t.Run("CrawlProperties", testCrawlPropertiesInsert)
	t.Run("CrawlProperties", testCrawlPropertiesInsertWhitelist)
	t.Run("Crawls", testCrawlsInsert)
	t.Run("Crawls", testCrawlsInsertWhitelist)
	t.Run("IPAddresses", testIPAddressesInsert)
	t.Run("IPAddresses", testIPAddressesInsertWhitelist)
	t.Run("Latencies", testLatenciesInsert)
	t.Run("Latencies", testLatenciesInsertWhitelist)
	t.Run("MultiAddresses", testMultiAddressesInsert)
	t.Run("MultiAddresses", testMultiAddressesInsertWhitelist)
	t.Run("MultiAddressesSets", testMultiAddressesSetsInsert)
	t.Run("MultiAddressesSets", testMultiAddressesSetsInsertWhitelist)
	t.Run("MultiAddressesXIPAddresses", testMultiAddressesXIPAddressesInsert)
	t.Run("MultiAddressesXIPAddresses", testMultiAddressesXIPAddressesInsertWhitelist)
	t.Run("Neighbors", testNeighborsInsert)
	t.Run("Neighbors", testNeighborsInsertWhitelist)
	t.Run("Peers", testPeersInsert)
	t.Run("Peers", testPeersInsertWhitelist)
	t.Run("PegasysConnections", testPegasysConnectionsInsert)
	t.Run("PegasysConnections", testPegasysConnectionsInsertWhitelist)
	t.Run("PegasysNeighbours", testPegasysNeighboursInsert)
	t.Run("PegasysNeighbours", testPegasysNeighboursInsertWhitelist)
	t.Run("Protocols", testProtocolsInsert)
	t.Run("Protocols", testProtocolsInsertWhitelist)
	t.Run("ProtocolsSets", testProtocolsSetsInsert)
	t.Run("ProtocolsSets", testProtocolsSetsInsertWhitelist)
	t.Run("RawVisits", testRawVisitsInsert)
	t.Run("RawVisits", testRawVisitsInsertWhitelist)
	t.Run("Sessions", testSessionsInsert)
	t.Run("Sessions", testSessionsInsertWhitelist)
	t.Run("Visits", testVisitsInsert)
	t.Run("Visits", testVisitsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("CrawlPropertyToAgentVersionUsingAgentVersion", testCrawlPropertyToOneAgentVersionUsingAgentVersion)
	t.Run("CrawlPropertyToCrawlUsingCrawl", testCrawlPropertyToOneCrawlUsingCrawl)
	t.Run("CrawlPropertyToProtocolUsingProtocol", testCrawlPropertyToOneProtocolUsingProtocol)
	t.Run("LatencyToPeerUsingPeer", testLatencyToOnePeerUsingPeer)
	t.Run("MultiAddressesXIPAddressToIPAddressUsingIPAddress", testMultiAddressesXIPAddressToOneIPAddressUsingIPAddress)
	t.Run("MultiAddressesXIPAddressToMultiAddressUsingMultiAddress", testMultiAddressesXIPAddressToOneMultiAddressUsingMultiAddress)
	t.Run("NeighborToCrawlUsingCrawl", testNeighborToOneCrawlUsingCrawl)
	t.Run("NeighborToPeerUsingPeer", testNeighborToOnePeerUsingPeer)
	t.Run("PeerToAgentVersionUsingAgentVersion", testPeerToOneAgentVersionUsingAgentVersion)
	t.Run("PeerToProtocolsSetUsingProtocolsSet", testPeerToOneProtocolsSetUsingProtocolsSet)
	t.Run("SessionToPeerUsingPeer", testSessionToOnePeerUsingPeer)
	t.Run("VisitToAgentVersionUsingAgentVersion", testVisitToOneAgentVersionUsingAgentVersion)
	t.Run("VisitToCrawlUsingCrawl", testVisitToOneCrawlUsingCrawl)
	t.Run("VisitToMultiAddressesSetUsingMultiAddressesSet", testVisitToOneMultiAddressesSetUsingMultiAddressesSet)
	t.Run("VisitToPeerUsingPeer", testVisitToOnePeerUsingPeer)
	t.Run("VisitToProtocolsSetUsingProtocolsSet", testVisitToOneProtocolsSetUsingProtocolsSet)
	t.Run("VisitToSessionUsingSession", testVisitToOneSessionUsingSession)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("AgentVersionToCrawlProperties", testAgentVersionToManyCrawlProperties)
	t.Run("AgentVersionToPeers", testAgentVersionToManyPeers)
	t.Run("AgentVersionToVisits", testAgentVersionToManyVisits)
	t.Run("CrawlToCrawlProperties", testCrawlToManyCrawlProperties)
	t.Run("CrawlToNeighbors", testCrawlToManyNeighbors)
	t.Run("CrawlToVisits", testCrawlToManyVisits)
	t.Run("IPAddressToMultiAddressesXIPAddresses", testIPAddressToManyMultiAddressesXIPAddresses)
	t.Run("MultiAddressToMultiAddressesXIPAddresses", testMultiAddressToManyMultiAddressesXIPAddresses)
	t.Run("MultiAddressToPeers", testMultiAddressToManyPeers)
	t.Run("MultiAddressesSetToVisits", testMultiAddressesSetToManyVisits)
	t.Run("PeerToLatencies", testPeerToManyLatencies)
	t.Run("PeerToNeighbors", testPeerToManyNeighbors)
	t.Run("PeerToMultiAddresses", testPeerToManyMultiAddresses)
	t.Run("PeerToSessions", testPeerToManySessions)
	t.Run("PeerToVisits", testPeerToManyVisits)
	t.Run("ProtocolToCrawlProperties", testProtocolToManyCrawlProperties)
	t.Run("ProtocolsSetToPeers", testProtocolsSetToManyPeers)
	t.Run("ProtocolsSetToVisits", testProtocolsSetToManyVisits)
	t.Run("SessionToVisits", testSessionToManyVisits)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("CrawlPropertyToAgentVersionUsingCrawlProperties", testCrawlPropertyToOneSetOpAgentVersionUsingAgentVersion)
	t.Run("CrawlPropertyToCrawlUsingCrawlProperties", testCrawlPropertyToOneSetOpCrawlUsingCrawl)
	t.Run("CrawlPropertyToProtocolUsingCrawlProperties", testCrawlPropertyToOneSetOpProtocolUsingProtocol)
	t.Run("LatencyToPeerUsingLatencies", testLatencyToOneSetOpPeerUsingPeer)
	t.Run("MultiAddressesXIPAddressToIPAddressUsingMultiAddressesXIPAddresses", testMultiAddressesXIPAddressToOneSetOpIPAddressUsingIPAddress)
	t.Run("MultiAddressesXIPAddressToMultiAddressUsingMultiAddressesXIPAddresses", testMultiAddressesXIPAddressToOneSetOpMultiAddressUsingMultiAddress)
	t.Run("NeighborToCrawlUsingNeighbors", testNeighborToOneSetOpCrawlUsingCrawl)
	t.Run("NeighborToPeerUsingNeighbors", testNeighborToOneSetOpPeerUsingPeer)
	t.Run("PeerToAgentVersionUsingPeers", testPeerToOneSetOpAgentVersionUsingAgentVersion)
	t.Run("PeerToProtocolsSetUsingPeers", testPeerToOneSetOpProtocolsSetUsingProtocolsSet)
	t.Run("SessionToPeerUsingSessions", testSessionToOneSetOpPeerUsingPeer)
	t.Run("VisitToAgentVersionUsingVisits", testVisitToOneSetOpAgentVersionUsingAgentVersion)
	t.Run("VisitToCrawlUsingVisits", testVisitToOneSetOpCrawlUsingCrawl)
	t.Run("VisitToMultiAddressesSetUsingVisits", testVisitToOneSetOpMultiAddressesSetUsingMultiAddressesSet)
	t.Run("VisitToPeerUsingVisits", testVisitToOneSetOpPeerUsingPeer)
	t.Run("VisitToProtocolsSetUsingVisits", testVisitToOneSetOpProtocolsSetUsingProtocolsSet)
	t.Run("VisitToSessionUsingVisits", testVisitToOneSetOpSessionUsingSession)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("CrawlPropertyToAgentVersionUsingCrawlProperties", testCrawlPropertyToOneRemoveOpAgentVersionUsingAgentVersion)
	t.Run("CrawlPropertyToProtocolUsingCrawlProperties", testCrawlPropertyToOneRemoveOpProtocolUsingProtocol)
	t.Run("PeerToAgentVersionUsingPeers", testPeerToOneRemoveOpAgentVersionUsingAgentVersion)
	t.Run("PeerToProtocolsSetUsingPeers", testPeerToOneRemoveOpProtocolsSetUsingProtocolsSet)
	t.Run("VisitToAgentVersionUsingVisits", testVisitToOneRemoveOpAgentVersionUsingAgentVersion)
	t.Run("VisitToCrawlUsingVisits", testVisitToOneRemoveOpCrawlUsingCrawl)
	t.Run("VisitToMultiAddressesSetUsingVisits", testVisitToOneRemoveOpMultiAddressesSetUsingMultiAddressesSet)
	t.Run("VisitToProtocolsSetUsingVisits", testVisitToOneRemoveOpProtocolsSetUsingProtocolsSet)
	t.Run("VisitToSessionUsingVisits", testVisitToOneRemoveOpSessionUsingSession)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("AgentVersionToCrawlProperties", testAgentVersionToManyAddOpCrawlProperties)
	t.Run("AgentVersionToPeers", testAgentVersionToManyAddOpPeers)
	t.Run("AgentVersionToVisits", testAgentVersionToManyAddOpVisits)
	t.Run("CrawlToCrawlProperties", testCrawlToManyAddOpCrawlProperties)
	t.Run("CrawlToNeighbors", testCrawlToManyAddOpNeighbors)
	t.Run("CrawlToVisits", testCrawlToManyAddOpVisits)
	t.Run("IPAddressToMultiAddressesXIPAddresses", testIPAddressToManyAddOpMultiAddressesXIPAddresses)
	t.Run("MultiAddressToMultiAddressesXIPAddresses", testMultiAddressToManyAddOpMultiAddressesXIPAddresses)
	t.Run("MultiAddressToPeers", testMultiAddressToManyAddOpPeers)
	t.Run("MultiAddressesSetToVisits", testMultiAddressesSetToManyAddOpVisits)
	t.Run("PeerToLatencies", testPeerToManyAddOpLatencies)
	t.Run("PeerToNeighbors", testPeerToManyAddOpNeighbors)
	t.Run("PeerToMultiAddresses", testPeerToManyAddOpMultiAddresses)
	t.Run("PeerToSessions", testPeerToManyAddOpSessions)
	t.Run("PeerToVisits", testPeerToManyAddOpVisits)
	t.Run("ProtocolToCrawlProperties", testProtocolToManyAddOpCrawlProperties)
	t.Run("ProtocolsSetToPeers", testProtocolsSetToManyAddOpPeers)
	t.Run("ProtocolsSetToVisits", testProtocolsSetToManyAddOpVisits)
	t.Run("SessionToVisits", testSessionToManyAddOpVisits)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("AgentVersionToCrawlProperties", testAgentVersionToManySetOpCrawlProperties)
	t.Run("AgentVersionToPeers", testAgentVersionToManySetOpPeers)
	t.Run("AgentVersionToVisits", testAgentVersionToManySetOpVisits)
	t.Run("CrawlToVisits", testCrawlToManySetOpVisits)
	t.Run("MultiAddressToPeers", testMultiAddressToManySetOpPeers)
	t.Run("MultiAddressesSetToVisits", testMultiAddressesSetToManySetOpVisits)
	t.Run("PeerToMultiAddresses", testPeerToManySetOpMultiAddresses)
	t.Run("ProtocolToCrawlProperties", testProtocolToManySetOpCrawlProperties)
	t.Run("ProtocolsSetToPeers", testProtocolsSetToManySetOpPeers)
	t.Run("ProtocolsSetToVisits", testProtocolsSetToManySetOpVisits)
	t.Run("SessionToVisits", testSessionToManySetOpVisits)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("AgentVersionToCrawlProperties", testAgentVersionToManyRemoveOpCrawlProperties)
	t.Run("AgentVersionToPeers", testAgentVersionToManyRemoveOpPeers)
	t.Run("AgentVersionToVisits", testAgentVersionToManyRemoveOpVisits)
	t.Run("CrawlToVisits", testCrawlToManyRemoveOpVisits)
	t.Run("MultiAddressToPeers", testMultiAddressToManyRemoveOpPeers)
	t.Run("MultiAddressesSetToVisits", testMultiAddressesSetToManyRemoveOpVisits)
	t.Run("PeerToMultiAddresses", testPeerToManyRemoveOpMultiAddresses)
	t.Run("ProtocolToCrawlProperties", testProtocolToManyRemoveOpCrawlProperties)
	t.Run("ProtocolsSetToPeers", testProtocolsSetToManyRemoveOpPeers)
	t.Run("ProtocolsSetToVisits", testProtocolsSetToManyRemoveOpVisits)
	t.Run("SessionToVisits", testSessionToManyRemoveOpVisits)
}

func TestReload(t *testing.T) {
	t.Run("AgentVersions", testAgentVersionsReload)
	t.Run("CrawlProperties", testCrawlPropertiesReload)
	t.Run("Crawls", testCrawlsReload)
	t.Run("IPAddresses", testIPAddressesReload)
	t.Run("Latencies", testLatenciesReload)
	t.Run("MultiAddresses", testMultiAddressesReload)
	t.Run("MultiAddressesSets", testMultiAddressesSetsReload)
	t.Run("MultiAddressesXIPAddresses", testMultiAddressesXIPAddressesReload)
	t.Run("Neighbors", testNeighborsReload)
	t.Run("Peers", testPeersReload)
	t.Run("PegasysConnections", testPegasysConnectionsReload)
	t.Run("PegasysNeighbours", testPegasysNeighboursReload)
	t.Run("Protocols", testProtocolsReload)
	t.Run("ProtocolsSets", testProtocolsSetsReload)
	t.Run("RawVisits", testRawVisitsReload)
	t.Run("Sessions", testSessionsReload)
	t.Run("Visits", testVisitsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("AgentVersions", testAgentVersionsReloadAll)
	t.Run("CrawlProperties", testCrawlPropertiesReloadAll)
	t.Run("Crawls", testCrawlsReloadAll)
	t.Run("IPAddresses", testIPAddressesReloadAll)
	t.Run("Latencies", testLatenciesReloadAll)
	t.Run("MultiAddresses", testMultiAddressesReloadAll)
	t.Run("MultiAddressesSets", testMultiAddressesSetsReloadAll)
	t.Run("MultiAddressesXIPAddresses", testMultiAddressesXIPAddressesReloadAll)
	t.Run("Neighbors", testNeighborsReloadAll)
	t.Run("Peers", testPeersReloadAll)
	t.Run("PegasysConnections", testPegasysConnectionsReloadAll)
	t.Run("PegasysNeighbours", testPegasysNeighboursReloadAll)
	t.Run("Protocols", testProtocolsReloadAll)
	t.Run("ProtocolsSets", testProtocolsSetsReloadAll)
	t.Run("RawVisits", testRawVisitsReloadAll)
	t.Run("Sessions", testSessionsReloadAll)
	t.Run("Visits", testVisitsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("AgentVersions", testAgentVersionsSelect)
	t.Run("CrawlProperties", testCrawlPropertiesSelect)
	t.Run("Crawls", testCrawlsSelect)
	t.Run("IPAddresses", testIPAddressesSelect)
	t.Run("Latencies", testLatenciesSelect)
	t.Run("MultiAddresses", testMultiAddressesSelect)
	t.Run("MultiAddressesSets", testMultiAddressesSetsSelect)
	t.Run("MultiAddressesXIPAddresses", testMultiAddressesXIPAddressesSelect)
	t.Run("Neighbors", testNeighborsSelect)
	t.Run("Peers", testPeersSelect)
	t.Run("PegasysConnections", testPegasysConnectionsSelect)
	t.Run("PegasysNeighbours", testPegasysNeighboursSelect)
	t.Run("Protocols", testProtocolsSelect)
	t.Run("ProtocolsSets", testProtocolsSetsSelect)
	t.Run("RawVisits", testRawVisitsSelect)
	t.Run("Sessions", testSessionsSelect)
	t.Run("Visits", testVisitsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("AgentVersions", testAgentVersionsUpdate)
	t.Run("CrawlProperties", testCrawlPropertiesUpdate)
	t.Run("Crawls", testCrawlsUpdate)
	t.Run("IPAddresses", testIPAddressesUpdate)
	t.Run("Latencies", testLatenciesUpdate)
	t.Run("MultiAddresses", testMultiAddressesUpdate)
	t.Run("MultiAddressesSets", testMultiAddressesSetsUpdate)
	t.Run("MultiAddressesXIPAddresses", testMultiAddressesXIPAddressesUpdate)
	t.Run("Neighbors", testNeighborsUpdate)
	t.Run("Peers", testPeersUpdate)
	t.Run("PegasysConnections", testPegasysConnectionsUpdate)
	t.Run("PegasysNeighbours", testPegasysNeighboursUpdate)
	t.Run("Protocols", testProtocolsUpdate)
	t.Run("ProtocolsSets", testProtocolsSetsUpdate)
	t.Run("RawVisits", testRawVisitsUpdate)
	t.Run("Sessions", testSessionsUpdate)
	t.Run("Visits", testVisitsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("AgentVersions", testAgentVersionsSliceUpdateAll)
	t.Run("CrawlProperties", testCrawlPropertiesSliceUpdateAll)
	t.Run("Crawls", testCrawlsSliceUpdateAll)
	t.Run("IPAddresses", testIPAddressesSliceUpdateAll)
	t.Run("Latencies", testLatenciesSliceUpdateAll)
	t.Run("MultiAddresses", testMultiAddressesSliceUpdateAll)
	t.Run("MultiAddressesSets", testMultiAddressesSetsSliceUpdateAll)
	t.Run("MultiAddressesXIPAddresses", testMultiAddressesXIPAddressesSliceUpdateAll)
	t.Run("Neighbors", testNeighborsSliceUpdateAll)
	t.Run("Peers", testPeersSliceUpdateAll)
	t.Run("PegasysConnections", testPegasysConnectionsSliceUpdateAll)
	t.Run("PegasysNeighbours", testPegasysNeighboursSliceUpdateAll)
	t.Run("Protocols", testProtocolsSliceUpdateAll)
	t.Run("ProtocolsSets", testProtocolsSetsSliceUpdateAll)
	t.Run("RawVisits", testRawVisitsSliceUpdateAll)
	t.Run("Sessions", testSessionsSliceUpdateAll)
	t.Run("Visits", testVisitsSliceUpdateAll)
}
