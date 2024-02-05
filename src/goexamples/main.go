// File: main.go
package main

import (
	"fmt"
)

func main1() {
	var a uint8

	//a = 12
	fmt.Printf("Hello %d\n", a)
	var name string
	var age int

	fmt.Println("What is your name?")
	fmt.Scanln(&name)

	fmt.Println("What is your age?")
	_, ee := fmt.Scanln(&age)
	if ee != nil {
		fmt.Println("Error: ")
	}
	fmt.Printf("Hello %s you are %d years old", name, age)
}
