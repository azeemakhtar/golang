package main

import (
	"fmt"
	"time"
)

// Initializing a ping pong channel
func pinger(pinger <-chan int, ponger chan<- int) {
	for {
		<-pinger
		fmt.Println("Ping")
		time.Sleep(1 * time.Second)
		ponger <- 1
	}
}
func ponger(pinger chan<- int, ponger <-chan int) {
	for {
		<-ponger
		fmt.Println("Pong")
		time.Sleep(1 * time.Second)
		pinger <- 1
	}
}

func main() {
	//initiate a buffered channel
	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)

	//The main go rountine starts the gameby sending a value to the ping channel
	ping <- 1
	for {
		// Block the main thread until an intrupt
		time.Sleep(time.Second)
	}
}
