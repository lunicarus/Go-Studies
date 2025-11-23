package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	side float64
}
type circle struct {
	rad float64
}
type geometry interface {
	area() float64
}

func (r rectangle) area() float64 {
	return r.side * r.side
}

func (c circle) area() float64 {
	return c.rad * math.Pi
}

func evalutateSize(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
}
func main() {
	rectangle1 := rectangle{side: 4}
	//fmt.Println(rectangle1.area())
	evalutateSize(rectangle1)
}
