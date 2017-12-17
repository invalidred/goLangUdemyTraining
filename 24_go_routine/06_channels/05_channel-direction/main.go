package main

import "fmt"

func main() {
	c := incrementor()
	for n := range puller(c) {
		fmt.Println(n)
	}
}

func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		sum := 0
		for i := range c {
			sum += i
		}
		out <- sum
		close(out)
	}()
	return out
}
