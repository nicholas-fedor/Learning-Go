package main

import "fmt"

func main() {
	totalWins := map[string]int{}

	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2

	// Prints 1
	fmt.Println(totalWins["Orcas"])

	// Prints 0
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++
	// Prints 1
	fmt.Println(totalWins["Kittens"])

	totalWins["Lions"] = 3
	// Prints 3
	fmt.Println(totalWins["Lions"])
}
