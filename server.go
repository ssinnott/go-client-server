package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)


func main() {

	l, err := net.Listen(CONN_TYPE, fmt.Sprintf("%s:%s", CONN_HOST, CONN_PORT))
	if err != nil {
		fmt.Println(fmt.Sprintf(
			"ERROR: Failed to register port at %s:%s message is %s",
			CONN_HOST,
			CONN_PORT,
			err.Error()))
		os.Exit(1)
	}

	defer l.Close()
	fmt.Println(fmt.Sprintf("ECHO server running on %s:%s. Have fun!", CONN_HOST, CONN_PORT))
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("ERROR accepting: ", err.Error())
			os.Exit(1)
		}

		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	for {
		var buffer bytes.Buffer

		buf := make([]byte, BUFFER_LENGTH)
		for {
			length, err := conn.Read(buf)
			if err != nil && err != io.EOF{
				fmt.Println("Server read error: ", err)
				os.Exit(1)
			}else if err != io.EOF {
				buffer.Write(buf[0:length])
				if length != BUFFER_LENGTH - 1{
					break
				}
			}
		}

		var value = strings.TrimSpace(buffer.String())
		if value == "exit" {
			fmt.Fprint(conn, "Bye bye!!!\n")
			break
		}else{
			fmt.Fprint(conn, value + " ..... right back at you!\n")
		}

	}
	conn.Close()
}
