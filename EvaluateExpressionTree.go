package main

// import("fmt")

type BinaryTree struct{
	value int
	left *BinaryTree
	right *BinaryTree
}

type Operator int

const (
	Plus Operator = -1
	Minus  = -2
	Division = -3
	Multiplication = -4
)


func EvaluateExpressiontTree(tree *BinaryTree) int {
	if tree == nil { return 0 }

	if tree.value > 0 { return tree.value }

	left, right := EvaluateExpressiontTree(tree.left), EvaluateExpressiontTree(tree.right)

	operator := Operator(tree.value)

	switch operator {
	case Plus: return left+right
	case Minus: return left-right
	case Division: return left/right
	case Multiplication: return left*right
	default: return tree.value
	}
}

// func main() {
// 	tree := BinaryTree{value: 0}

// 	fmt.Println(EvaluateExpressiontTree(&tree))

// }