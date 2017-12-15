package main

import (
	"fmt"
)

// Person struct
type Person struct {
	Name string
	Age  int
}

// DoubleZero struct
type DoubleZero struct {
	Person
	LicenseToKill bool
}

// Greeting method
func (p Person) Greeting() {
	fmt.Println("I am just a regular person")
}

// Greeting method
func (dz DoubleZero) Greeting() {
	fmt.Println("I am a Double Zero")
}

func main() {
	p1 := Person{"Abdul", 12}
	p2 := DoubleZero{
		Person: Person{
			Name: "James Bond",
			Age:  23,
		},
		LicenseToKill: true,
	}
	p1.Greeting()        // I am just a regular person
	p2.Greeting()        // I am a Double Zero
	p2.Person.Greeting() // I am just a regular person
}
