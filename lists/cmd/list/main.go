package main

import (
	"fmt"

	"github.com/tankcdr/lists"
)

func main() {
	aCell := lists.Cell{"Apple", nil}
	bCell := lists.Cell{"Banana", nil}
	cCell := lists.Cell{"Cherry", nil}

	aCell.AddAfter(&bCell)
	bCell.AddAfter(&cCell)
	top := &aCell

	for cell := top; cell != nil; cell = cell.Next {
		fmt.Printf("%s\n", cell.Data)
	}
	fmt.Println()

	cCell.DeleteAfter()

	for cell := top; cell != nil; cell = cell.Next {
		fmt.Printf("%s\n", cell.Data)
	}
	fmt.Println()
}
