/*
## Create a basic server using TCP.
The server should use net.Listen to listen on port 8080

Remember to close the listener using defer.

Remember that from the "net" package you first need to LISTEN, then you need to accept
an incoming connection.

Now write a response back on the connection.

Use io.WriteString to write the response. I see you connected.

Remember to close the connection.

Once you have all of that working, run your TCP server and from test it from telnet
(telnet localhost 8080)
*/

package main

import (
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
		var c net.Conn
		c, err = l.Accept()
		if err != nil {
			log.Println(err)
		}

		// write to connection
		io.WriteString(c, "I see you connected.")
		c.Close()
	}
}
