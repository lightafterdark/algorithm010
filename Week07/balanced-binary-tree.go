package main

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

//func isBalanced(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	if math.Abs(float64(height(root.Left)-height(root.Right))) > 1 {
//		return false
//	} else {
//		return isBalanced(root.Left) && isBalanced(root.Right)
//	}
//}
//
//func height(root *TreeNode) int {
//	if root == nil {
//		return -1
//	}
//	if height(root.Left) > height(root.Right) {
//		return height(root.Left) + 1
//	} else {
//		return height(root.Right) + 1
//	}
//}

func isBalanced(root *TreeNode) bool {
	return helper(root) != -1
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left)
	if left == -1 {
		return -1
	}
	right := helper(root.Right)
	if right == -1 {
		return -1
	}
	if math.Abs(float64(left-right)) < 2 {
		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	} else {
		return -1
	}
}

func main() {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	fmt.Println(isBalanced(root))
}
