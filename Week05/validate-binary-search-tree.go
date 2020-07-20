package main

import (
	"fmt"
	"math"
)

//98. 验证二叉搜索树

//type TreeNode struct {
//	Val int
//	Left *TreeNode
//	Right *TreeNode
//}
func isValidBST(root *TreeNode) bool {
	return helper(root,math.MinInt64,math.MaxInt64)
}

func helper(root *TreeNode,lower,upper int)  bool{
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left,lower,root.Val) && helper(root.Right,root.Val,upper)
}

func main()  {
	root := &TreeNode{
		Val:10,
		Left:&TreeNode{
			Val:5,
			Left:nil,
			Right:nil,
		},
		Right:&TreeNode{
			Val:4,
			Left:&TreeNode{
				Val:3,
				Left:nil,
				Right:nil,
			},
			Right:&TreeNode{
				Val:6,
				Left:nil,
				Right:nil,
			},
		},
	}
	fmt.Println(isValidBST(root))
}