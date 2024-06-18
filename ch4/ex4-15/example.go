package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
}

// Example Output:
// Loop 0
// a 1
// c 3
// b 2
// Loop 1
// a 1
// c 3
// b 2
// Loop 2
// a 1
// c 3
// b 2

// Note that the output can vary for each run due to how Go's map hashing
// algorithm works.
// fmt.Println always outputs maps with their keys in ascending sorted order.
