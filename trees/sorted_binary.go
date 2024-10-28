package trees

func sortedInsert(node *Node, data string) *Node {
	if node == nil {
		return &Node{Data: data}
	}

	if data < node.Data {
		node.Left = sortedInsert(node.Left, data)
	} else {
		node.Right = sortedInsert(node.Right, data)
	}

	return node
}

func sortedFind(node *Node, target string) *Node {
	if node == nil {
		return nil
	}

	if target == node.Data {
		return node
	}

	if target < node.Data {
		return sortedFind(node.Left, target)
	}

	if target > node.Data {
		return sortedFind(node.Right, target)
	}

	return nil
}

/**************************************************
 * Sorted Binary Tree definition and operations
 **************************************************/
type SortedBinaryTree struct {
	Tree *BinaryTree
}

func NewSortedBinaryTree() *SortedBinaryTree {
	return &SortedBinaryTree{NewBinaryTree()}
}

func (t *SortedBinaryTree) Insert(data string) {
	sortedInsert(t.Tree.Root, data)
}

func (t *SortedBinaryTree) Find(target string) *Node {
	return sortedFind(t.Tree.Root, target)
}
