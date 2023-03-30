package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := sequentialRouteTo("Binzmühle")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)

	start = time.Now()
	results = concurrentRouteTo("Binzmühle")
	elapsed = time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)

	start = time.Now()
	results = replicatedRouteTo("Binzmühle")
	elapsed = time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

var publicTransport = fakeRoute("publicTransport")
var car = fakeRoute("car")
var bike = fakeRoute("bike")

type Route func(destination string) string

func fakeRoute(mode string) Route {
	return func(destination string) string {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		return fmt.Sprintf("quickest %s route to %q: ...\n", mode, destination)
	}
}

func sequentialRouteTo(destination string) (routes []string) {
	routes = append(routes, publicTransport(destination))
	routes = append(routes, car(destination))
	routes = append(routes, bike(destination))
	return
}

func concurrentRouteTo(destination string) (routes []string) {
	c := make (chan string)
	go func() { c <- publicTransport(destination) } ()
	go func() { c <- car(destination) } () 
	go func() { c <- bike(destination) } ()

	for i := 0; i < 3; i++ {
		route := <-c
		routes = append(routes, route)
	}
	// timeout := time.After(120 * time.Millisecond)
	// for i := 0; i < 3; i++ {
	// 	select {
	// 	case route := <-c:
	// 		routes = append(routes, route)
	// 	case <-timeout:
	// 		fmt.Println("timed out")
	// 		return
	// 	}
	// }
	return
}

func replicatedRouteTo(destination string) (routes []string) {
	// same as before, but use multiple replicas to decrease the chance of a timeout
	return
}

func First(destination string, replicas ...Route) string {
	// create a channel that handles multiple replicas of each transportation system and returns only the fastest one
	return "not implemented"
}