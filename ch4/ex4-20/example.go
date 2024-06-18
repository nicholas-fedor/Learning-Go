package main

import "fmt"

func main() {
	loopOne()
	fmt.Println()
	loopTwo()
}

// Loop missing "loop" label
func loopOne() {
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3, but not 2")
		case i%7 == 0:
			fmt.Println("exit the loop!")
			break
		default:
			fmt.Println(i, "is boring")
		}
	}
}

// Loop with "loop" label
func loopTwo() {
loop:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3, but not 2")
		case i%7 == 0:
			fmt.Println("exit the loop!")
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}
}

// Output:
// 0 is even
// 1 is boring
// 2 is even
// 3 is divisible by 3, but not 2
// 4 is even
// 5 is boring
// 6 is even
// exit the loop!
// 8 is even
// 9 is divisible by 3, but not 2

// 0 is even
// 1 is boring
// 2 is even
// 3 is divisible by 3, but not 2
// 4 is even
// 5 is boring
// 6 is even
// exit the loop!
