package p2p

// This
type HandShakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
