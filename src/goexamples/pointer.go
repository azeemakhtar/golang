package main

import "fmt"

func swap(a, b *int) {
	var temp int
	temp = *b
	*b = *a
	*a = temp
}

func main() {
	a, b := 2, 3
	//m1 := 2
	//ptr := &m1 // pointer is a variable that stores the address of another variable
	// & is reference operator that used to get the address of a variable
	//fmt.Println(ptr)  //this will print the address of m1 hexa decimal format
	//fmt.Println(*ptr) // * is dereference operator that used to get the value of a variable
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)
}
