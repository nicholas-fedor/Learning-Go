package main

import "fmt"

func main() {
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}
}

// Expected Output:
// 1
// 2
// 4
// 8
// 16
// 32
// 64