package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 10
	}()
	fmt.Println(<-c)
}
