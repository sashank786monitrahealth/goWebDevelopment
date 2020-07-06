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
	request(conn)
}

func request(conn net.Conn) {
	var i int = 0
	var scanner *bufio.Scanner = bufio.NewScanner(conn)

	for scanner.Scan() {
		var ln string = scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}

		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	// request line
	var m string = strings.Fields(ln)[0] // method
	var u string = strings.Fields(ln)[1] // URI
	fmt.Println("***METHOD", m)
	fmt.Println("***URI", u)

	//multiplexer
	if m == "GET" && u == "/" {
		index(conn)
	}

	if m == "GET" && u == "/about" {
		about(conn)
	}

	if m == "GET" && u == "/contact" {
		contact(conn)
	}

	if m == "GET" && u == "/apply" {
		apply(conn)
	}

	if m == "POST" && u == "/apply" {
		applyProcess(conn)
	}

}

func index(conn net.Conn) {
	var body string = `<!DOCTYPE html> <html lang="en"><head><meta charset="UTF-8"><title></title></head>
	<body>
	<strong>INDEX</strong><br>
	<a href="/"></a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>

	</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func about(conn net.Conn) {
	var body string = `<!DOCTYPE html> <html lang="en"><head><meta charset="UTF-8"><title></title></head>
	<body>
	<strong>ABOUT</strong><br>
	<a href="/"></a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>

	</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func contact(conn net.Conn) {
	var body string = `<!DOCTYPE html> <html lang="en"><head><meta charset="UTF-8"><title></title></head>
	<body>
	<strong>CONTACT</strong><br>
	<a href="/"></a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>

	</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func apply(conn net.Conn) {
	var body string = `<!DOCTYPE html> <html lang="en"><head><meta charset="UTF-8"><title></title></head>
	<body>
	<strong>APPLY</strong><br>
	<a href="/"></a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	<form method="POST" action="/apply">
	<input type="submit" value="apply">
	</form>
	
	</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func applyProcess(conn net.Conn) {
	var body string = `<!DOCTYPE html> <html lang="en"><head><meta charset="UTF-8"><title></title></head>
	<body>
	<strong>APPLY PROCESS</strong><br>
	<a href="/"></a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>

	</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
