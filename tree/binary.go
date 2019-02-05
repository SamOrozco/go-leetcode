package main

import "math/rand"

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return &Tree{}
}

func (tree *Tree) printInOrder() {
	if tree.root == nil {
		println("Empty Tree")
	}
	tree.root.printInOrder()
}

func (tree *Tree) printPreOrder() {
	if tree.root == nil {
		println("Empty Tree")
	}
	tree.root.printPreOrder()
}

func (tree *Tree) printPostOrder() {
	if tree.root == nil {
		println("Empty Tree")
	}
	tree.root.printPostOrder()
}

func (tree *Tree) contains(val int) bool {
	if tree.root == nil {
		return false
	}
	return tree.root.contains(val)
}

func (tree *Tree) add(val int) {
	if tree.root == nil { // val is the root now
		tree.root = NewNode(val)
	} else {
		tree.root = insert(tree.root, val)
	}
}

func insert(node *Node, val int) *Node {
	if node == nil {
		return NewNode(val)
	}
	if val > node.val {
		node.right = insert(node.right, val)
	}
	if val < node.val {
		node.left = insert(node.left, val)
	}
	return node
}

type Node struct {
	left  *Node
	right *Node
	val   int
}

func (node *Node) contains(val int) bool {
	if node.val == val {
		return true
	}
	if val > node.val {
		if node.right != nil {
			return node.right.contains(val)
		}
		return false
	}
	if node.left != nil {
		return node.left.contains(val)
	}
	return false
}

func (node *Node) printInOrder() {
	if node.left != nil {
		node.left.printInOrder()
	}
	println(node.val)
	if node.right != nil {
		node.right.printInOrder()
	}
}

func (node *Node) printPreOrder() {
	println(node.val)
	if node.left != nil {
		node.left.printInOrder()
	}
	if node.right != nil {
		node.right.printInOrder()
	}
}

func (node *Node) printPostOrder() {
	if node.left != nil {
		node.left.printInOrder()
	}
	if node.right != nil {
		node.right.printInOrder()
	}
	println(node.val)
}

func NewNode(val int) *Node {
	return &Node{
		val: val,
	}
}

func main() {
	tree := NewTree()
	for i := 0; i < 10; i ++ {
		tree.add(rand.Int() % 1024)
	}
	tree.printInOrder()
	println(tree.contains(3))
}
