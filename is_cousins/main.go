package main

/**
 * Definition for a binary tree node.
 */
func isCousins(root *TreeNode, x int, y int) bool {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDepthAndParent(node *TreeNode, val int) (int, int) {
	if node == nil ||
		!node.hasChildren() {
		return -1, -1
	}

	if node.Left != nil {
		return findDepthAndParent();
	};
}

func (nd *TreeNode) hasChildren() bool {
	return nd.Left != nil && nd.Right != nil
}
