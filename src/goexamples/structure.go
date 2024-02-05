package main

import "fmt"

// Structure is a user defined data type which holds its own data members and member functions

type Car struct {
	Name  string
	Age   int
	Model int
}

func (c Car) Print() {
	fmt.Println(c)
}
func (c Car) Drive() {
	fmt.Println("Driving...")
}
func (c Car) Stop() {
	fmt.Println(("Stopping..."))
}
func (c Car) GetName() string {
	return c.Name
}
func (c Car) GetModel() int {
	return c.Model
}
func main() {
	c := Car{
		Name:  "BMW",
		Age:   2,
		Model: 2022,
	}
	c.Print()
	c.Drive()
	c.Stop()
	fmt.Println(c.GetName())
	fmt.Println(c.GetModel())
}
