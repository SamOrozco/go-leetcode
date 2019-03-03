package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 */
type BSTIterator struct {
	root *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{root: root}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.root != nil

}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
