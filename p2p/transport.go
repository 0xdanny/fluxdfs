package p2p

// This represents a remote node
type Peer interface {
}

// This handles the communication between the nodes
// in the network and it can be either TCP, UDP,
// websockets, etc.
type Transport interface {
	ListenAndAccept() error
}
