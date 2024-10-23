package trees

import (
	"strings"

	"github.com/tankcdr/lists"
)

/**************************************************
 * Tree Node definition and operations
 **************************************************/
type Node struct {
	Data  string
	Left  *Node
	Right *Node
}

func (n *Node) displayIndented(seperator string, depth int) string {
	result := strings.Repeat(seperator, depth) + n.Data + "\n"
	if n.Left != nil {
		result += n.Left.displayIndented(seperator, depth+1)
	}

	if n.Right != nil {
		result += n.Right.displayIndented(seperator, depth+1)
	}

	return result
}

func (n *Node) preOrder() string {
	result := n.Data + " "
	if n.Left != nil {
		result += n.Left.preOrder()
	}

	if n.Right != nil {
		result += n.Right.preOrder()
	}

	return result
}

func (n *Node) inOrder() string {
	result := ""
	if n.Left != nil {
		result += n.Left.inOrder()
	}

	result += n.Data + " "

	if n.Right != nil {
		result += n.Right.inOrder()
	}

	return result
}

func (n *Node) postOrder() string {
	result := ""
	if n.Left != nil {
		result += n.Left.postOrder()
	}

	if n.Right != nil {
		result += n.Right.postOrder()
	}

	result += n.Data + " "

	return result
}

/**************************************************
 * Binary Tree definition and operations
 **************************************************/
type BinaryTree struct {
	Root *Node
}

func (t *BinaryTree) DisplayIndented(seperator string) string {
	results := t.Root.displayIndented(seperator, 0)
	return results
}

func (t *BinaryTree) PreOrder() string {
	results := t.Root.preOrder()
	return results
}

func (t *BinaryTree) InOrder() string {
	results := t.Root.inOrder()
	return results
}

func (t *BinaryTree) PostOrder() string {
	results := t.Root.postOrder()
	return results
}

func (t *BinaryTree) BreadthFirst() string {
	queue := lists.NewDoublyLinkedList[*Node]()
	queue.Enqueue(t.Root)
	results := ""

	for queue.Length() > 0 {
		results += " "

		node := queue.Dequeue()
		results += node.Data

		if node.Left != nil {
			queue.Enqueue(node.Left)
		}

		if node.Right != nil {
			queue.Enqueue(node.Right)
		}
	}

	return results
}
