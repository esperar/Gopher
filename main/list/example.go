package main

import (
	"container/list"
	"fmt"
)

func main() {

	ll := list.New()

	ll.PushBack("A")
	ll.PushBack(100)
	ll.PushBack(true)
	ll.PushFront("B")
	ll.PushFront(200)

	for e := ll.Front(); e != nil; e = e.Next() {
		fmt.Printf("[%T] %v\n", e.Value, e.Value)
	}

	fmt.Println("--------------")

	for e := ll.Back(); e != nil; e = e.Prev() {
		fmt.Printf("[%T] %v\n", e.Value, e.Value)
	}
}
