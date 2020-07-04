package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	var li net.Listener
	var err error
	li, err = net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalln(err)
	}

	defer li.Close()

	for {
		var conn net.Conn
		conn, err = li.Accept()

		if err != nil {
			log.Println(err)
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	var err error = conn.SetDeadline(time.Now().Add(10 * time.Second))

	if err != nil {
		log.Fatalln("CONN TIMEOUT")
	}

	var scanner *bufio.Scanner = bufio.NewScanner(conn)

	for scanner.Scan() {
		var ln string = scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	defer conn.Close()

	// now we get here
	// the connection will time out
	// that breaks us out of the scanner loop
	fmt.Println("***CODE GOT HERE***")
}

/*
Run it using:
telnet localhost 8080
*/
