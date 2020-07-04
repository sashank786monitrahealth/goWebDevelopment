package main

import (
	"io"
	"log"
	"net"
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
}
