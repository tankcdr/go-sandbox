package trees

import "strings"

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

	if n.Data != "" {
		result += n.Data + " "
	}

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
