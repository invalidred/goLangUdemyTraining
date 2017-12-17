package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

// special func call before main is executed
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // No longer required from Go 1.5
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo: ", i)
		time.Sleep(time.Duration(100 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 100; i++ {
		fmt.Println("Bar: ", i)
		time.Sleep(time.Duration(100 * time.Millisecond))
	}
	wg.Done()
}
