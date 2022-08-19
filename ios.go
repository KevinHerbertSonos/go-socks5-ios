package socks

import (
	"fmt"
	"log"
	"os"

	"github.com/things-go/go-socks5"
)

func startServer() {
	fmt.Printf("gomobile: In startServer goroutine\n")
	// Create a SOCKS5 server
	server := socks5.NewServer(
		socks5.WithLogger(socks5.NewLogger(log.New(os.Stdout, "socks5: ", log.LstdFlags))),
	)

	// Create SOCKS5 proxy on localhost port 10800
	if err := server.ListenAndServe("tcp", ":10800"); err != nil {
		panic(err)
	}
}

func StartServer() {
	fmt.Printf("gomobile: Made it to StartServer\n")
	go startServer()
}
