package main

import (
	"fmt"
	"io/ioutil"
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

	var bs []byte
	bs, err = ioutil.ReadAll(conn)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))

}
