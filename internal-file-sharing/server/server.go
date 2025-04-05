package server

import (
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}
	if n == 0 {
		fmt.Printf("Req is empty: %d", n)
		return
	}
	defer conn.Close()
	fmt.Printf("Connection: %s, Body: %s", conn.LocalAddr().String(), string(buffer))
}

func Run() error {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return fmt.Errorf("error starting server: %s", err.Error())
	}
	log.Printf("server started at %s", ":8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			return fmt.Errorf("error connecting failed: %s", err.Error())
		}

		go handleConnection(conn)
	}
	return nil
}
