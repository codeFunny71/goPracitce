package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	foo()
	fmt.Println("something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func bar() {
	fmt.Println("and then we exited")
}

func foo() {
	fmt.Println("I'm in foo")
}

// control flow:
// (1) sequence
// (2) loop; iterative
// (1) conditional
