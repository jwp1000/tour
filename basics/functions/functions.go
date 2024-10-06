package main

import "fmt"

func add(x, y, z int) string {
	return string(x + y + z)
}

func main() {
	fmt.Println(add(42, 13, 11))
}
