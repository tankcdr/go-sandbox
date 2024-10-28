package trees_test

import (
	"testing"

	"github.com/tankcdr/trees"
)

func TestSortedBinaryTree_Create(t *testing.T) {
	t.Parallel()

	trees.NewSortedBinaryTree()
}

func TestSortedBinaryTree_Insert(t *testing.T) {
	t.Parallel()

	tree := trees.NewSortedBinaryTree()

	tree.Insert("a")
	tree.Insert("b")
	tree.Insert("c")
	tree.Insert("d")
	tree.Insert("e")

	result := tree.Tree.InOrder()

	if result != "a b c d e " {
		t.Errorf("Expected 'a b c d e ', but got '%v'", result)
	}
}

func TestSortedBinaryTree_Find(t *testing.T) {
	t.Parallel()

	tree := trees.NewSortedBinaryTree()

	tree.Insert("a")
	tree.Insert("b")
	tree.Insert("c")
	tree.Insert("d")
	tree.Insert("e")

	result := tree.Find("c")

	if result == nil {
		t.Errorf("Expected 'found', but got '%v'", result)
	}

	if result.Data != "c" {
		t.Errorf("Expected 'c', but got '%v'", result.Data)
	}
}
