package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	var l net.Listener
	var err error

	l, err = net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer l.Close()

	for {
		var conn net.Conn
		var err error
		conn, err = l.Accept()

		if err != nil {
			log.Println(err)
		}

		var scanner *bufio.Scanner = bufio.NewScanner(conn)
		for scanner.Scan() {
			var ln string = scanner.Text()
			fmt.Println(ln)
		}
		defer conn.Close()

		// we never get here
		// we have an open stream connection
		// how does the above reader know when it is done?
		fmt.Println("Code got here")
		io.WriteString(conn, "I see you connected.")

		conn.Close()

	}
}
