package main

import "math"
import "fmt"

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	side float64
}

func (s *square) area() float64 {
	return s.side * s.side
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	c := circle{23.3}
	s := square{12.3}

	info(c)
	info(&c) // This works since c has a value reciever which accetps values & pointer
	// info(s) this will not work since c has a pointer reciever which only accepts an address
	info(&s)
}
