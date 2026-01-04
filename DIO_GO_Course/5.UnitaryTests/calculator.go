package main

import "fmt"

func main() {

}

func multiply(i ...int) int {
	total := i[0]
	for x := 1; x < len(i); x++ {
		total *= i[x]
	}

	return total
}
func divide(i ...float32) (float32, error) {

	//dead man tell no tales
	if len(i) == 0 {
		return 0, fmt.Errorf("no values provided")
	}
	if i[0] == 0 {
		return 0, fmt.Errorf("impossible to divide by zero")
	}
	total := i[0]
	for x := 1; x < len(i); x++ {
		total /= i[x]
	}

	return total, nil
}
func add(i ...int) int {
	total := i[0]
	for x := 1; x < len(i); x++ {
		total += i[x]
	}

	return total
}
func sub(i ...int) int {
	total := i[0]
	for x := 1; x < len(i); x++ {
		total -= i[x]
	}

	return total
}
