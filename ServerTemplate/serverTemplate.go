package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

type Route func(destination string) string

func fakeRoute(mode string) Route {
	return func(destination string) string {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		return fmt.Sprintf("quickest %s route to %q: ...\n", mode, destination)
	}
}

func concurrentRouteTo(destination string) (routes []string) {
	c := make(chan string)
	go func() { c <- fakeRoute("publicTransport")(destination) }()
	go func() { c <- fakeRoute("car")(destination) }()
	go func() { c <- fakeRoute("bike")(destination) }()

	for i := 0; i < 3; i++ {
		route := <-c
		routes = append(routes, route)
	}
	return
}

func calculateWithRoute(mode string, destination string) string {
	r := fakeRoute(mode)
	erg := r(destination)
	return erg
}

func calculateAllRoutes(destination string) string {
	routes := concurrentRouteTo(destination)
	erg := ""
	for index, value := range routes {
		if index == len(routes)-1 {
			erg += value
		} else {
			erg += strings.TrimSuffix(value, "\n") + ", "
		}
	}
	return erg
}

func handleConnection(c net.Conn) {
	fmt.Printf("Starting new Connection with %s\n", c.RemoteAddr())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Printf("Closing connection with %s, because of error:\n\t%s\n", c.RemoteAddr(), err)
			return
		}

		fmt.Println(strings.TrimSpace(string(netData)))

		args := strings.Fields(netData)

		if args[0] == "STOP" {
			break
		} else if args[0] == "car" || args[0] == "bike" || args[0] == "publicTransport" {
			//send result of calculateWithRoute(args[0], args[1]) to client

		} else {
			//send result of calculateAllRoutes(args[0]) to client

		}
	}
	fmt.Printf("Closing connection with %s\n", c.RemoteAddr())
	c.Close()
}

func main() {
	// read the port the server should be using from standard input, check if port provided
	arguments := os.Args

	PORT := ":" + arguments[1]
	//listen for incoming connections

	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	//accept and handle single connection

	if err != nil {
		fmt.Println(err)
		return
	}
	handleConnection(c)
}
