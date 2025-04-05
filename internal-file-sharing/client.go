package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("error: %s", err.Error());
	}
	defer conn.Close();
	s := "Hello TCP /n"
	conn.Write([]byte(s))
}