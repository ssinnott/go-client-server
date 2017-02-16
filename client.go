package main

import (
	"bufio"
	"fmt"
	"os"
	"net"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	conn, err := net.Dial("tcp", CONN_HOST + ":" + CONN_PORT)
	if err != nil {
		fmt.Print("Error: count not connect to " + CONN_HOST + ":" + CONN_PORT)
	}

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print("Error: " + line)
			os.Exit(1)
		}

	}
	reader.ReadString('\n')
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	fmt.Println("Enter text: ")
	text2 := ""
	fmt.Scanln(text2)
	fmt.Println(text2)

	ln := ""
	fmt.Sscanln("%v", ln)
	fmt.Println(ln)
}