package main

import (
	"fmt"

	"github.com/tankcdr/lists"
)

func main() {
	// Test queue functions.
	fmt.Printf("*** Queue Functions ***\n")
	queue := lists.NewDoublyLinkedList[string]()
	queue.Enqueue("Agate")
	queue.Enqueue("Beryl")
	fmt.Printf("%s ", queue.Dequeue())
	queue.Enqueue("Citrine")
	fmt.Printf("%s ", queue.Dequeue())
	fmt.Printf("%s ", queue.Dequeue())
	queue.Enqueue("Diamond")
	queue.Enqueue("Emerald")
	for !queue.IsEmpty() {
		fmt.Printf("%s ", queue.Dequeue())
	}
	fmt.Printf("\n\n")

	// Test deque functions. Names starting
	// with F have a fast pass.
	fmt.Printf("*** Deque Functions ***\n")
	deque := lists.NewDoublyLinkedList[string]()
	deque.PushTop("Ann")
	deque.PushTop("Ben")
	fmt.Printf("%s ", deque.PopBottom())
	deque.PushBottom("F-Cat")
	fmt.Printf("%s ", deque.PopBottom())
	fmt.Printf("%s ", deque.PopBottom())
	deque.PushBottom("F-Dan")
	deque.PushTop("Eva")
	for !deque.IsEmpty() {
		fmt.Printf("%s ", deque.PopBottom())
	}
	fmt.Printf("\n")
}
