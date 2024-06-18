package main

import "fmt"

func main() {
	evanVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evanVals {
		fmt.Println(i, v)
	}
}

// Output:
// 0 2
// 1 4
// 2 6
// 3 8
// 4 10
// 5 12
