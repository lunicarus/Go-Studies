package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan bool)
	c2 := make(chan bool)

	go func() {
		for {
			<-c1
			fmt.Println("Ping")
			time.Sleep(time.Second)
			c2 <- true
		}
	}()

	go func() {
		for {
			<-c2
			fmt.Println("Pong")
			time.Sleep(time.Second)
			c1 <- true
		}
	}()
	c1 <- true

	select {}
}
