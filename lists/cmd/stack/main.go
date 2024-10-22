package main

import (
	"fmt"

	"github.com/tankcdr/lists"
)

func main() {
	greekLetters := []string{
		"α", "β", "γ", "δ", "ε",
	}
	list := lists.NewLinkedList()
	list.AddRange(greekLetters)
	fmt.Println(list.ToString(" "))
	fmt.Println()

	// Demonstrate a stack.
	stack := lists.NewLinkedList()
	stack.Push("Apple")
	stack.Push("Banana")
	stack.Push("Coconut")
	stack.Push("Date")

	fmt.Printf("Stack %d: %s\n",
		stack.Length(),
		stack.ToString(" "))

	for !stack.IsEmpty() {
		fmt.Printf("Popped: %-7s   Remaining %d: %s\n",
			stack.Pop(),
			stack.Length(),
			stack.ToString(" "))
	}
}
