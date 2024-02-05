package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel. Declare a channel of type int
	// var name of channel chan keyword & data type
	//var c chan int
	//initialize the channel
	c := make((chan int))

	//send a value to the channel in a goroutine
	go func() {
		c <- 42
	}()
	//sniff the value from the channel
	val := <-c
	fmt.Println(val)
	time.Sleep(time.Second * 3)
	//second goroutine to send a value to the channel
	go func() {
		c <- 44
	}()
	val = <-c
	fmt.Println(val)
}
