package main

import (
	"fmt"

	"github.com/tankcdr/trees"
)

func buildTree() *trees.BinaryTree {

	jNode := &trees.Node{"J", nil, nil}
	iNode := &trees.Node{"I", nil, nil}
	hNode := &trees.Node{"H", iNode, jNode}

	fNode := &trees.Node{"F", hNode, nil}
	cNode := &trees.Node{"C", nil, fNode}

	gNode := &trees.Node{"G", nil, nil}
	eNode := &trees.Node{"E", gNode, nil}
	dNode := &trees.Node{"D", nil, nil}
	bNode := &trees.Node{"B", dNode, eNode}

	aNode := &trees.Node{"A", bNode, cNode}

	return &trees.BinaryTree{Root: aNode}
}

func main() {
	tree := buildTree()
	fmt.Print(tree.DisplayIndented("  "))
	fmt.Println()

	fmt.Printf("PreOrder: %s\n", tree.PreOrder())
	fmt.Println()

	fmt.Printf("InOrder: %s\n", tree.InOrder())
	fmt.Println()

	fmt.Printf("PostOrder: %s\n", tree.PostOrder())
	fmt.Println()

	fmt.Printf("BreadthFirst: %s\n", tree.BreadthFirst())
	fmt.Println()
}
