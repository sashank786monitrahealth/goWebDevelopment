package main

import (
	"bufio"
	"fmt"
	"io"
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
	defer conn.Close()

	// instructions
	var vWriteStr string = "\r\nIN-MEMORY DATABASE\r\n\r\n" +
		"USE:\r\n" +
		"\tSET key value \r\n" +
		"\tGET key \r\n" +
		"\tDEL key \r\n\r\n" +
		"EXAMPLE:\r\n" +
		"\tSET fav chocolate \r\n" +
		"\t GET fav \r\n\r\n\r\n"

	io.WriteString(conn, vWriteStr)

	// read & write
	var data map[string]string = make(map[string]string)
	var scanner *bufio.Scanner = bufio.NewScanner(conn)

	for scanner.Scan() {
		var ln string = scanner.Text()
		var fs []string = strings.Fields(ln)

		//logic
		if len(fs) < 1 {
			continue
		}

		switch fs[0] {
		case "GET":
			var k string = fs[1]
			var v string = data[k]
			fmt.Fprintf(conn, "%s\r\n", v)
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "EXPECTED VALUE\r\n")
				continue
			}
			var k string = fs[1]
			var v string = fs[2]
			data[k] = v
		case "DEL":
			var k string = fs[1]
			delete(data, k)
		default:
			fmt.Fprintln(conn, "INVALID COMMAND "+fs[0]+"\r\n")
			continue
		}
	}
}
