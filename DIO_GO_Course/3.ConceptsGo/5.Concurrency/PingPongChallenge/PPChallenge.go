package main

import (
	"fmt"
	"time"
)

func main() {
	ping := make(chan bool)
	pong := make(chan bool)

	go func() {
		for {
			<-ping
			fmt.Println("Ping")
			time.Sleep(time.Second)
			pong <- true
		}
	}()

	go func() {
		for {
			<-pong
			fmt.Println("Pong")
			time.Sleep(time.Second)
			ping <- true
		}
	}()
	ping <- true

	select {}
}
