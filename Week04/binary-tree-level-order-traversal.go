package main

import "fmt"

//102. 二叉树的层序遍历
/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/

//type TreeNode struct {
//	Val int
//	Left *TreeNode
//	Right *TreeNode
//}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int,0)
	nodes := make([]*TreeNode,0)
	var helper  func(nodes []*TreeNode)
	helper = func(nodes []*TreeNode) {
		if len(nodes) == 0 {
			return
		}
		tmp := make([]int,0)
		newNodes := make([]*TreeNode,0)
		for _,node := range nodes {
			if node != nil {
				tmp = append(tmp,node.Val)
				if node.Left != nil {
					newNodes = append(newNodes,node.Left)
				}
				if node.Right != nil {
					newNodes = append(newNodes,node.Right)
				}
			}

		}
		if len(tmp)>0 {
			res = append(res,tmp)
		}
		helper(newNodes)
	}
	if root != nil {
		res = append(res,[]int{root.Val})
		nodes = append(nodes,root.Left)
		nodes = append(nodes,root.Right)
		helper(nodes)
	}
	return res
}



func main()  {
	//root := &TreeNode{
	//	Val:3,
	//	Left:&TreeNode{
	//		Val:9,
	//		Left:nil,
	//		Right:nil,
	//	},
	//	Right:&TreeNode{
	//		Val:20,
	//		Left:&TreeNode{
	//			Val:15,
	//			Left:nil,
	//			Right:nil,
	//		},
	//		Right:&TreeNode{
	//			Val:7,
	//			Left:nil,
	//			Right:nil,
	//		},
	//	},
	//}
	root := &TreeNode{
		Val:1,
		Left:nil,
		Right:nil,
	}
	fmt.Println(levelOrder(root))
}