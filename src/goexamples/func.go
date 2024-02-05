package main

import "fmt"

func myfunc(a int, b int) {
	total := 0
	total = a + b
	return total
	//fmt.Println("Total:", total)
}
func main() {
	sum := add(10, 20)
	fmt.Println("Sum:", sum)
}
