package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		// Prints 10
		fmt.Println(x)
		x := 5
		// Prints 5
		fmt.Println(x)
	}
	// Prints 10
	fmt.Println(x)
}
