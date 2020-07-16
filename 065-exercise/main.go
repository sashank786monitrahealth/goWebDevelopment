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
	var i int64
	var rMethod, rURI string
	var scanner *bufio.Scanner = bufio.NewScanner(conn)
	for scanner.Scan() {
		var ln string = scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			// we are in REQUEST LINE
			var xs []string = strings.Fields(ln)
			rMethod = xs[0]
			rURI = xs[1]
			fmt.Println("METHOD:", rMethod)
			fmt.Println("URI:", rURI)
		}

		if ln == "" {
			// when ln is empty, header is done
			fmt.Println("This is the end of the http request headers.")
			break
		}
		i++
	}

	switch {
	case rMethod == "GET" && rURI == "/":
		handleIndex(conn)
	case rMethod == "GET" && rURI == "/apply":
		handleApply(conn)
	case rMethod == "POST" && rURI == "":
		handleApplyPost(conn)
	default:
		handleDefault(conn)
	}

}

func handleIndex(conn net.Conn) {
	var body string = `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset = "UTF-8">
	<title></title>

	</head>
	<body>
	<h1>"GET INDEX"</h1>
	<a href="/">index</a>
	<a href="/apply">apply</a>
	</body>
	</html>
	`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)

	//conn.Close()
}

func handleApply(conn net.Conn) {
	var body string = `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<meta http-equiv="X-UA-Compatible" content="ie=edge">
			<title>GET DOG</title>
		</head>
		<body>
			<h1>"GET APPLY"</h1>
			<a href="/">index</a> <br>
			<a href="/apply">apply</a><br>
			<form action="/apply" method="POST">
			<input type="hidden" value="In my good death">
			<input type="submit" value="submit">
			</form>
		</body>
		</html>

		`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)

}

func handleApplyPost(conn net.Conn) {
	var body string = `
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>POST APPLY</title>
</head>
<body>
    <h1>"POST APPLY"</h1>
    <a href="/">index</a> <br>
    <a href="/apply">apply</a> <br>
</body>
</html>
	`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func handleDefault(conn net.Conn) {
	var body string = `
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>default</title>
</head>
<body>
    <h1>"default"</h1>
</body>
</html>
	`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
