package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	fmt.Println("Proxy Running")
	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		ClientConn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleProxyConnection(ClientConn)
	}
}

func handleProxyConnection(client net.Conn) {
	defer client.Close()
	ServerConn, err := net.Dial("tcp", "127.0.0.1:1162")
	if err != nil {
		panic(err)
	}
	defer ServerConn.Close()
	go func() {
		_, err := io.Copy(ServerConn, client)
		if err != nil {
			panic(err)
		}
	}()

	_, err = io.Copy(client, ServerConn)
	if err != nil {
		panic(err)
	}
}
