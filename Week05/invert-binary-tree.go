package main

import "fmt"

//226. 翻转二叉树
//type TreeNode struct {
//	Val int
//	Left *TreeNode
//	Right *TreeNode
//}

func invertTree(root *TreeNode) *TreeNode {
	var helper func(root *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		node.Left,node.Right = node.Right,node.Left
		helper(node.Left)
		helper(node.Right)
	}
	helper(root)
	return root
}

func main()  {
	root:= &TreeNode{
		Val:4,
		Left:&TreeNode{
			Val:2,
			Left:&TreeNode{
				Val:1,
				Left:nil,
				Right:nil,
			},
			Right:&TreeNode{
				Val:3,
				Left:nil,
				Right:nil,
			},
		},
		Right:&TreeNode{
			Val:2,
			Left:&TreeNode{
				Val:3,
				Left:nil,
				Right:nil,
			},
			Right:&TreeNode{
				Val:1,
				Left:nil,
				Right:nil,
			},
		},
	}

	fmt.Println(invertTree(root))

}

