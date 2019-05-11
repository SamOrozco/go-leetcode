package main

func main() {

}

type TreeNode struct {
        Val   int
        Left  *TreeNode
        Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	return sumNode(root, L, R)
}

func sumNode(node *TreeNode, l int, r int) int {
	if node == nil {
		return 0
	}

	sum := 0
	nodeVal := node.Val
	if inRange(nodeVal, l, r) {
		sum += nodeVal
	}
	if node := node.Right; node != nil {
		sum += sumNode(node, l, r)
	}
	if node := node.Left; node != nil {
		sum += sumNode(node, l, r)
	}

	return sum
}

func inRange(val, l, r int) bool {
	return l <= val && r >= val
}