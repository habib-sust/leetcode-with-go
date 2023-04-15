package main

// Problem Link: https://leetcode.com/problems/binary-tree-inorder-traversal/description/

func inorderTraversal(tree *TreeNode) []int {
	result := []int{}

	var inorder func(*TreeNode)

	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorder(root.Left)
		result = append(result, root.Val)
		inorder(root.Right)
	}

	inorder(tree)

	return result
}
