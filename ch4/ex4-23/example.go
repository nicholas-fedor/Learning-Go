package main

import "fmt"

func main() {
	a := 10

	goto skip

	b := 20

skip:
	c := 30

	fmt.Println(a, b, c)

	if c > a {
		goto inner
	}

	if a < b {
	inner:
		fmt.Println("a is less than b")
	}
}

// Expected Errors:
// # command-line-arguments
// ./example.go:8:7: goto skip jumps over declaration of b at ./example.go:10:4
// ./example.go:18:8: goto inner jumps into block starting at ./example.go:21:11
