package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)
	fmt := "ooops"
	// Expected error: fmt.Println undefined (type string has no field or method Println)
	fmt.Println(fmt)
}
