package main

import (
	"log"

	"github.com/0xdanny/fluxdfs/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3200")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}

	// fmt.Println("Flux DFS")
}
