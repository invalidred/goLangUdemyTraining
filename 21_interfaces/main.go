package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
}

func main() {
	s := square{4.5}
	c := circle{4.2}
	info(s)
	info(c)
}
