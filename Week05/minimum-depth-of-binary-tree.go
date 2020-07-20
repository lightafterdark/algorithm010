package main

import "fmt"

//111. 二叉树的最小深度
//type TreeNode struct {
//	Val int
//	Left *TreeNode
//	Right *TreeNode
//}
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lHeight := minDepth(root.Left)
	rRight := minDepth(root.Right)
	if lHeight==0 || rRight == 0 {
		return lHeight+rRight + 1
	}else{
		if lHeight < rRight {
			return lHeight+1
		}else{
			return rRight+1
		}
	}
}

func main()  {
	root:= &TreeNode{
		Val:3,
		Left:&TreeNode{
			Val:9,
			Left:nil,
			Right:nil,
		},
		Right:&TreeNode{
			Val:20,
			Left:&TreeNode{
				Val:15,
				Left:nil,
				Right:nil,
			},
			Right:&TreeNode{
				Val:7,
				Left:nil,
				Right:nil,
			},
		},
	}

	fmt.Println(minDepth(root))

}