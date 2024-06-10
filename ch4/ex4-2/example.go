package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		x, y := 5, 20
		// Prints 5 20
		fmt.Println(x, y)
	}
	// Prints 10
	fmt.Println(x)
}
