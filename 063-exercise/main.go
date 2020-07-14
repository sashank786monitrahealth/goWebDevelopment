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

		go serve(conn)

	}
}

func serve(conn net.Conn) {

	defer conn.Close()

	var scanner *bufio.Scanner = bufio.NewScanner(conn)
	for scanner.Scan() {
		var ln string = scanner.Text()
		fmt.Println(ln)

		if ln == "" {
			// when ln is empty, header is done
			fmt.Println("This is the end of the http request headers.")
			break
		}
	}

	var body string = "CHECK OUT THE RESPONSE BODY PAYLOAD"

	// we never get here
	// we have an open stream connection
	// how does the above reader know when it is done?
	fmt.Println("Code got here")
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)

	//conn.Close()

}
