package main

import (
	"fmt"
	"os"
	"time"
)

func Select(c chan int, quits chan struct{}) {
	//switch case
	//Select for Channel asyn
	for {
		time.Sleep(time.Second)
		select {
		case <-c:
			fmt.Println("Received")
		case <-quits:
			fmt.Println("Quits")
			os.Exit(0)
		}
	}

}
func main() {
	c := make(chan int, 2)
	quits := make(chan struct{})

	go func() {
		c <- 1
		quits <- struct{}{}
	}()
	go Select(c, quits)
	select {}
}
