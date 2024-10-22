package main

import (
	"fmt"

	"github.com/tankcdr/lists"
)

func main() {
	// Make a list from an array of values.
	values := []string{
		"0", "1", "2", "3", "4", "5",
	}
	list := lists.NewLinkedList()
	list.AddRange(values)

	fmt.Println(list.ToString(" "))
	if list.HasLoop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
	fmt.Println()

	// Make cell 5 point to cell 2.
	five := list.Find("4")
	two := list.Find("1")
	five.Next = two

	fmt.Println(list.ToStringMax(" ", 10))
	if list.HasLoop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
	fmt.Println()

	// Make cell 4 point to cell 2.
	four := list.Find("3")
	four.Next = two

	fmt.Println(list.ToStringMax(" ", 10))
	if list.HasLoop() {
		fmt.Println("Has loop")
	} else {
		fmt.Println("No loop")
	}
}
