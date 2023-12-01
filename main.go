package imp

import "fmt"

func SayHello() {
	fmt.Println("Hello world!")
}

func SayGoodbye() {
	printGoodBye()
}

func printGoodBye() {
	fmt.Println("Goodbye")
}
