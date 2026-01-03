package main

import "fmt"

func median(listOfGrades []float64) float64 {
	total := 0.0

	for _, valor := range listOfGrades {
		total += valor
	}
	return (total / float64(len(listOfGrades)))
}

func main() {
	listOfGrades := []float64{67, 98, 100, 32, 77}
	fmt.Println(median(listOfGrades))
}
