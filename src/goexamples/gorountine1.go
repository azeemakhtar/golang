package main

import (
	"fmt"
	"sync"
)

func main() {
	// wait group gives us the ability to wait for a group of go routines to finish
	wg := &sync.WaitGroup{}
	//wiat goups give add, done and wait methods
	// add 2 go routines
	wg.Add(2)
	// run 2 go routines
	go func() {
		fmt.Println("Hello, World!")
		wg.Done()
	}()
	go func() {
		fmt.Println("World!")
		wg.Done()
	}()
	fmt.Println("Waiting...")
	wg.Wait()
	fmt.Println("Fin")
}
