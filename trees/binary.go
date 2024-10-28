package trees

import (
	"github.com/tankcdr/lists"
)

/**************************************************
 * Binary Tree definition and operations
 **************************************************/
type BinaryTree struct {
	Root *Node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{&Node{}}
}

func (t *BinaryTree) DisplayIndented(seperator string) string {
	return t.Root.displayIndented(seperator, 0)
}

func (t *BinaryTree) PreOrder() string {
	return t.Root.preOrder()
}

func (t *BinaryTree) InOrder() string {
	return t.Root.inOrder()
}

func (t *BinaryTree) PostOrder() string {
	return t.Root.postOrder()
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
