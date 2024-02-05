package main

import "fmt"

func main() {
	fmt.Println("Waht is the day today?")
	var day string
	fmt.Scanln(&day)

	if day == "monday" {
		fmt.Println("Shahzain will go to school")
	} else if day == "tuesday" {
		fmt.Println("Shahzain will not go to school")
	}
}
