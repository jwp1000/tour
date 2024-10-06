package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	previous := 0
	next := 1
	return func() int {
		old := previous
		sum := previous + next
		previous = next
		next = sum
		return old
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
