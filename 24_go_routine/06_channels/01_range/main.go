package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	// no longer need time.Sleep(), the program blocks
	for n := range c {
		fmt.Println(n)
	}
}
