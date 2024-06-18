package main

import (
	"fmt"
	"math/rand"
)

func main() {
	switch n := rand.Intn(10); {
	case n == 0:
		fmt.Println("That's too low")
	case n > 5:
		fmt.Println("That's too big:", n)
	default:
		fmt.Println("That's a good number:", n)
	}
}

// Example Output:
// That's a good number: 2

// Varies as using pseudo-random generator.
