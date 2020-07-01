package main

import "fmt"

//104. 二叉树的最大深度
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}else{
		lHeight := maxDepth(root.Left)
		RHeight := maxDepth(root.Right)
		if lHeight > RHeight {
			return lHeight
		}else{
			return RHeight
		}
	}
}

func main()  {
	root := &TreeNode{
		Val:3,
		Left:&TreeNode{
			Val:9,Left:nil,Right:nil,
		},
		Right:&TreeNode{Val:20,Left:&TreeNode{Val:15,Left:nil,Right:nil,},Right:&TreeNode{Val:7,Left:nil,Right:nil}},
	}

	fmt.Println(maxDepth(root))
}