package main

import (
	"fmt"
)

func printN(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(n, ":", i)
	}

}
func main() {
	go printN(10)
	var written string
	fmt.Scanln(&written)
}
