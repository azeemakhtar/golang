package main

import "fmt"

func todo() {
	arr := []int{1, 2, 3, 4, 5}
	arr2 := []string{"a", "b", "c", "d", "e"}
	fmt.Println("emp:", arr)

	arr2 = append(arr2, "f", "g")
	fmt.Println(arr, arr2)
}
func main() {
	todo()
}
