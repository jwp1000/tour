package main

import "fmt"

func main() {
	v := func() string { return "j" } // change me!
	fmt.Printf("v is of type %T\n", v)
}
