package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	// Listen on TCP port 8080
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer ln.Close()

	// Accept connection on port
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		defer conn.Close()

		// Handle the connection
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	// Create a new scanner and read the first line of the request
	i := 0
	scanner := bufio.NewScanner(conn)

	// Print the request
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			mux(conn, line)
		}
		if line == "" {
			break
		}
		fmt.Println(line)

		i++
	}
}

func mux(conn net.Conn, line string) {
	// request line
	m := strings.Fields(line)[0] // method
	u := strings.Fields(line)[1] // uri
	fmt.Println("***METHOD", m)
	fmt.Println("***URI", u)

	// multiplexer
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
	body := `<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charset="UTF-8">
			<title>Index</title>
			</head>
			<body>
			<h1>Index</h1>
			<a href="/about">About</a><br>
			<a href="/contact">Contact</a><br>
			<a href="/apply">Apply</a><br>
			</body>
			</html>
			`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "Content-Length: ", len(body), "\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func about(conn net.Conn) {
	body := `<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charset="UTF-8">
			<title>About</title>
			</head>
			<body>
			<h1>About</h1>
			<a href="/">Index</a><br>
			<a href="/contact">Contact</a><br>
			<a href="/apply">Apply</a><br>
			</body>
			</html>
			`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "Content-Length: ", len(body), "\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func contact(conn net.Conn) {
	body := `<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charset="UTF-8">
			<title>Contact</title>
			</head>
			<body>
			<h1>Contact</h1>
			<a href="/">Index</a><br>
			<a href="/about">About</a><br>
			<a href="/apply">Apply</a><br>
			</body>
			</html>
			`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "Content-Length: ", len(body), "\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func apply(conn net.Conn) {
	body := `<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charset="UTF-8">
			<title>Apply</title>
			</head>
			<body>
			<h1>Apply</h1>
			<a href="/">Index</a><br>
			<a href="/about">About</a><br>
			<a href="/contact">Contact</a><br>
			<form method="post" action="/apply">
			<input type="submit" value="Apply">
			</form>
			</body>
			</html>
			`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "Content-Length: ", len(body), "\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func applyProcess(conn net.Conn) {
	body := `<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charset="UTF-8">
			<title>Apply Process</title>
			</head>
			<body>
			<h1>Apply Process</h1>
			<a href="/">Index</a><br>
			<a href="/about">About</a><br>
			<a href="/contact">Contact</a><br>
			<a href="/apply">Apply</a><br>
			</body>
			</html>
			`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "Content-Length: ", len(body), "\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
