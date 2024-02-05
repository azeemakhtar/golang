package main

import "fmt"

type Car interface {
	Drive()
	Stop()
}
type Lambo struct {
	LamboModel string
}

func NewModel(arg string) Car {
	return &Lambo{arg}
}

type Chevy struct {
	ChevyModel string
}

func (l *Lambo) Stop() {
	fmt.Println("Lambo stopped")
	//fmt.Println(l.LamboModel)
}
func (l *Lambo) Drive() {
	fmt.Println("Lambo is being driven")
	fmt.Println(l.LamboModel)
}
func (c *Chevy) Drive() {
	fmt.Println("Chevy is being driven")
	fmt.Println((c.ChevyModel))
}
func main() {
	l := Lambo{"Gallardo"}
	c := Chevy{"Cruze"}
	l.Drive()
	c.Drive()
	l.Stop()
}
