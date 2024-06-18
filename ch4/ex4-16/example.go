package main

import "fmt"

func main() {
	samples := []string{"hello", "apple_Π!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
}

// Output:
// 0 104 h
// 1 101 e
// 2 108 l
// 3 108 l
// 4 111 o

// 0 97 a
// 1 112 p
// 2 112 p
// 3 108 l
// 4 101 e
// 5 95 _
// 6 110 Π
// 7 33 !
