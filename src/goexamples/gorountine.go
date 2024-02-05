package main

import (
	"fmt"
	"time"
)

func heavy() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("Hello!")
	}
}
func super_heavy() {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("'super heavy")
	}
}

func main() {
	go heavy()
	super_heavy()
	fmt.Println("Hello, World!")
	//select {}
}
