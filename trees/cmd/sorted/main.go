package main

import (
	"fmt"

	"github.com/tankcdr/trees"
)

func main() {
	// Make a root node to act as sentinel.
	sorted := trees.NewSortedBinaryTree()

	// Add some values.
	sorted.Insert("I")
	sorted.Insert("G")
	sorted.Insert("C")
	sorted.Insert("E")
	sorted.Insert("B")
	sorted.Insert("K")
	sorted.Insert("S")
	sorted.Insert("Q")
	sorted.Insert("M")

	// Add F.
	sorted.Insert("F")

	// Display the values in sorted order.
	fmt.Printf("Sorted values: %s\n", sorted.Tree.InOrder())

	// Let the user search for values.
	for {
		// Get the target value.
		target := ""
		fmt.Printf("String: ")
		fmt.Scanln(&target)
		if len(target) == 0 {
			break
		}

		// Find the value's node.
		node := sorted.Find(target)
		if node == nil {
			fmt.Printf("%s not found\n", target)
		} else {
			fmt.Printf("Found value %s\n", target)
		}
	}
}
