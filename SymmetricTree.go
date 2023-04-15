package main

// Problem: https://leetcode.com/problems/symmetric-tree/description/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var isSame func(*TreeNode, *TreeNode) bool

	isSame = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}

		if left == nil || right == nil {
			return false
		}

		if left.Val != right.Val {
			return false
		}

		return isSame(left.Left, right.Right) && isSame(left.Right, right.Left)
	}

	return isSame(root.Left, root.Right)

}
