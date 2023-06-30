package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{}
	root.Val = 5
	root.Left = &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	root.Right = &TreeNode{
		Left:  nil,
		Right: nil,
	}

	fmt.Println(hasPathSum(root, 9))
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {
			return true
		}

		return false
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
