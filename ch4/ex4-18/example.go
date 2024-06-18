package main

import "fmt"

func main() {
	samples := []string{"hello", "apple_Π"}
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}
}

// Output:
// 0 104 h
// 1 101 e
// 2 108 l
// 0 97 a
// 1 112 p
// 2 112 p
// 3 108 l
