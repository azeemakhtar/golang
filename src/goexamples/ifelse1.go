package main

import "fmt"

func main() {
	// Here’s a basic if & elsse example.
	if 7%2 == 0 {
		fmt.Println("7 ia even")
	} else {
		fmt.Println("7 is odd")
	}
	// if statement can used without an else.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	//if statement with comparison operator.
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 9 are even")
	}
	// if else if else statement.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
