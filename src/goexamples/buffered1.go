package main

import "fmt"

type Car struct {
	Model string
}

func main() {
	//initiate a buffered channel
	c := make(chan *Car, 3)

	//send 3 values to the channel with a go routine
	go func() {
		c <- &Car{"Saab"} // Fix: Change the value 1 to a string value inside the struct literal
		c <- &Car{"Volvo"}
		c <- &Car{"Mercedes"}
		c <- &Car{"Toyota"}
		close(c)
	}()
	for i := range c {
		fmt.Println(i.Model)
	}
}
