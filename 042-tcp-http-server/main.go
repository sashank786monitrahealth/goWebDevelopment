package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	var li net.Listener
	var err error

	li, err = net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalln(err.Error())
	}

	defer li.Close()

	for {
		var conn net.Conn

		conn, err = li.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	request(conn)

	// write response
	respond(conn)
}

func request(conn net.Conn) {
	var i int = 0
	var scanner *bufio.Scanner = bufio.NewScanner(conn)

	for scanner.Scan() {
		var ln string = scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			// request line
			var m string = strings.Fields(ln)[0] // method
			var u string = strings.Fields(ln)[0] // uri
			fmt.Println("***METHOD", m)
			fmt.Println("***URI", u)
		}

		if ln == "" {
			// headers are done
			break
		}

		i++
	}
}

func respond(conn net.Conn) {
	var body string = `<!DOCTYPE html> <html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
