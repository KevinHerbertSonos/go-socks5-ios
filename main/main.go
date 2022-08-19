package main

import (
	"log"
	"os"

	"github.com/things-go/go-socks5"
)

func main() {
	// Create a SOCKS5 server
	server := socks5.NewServer(
		socks5.WithLogger(socks5.NewLogger(log.New(os.Stdout, "socks5: ", log.LstdFlags))),
	)

	// Create SOCKS5 proxy on localhost port 10800
	if err := server.ListenAndServe("tcp", ":10800"); err != nil {
		panic(err)
	}
}
