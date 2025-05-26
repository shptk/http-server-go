package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const CRLF = "\r\n"

func main() {
	fmt.Println("Starting the server....")
	l, err := net.Listen("tcp", "0.0.0.0:1234")
	if err != nil {
		fmt.Println("Error while setting up the tcp server: ", err)
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error while accepting connection: ", err)
			os.Exit(1)
		}

		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error while reading buffer: ", err)
		}
		res := ""
		if n > 0 {
			req := string(buffer[:n])
			if strings.Contains(req, "/echo") {
				path := strings.Split(strings.Split(req, CRLF)[0][len("GET /echo"):], " ")[0]
				fmt.Println("printing path: ", path)
				res = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(path), path)

			} else if strings.HasPrefix(req, "GET / HTTP/1.1") {
				res = "HTTP/1.1 200 OK\r\n\r\n"
			} else {
				res = "HTTP/1.1 404 Not Found\r\n\r\n"
			}

		}
		res_code, err := conn.Write([]byte(res))
		if err != nil {
			fmt.Println("Error while sending response: ", err)
			os.Exit(1)
		}
		fmt.Println("Response return code: ", res_code)
		conn.Close()
	}
}
