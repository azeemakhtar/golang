package main

import (
	"fmt"
	"strings"
)

func main() {
	m1 := "Hello"       // string
	m2 := "Hello World" // string

	fmt.Println(strings.Split(m2, " "), m1+m2)
}
