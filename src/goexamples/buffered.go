package main

import "fmt"

//initialize the buffered channel

func main() {
	buf := make(chan int, 3)

	//make 3 go routines to send values to the channel

	go func() {
		buf <- 1
		buf <- 2
		buf <- 3
		close(buf)
	}()
	//make 3 go routines to receive values from the channel
	for i := range buf {
		fmt.Println(i)
	}
	//close(buf)
}
