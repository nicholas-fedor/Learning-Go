package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big", n)
	} else {
		fmt.Println("That's a good number", n)
	}
	// Expected error: variable n is undefined as prior variable n is only
	// within scope of the if-statement
	fmt.Println(n)
}
