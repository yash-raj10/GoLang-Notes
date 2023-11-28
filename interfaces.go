package main

import "fmt"

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

type triangle struct {
	height float64
	base   float64
}

func (t triangle) getArea() float64 {
	return t.height * t.base
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Print(s.getArea())
}

func main() {
	t := triangle{base: 5, height: 10}
	s := square{sideLength: 10}

	printArea(t)
	printArea(s)
}
