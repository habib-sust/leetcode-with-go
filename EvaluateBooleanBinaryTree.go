package main

/* problem Link: https://leetcode.com/problems/evaluate-boolean-binary-tree*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func evaluateTree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	if root.Val == 0 || root.Val == 1 {
		return evaluate(root.Val)
	}

	left, right := evaluateTree(root.Left), evaluateTree(root.Right)

	if root.Val == 2 {
		return left || right
	} else {
		return left && right
	}
}

func evaluate(value int) bool {
	if value == 1 {
		return true
	} else {
		return false
	}
}
