package main

import "fmt"

func Anything(anything interface{}) {
	fmt.Println(anything)
}
func main() {
	fmt.Println("Hello, World!")
	fmt.Println(2.434)
	Anything(struct{}{})

	mymap := make(map[string]interface{})

	mymap["name"] = "Azeem"
	mymap["age"] = 42
	fmt.Println(mymap)

	f := true
	flag := &f

	if flag == nil {
		fmt.Println("Flag is nil")
	} else if *flag {
		fmt.Println("Flag is true")
	} else {
		fmt.Println("Flag is false")
	}

	mymap1 := make(map[string]interface{})
	mymap1["name"] = "Azeem"
	mymap1["age"] = 42

	for k, v := range mymap1 {
		fmt.Printf("Key is %s, and Value is %v", k, v)
	}
}
