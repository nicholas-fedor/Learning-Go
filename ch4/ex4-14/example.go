package main

import "fmt"

func main() {
	evanVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evanVals {
		fmt.Println(v)
	}
}

// Output:
// 2
// 4
// 6
// 8
// 10
// 12
