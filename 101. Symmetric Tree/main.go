package main

import "fmt"

func main() {
	p := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}

	q := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}

	fmt.Println(isSameTree(p, q))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	result := false
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val == q.Val {
			result = isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
		}
	}

	return result
}
