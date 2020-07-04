package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	var err error
	var li net.Listener

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

	var scanner *bufio.Scanner = bufio.NewScanner(conn)

	for scanner.Scan() {
		var ln string = strings.ToLower(scanner.Text())
		var bs []byte = []byte(ln)
		var r []byte = rot13(bs)

		fmt.Fprintf(conn, "%s - %s\n\n", ln, r)
	}
}

func rot13(bs []byte) []byte {
	var r13 []byte = make([]byte, len(bs))

	for i, v := range bs {
		// ascii 97 - 122
		if v <= 109 {
			r13[1] = v + 13
		} else {
			r13[i] = v - 13
		}
	}

	return r13
}
