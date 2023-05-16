package rbtree

import "testing"

func TestRBTree(t *testing.T) {
	tree := NewRedBlackTree()
	tree.Put(1, 10)
	tree.Put(2, 11)
	tree.Put(3, 12)

	val := tree.Get(2)
	if val != 11 {
		t.Fatal("Unexpected value for Get")
	}

	tree.Delete(2)
	val = tree.Get(2)
	if val != nil {
		t.Fatal("Unexpected value for Delete")
	}
	afterDeleteVal := tree.Get(1)
	if afterDeleteVal != 10 {
		t.Fatal("Unexpected value for Get")
	}

}
