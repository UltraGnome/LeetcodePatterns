package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

var DefaultValue int = -1024

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func InsertNodeToTree(tree *TreeNode, node *TreeNode) {
	if tree == nil {
		return
	}
	if tree.Value == DefaultValue {
		tree.Value = node.Value
		return
	}
	if node.Value > tree.Value {
		if tree.Right == nil {
			tree.Right = &TreeNode{Value: DefaultValue}
		}
		InsertNodeToTree(tree.Right, node)
	}
	if node.Value < tree.Value {
		if tree.Left == nil {
			tree.Left = &TreeNode{Value: DefaultValue}
		}
		InsertNodeToTree(tree.Left, node)
	}
}

func InitTree(values ...int) (root *TreeNode) {
	rootNode := TreeNode{Value: DefaultValue, Right: nil, Left: nil}
	for _, value := range values {
		node := TreeNode{Value: value}
		InsertNodeToTree(&rootNode, &node)
	}
	return &rootNode
}

func main() {
	tree := InitTree(3, 9, 20, 15, 7)

	fmt.Println(maxDepth(tree))
	//nums := []int{5, 7, 7, 8, 8, 10}
	//fmt.Println(searchRange(nums, 3))

}
