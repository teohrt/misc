package main

import (
	"fmt"
)

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	t := triangle{
		height: 5,
		base:   10,
	}

	s := square{sideLength: 10}

	printArea(t)
	printArea(s)

}

func printArea(sh shape) {
	fmt.Println("Area of the shape is:", sh.getArea())
}

func (t triangle) getArea() float64 {
	return .5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
