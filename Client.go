package main

import (
	"fmt"
)

func main() {
	root := TreeNode{Val: 0}

	// Evaluate Boolean Binary Tree
	fmt.Println(evaluateTree(&root))

	// Binary Tree Inorder Traversal
	fmt.Println(inorderTraversal(&root))

	// EvaluateExpressionTree
	tree := BinaryTree{value: 0}
	fmt.Println(EvaluateExpressiontTree(&tree))

	// Symmetric Tree
	fmt.Println(isSymmetric(&root))
}
