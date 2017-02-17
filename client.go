package main

import (
	"bufio"
	"fmt"
	"os"
	"net"
	"io"
	"bytes"
)


const (
	CONN_HOST = "localhost"
	CONN_PORT = "3334"
	CONN_TYPE = "tcp"
	BUFFER_LENGTH = 1024
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	conn, err := net.Dial("tcp", CONN_HOST + ":" + CONN_PORT)
	if err != nil {
		fmt.Print("Error: count not connect to " + CONN_HOST + ":" + CONN_PORT)
	}

	defer conn.Close()

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print("Error: " + line)
			os.Exit(1)
		}
		conn.Write([]byte(line))

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

		fmt.Print(buffer.String())

	}

}