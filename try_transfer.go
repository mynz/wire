package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("errCheck", err)
		os.Exit(1)
	}
}

func doServer() {
	ta, err := net.ResolveTCPAddr("tcp", ":5555")
	checkError(err)
	listner, err := net.ListenTCP("tcp4", ta)
	checkError(err)

	for {
		conn, err := listner.Accept()
		checkError(err)

		go func(conn net.Conn) {
			defer conn.Close()

			buf := make([]byte, 32)
			n, err := conn.Read(buf)
			checkError(err)
			fmt.Println("Accepted conn:", n, string(buf))
		}(conn)
	}
}

func doClient() {
	ta, err := net.ResolveTCPAddr("tcp", "localhost:5555")
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, ta)
	checkError(err)
	defer conn.Close()

	conn.Write([]byte("hello"))
}

func main() {
	isServer := flag.Bool("server", false, "Wheathr this app runs as a server")
	flag.Parse()

	if *isServer {
		fmt.Println("runs as server")
		doServer()
	} else {
		fmt.Println("runs as client")
		doClient()
	}
}
