package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (mylist *List[T]) append(val T) {
	if mylist.next == nil {
		mylist.next = &List[T]{val: val}
	} else {
		mylist.next.append(val)
	}
}

func (mylist List[T]) String() string {
	if mylist.next == nil {
		return fmt.Sprintf("{ val: %v, next: <nil> }", mylist.val)
	}
	return fmt.Sprintf("{ val: %v, next: %v }", mylist.val, mylist.next)
}

func main() {
	intList := new(List[int])
	intList.val = 5
	fmt.Println(intList)
	intList.append(6)
	fmt.Println(intList)
	intList.append(7)
	fmt.Println(intList)
}
