package binarytree

import "fmt"

// Tree дерево они же узлы
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Traverse Выводит все элементы двоичного дерева
func Traverse(t *Tree) {
	if t == nil {
		return
	}

	Traverse(t.Left)
	fmt.Print(t.Value, " ")
	Traverse(t.Right)
}

// Insert  добавить элементв дерево
func Insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v == t.Value {
		return t
	}
	if v < t.Value {
		t.Left = Insert(t.Left, v)
		return t
	}
	t.Right = Insert(t.Right, v)
	return t
}
