package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	var err error
	var conn net.Conn

	conn, err = net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	fmt.Fprintln(conn, "I dialed you.")
}
