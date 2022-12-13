package dht

type Peer struct {
	ID []byte
}

type Bucket struct {
	Peers []Peer
}

// RoutingTable defines the routing table
type RoutingTable struct {
	Buckets []Bucket
}
