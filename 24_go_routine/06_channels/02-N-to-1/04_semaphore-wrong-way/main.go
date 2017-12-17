package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	// blocks until done is sent a message
	<-done
	<-done
	close(c)

	// blocked untill done is sent a message (causing dead lock)
	for n := range c {
		fmt.Println(n)
	}
}
