package main

import "fmt"

func diaUm() {
	fmt.Println("Domingo")
}
func diaDois() {
	fmt.Println("Segunda")
}

func main() {

	defer diaDois()
	defer diaUm()
}
