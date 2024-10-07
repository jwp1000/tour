package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func WalkThenClose(treep *tree.Tree, ch chan int) {
	Walk(treep, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go WalkThenClose(t1, ch1)
	go WalkThenClose(t2, ch2)

	for v1 := range ch1 {
		v2 := <-ch2
		fmt.Println("v1:", v1, ", v2:", v2)
		if v1 != v2 {
			return false
		}
	}

	return true
}

func main() {
	t1, t2 := tree.New(1), tree.New(1)
	fmt.Println("t1:", t1)
	fmt.Println("t2:", t2)
	test := Same(t1, t2)
	fmt.Println("Same? ", test)
	t1, t2 = tree.New(1), tree.New(2)
	fmt.Println("t1:", t1)
	fmt.Println("t2:", t2)
	test = Same(t1, t2)
	fmt.Println("Same? ", test)

}
