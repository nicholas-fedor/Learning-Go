package main

import "fmt"

func main() {
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 2, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	// Prints 11 7
	fmt.Println(len(vals), len(intSet))
	// Prints true
	fmt.Println(intSet[5])
	// Prints false
	fmt.Println(intSet[500])
	// Prints 100 is not in the set
	if intSet[100] {
		fmt.Println("100 is in the set")
	} else {
		fmt.Println("100 is not in the set")
	}
}
