package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"

	"github.com/tkblack/uos-go/p2p"
)

var peer1 = flag.String("peer1", "localhost:9876", "peer 1")
var peer2 = flag.String("peer2", "localhost:2222", "peer 2")
var chainID = flag.String("chain-id", "308cae83a690640be3726a725dde1fa72a845e28cfc63f28c3fa0a6ccdb6faf0", "peer 1")
var showLog = flag.Bool("v", false, "show detail log")

func main() {
	flag.Parse()

	fmt.Println("P2P Proxy")

	if *showLog {
		p2p.EnableP2PLogging()
	}
	defer p2p.SyncLogger()

	cID, err := hex.DecodeString(*chainID)
	if err != nil {
		log.Fatal(err)
	}

	proxy := p2p.NewProxy(
		p2p.NewOutgoingPeer(*peer1, "uos-proxy", nil),
		p2p.NewOutgoingPeer(*peer2, "uos-proxy", &p2p.HandshakeInfo{
			ChainID: cID,
		}),
	)

	//proxy := p2p.NewProxy(
	//	p2p.NewOutgoingPeer("localhost:9876", chainID, "uos-proxy", false),
	//	p2p.NewIncommingPeer("localhost:1111", chainID, "uos-proxy"),
	//)

	//proxy := p2p.NewProxy(
	//	p2p.NewIncommingPeer("localhost:2222", "uos-proxy"),
	//	p2p.NewIncommingPeer("localhost:1111", "uos-proxy"),
	//)

	proxy.RegisterHandler(p2p.StringLoggerHandler)
	fmt.Println(proxy.ConnectAndStart())

}
