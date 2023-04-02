package main

import (
	"bufio"
	"fmt"

	//"net"
	"os"
	"strings"
)

func main() {
	//read command line arguments, provided: host:port
	arguments := os.Args

	CONNECT := arguments[1]
	//establish connection to specified host and port using the net package, handle error

	if err != nil {
		fmt.Println(err)
		return
	}

	//send requests until client wants to end connection
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		//send data to server using connection

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
