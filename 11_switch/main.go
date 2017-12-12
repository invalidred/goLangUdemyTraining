package main

import "fmt"

// Contact type
type Contact struct {
	greeting string
	name     string
}

func switchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	switch "Sara Abdo" {
	case "Sara Abdo", "Sara", "sara":
		fmt.Println("Hi Sara Abdo")
		fallthrough
	case "Deen Khan":
		fmt.Println("Hi Deen Khan")
	case "Zayn Khan":
		fmt.Println("Hi Zayn Khan")
	default:
		fmt.Println("Dont you have any family Abdul Khan?")
	}

	// no-expression switch
	myFriendName := "Mar"

	switch {
	case len(myFriendName) == 2:
		fmt.Println("Wassup my friend with name of length 2")
	case myFriendName == "Tim":
		fmt.Println("Wassup Tim")
	case myFriendName == "Julian":
		fmt.Println("Wassup Sushant")
	default:
		fmt.Println("Nothing matched")
	}

	// switch on type
	switchOnType(7)
	switchOnType("McLeod")
	var t = Contact{"Good to see you", "Tim"}
	switchOnType(t)
	switchOnType(t.greeting)
	switchOnType(t.name)
}
