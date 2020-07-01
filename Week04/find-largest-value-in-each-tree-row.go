package main

import "fmt"

//515. 在每个树行中找最大值

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func largestValues(root *TreeNode) []int {
	res := make([]int,0)
	var helper func(nodes []*TreeNode)
	helper = func(nodes []*TreeNode) {
		if len(nodes)  == 0 {
			return
		}
		max := 0
		newNodes := make([]*TreeNode,0)
		for index,node := range nodes {
			if index == 0 {
				max = node.Val
			}
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				newNodes = append(newNodes,node.Left)
			}

			if node.Right != nil {
				newNodes = append(newNodes,node.Right)
			}
		}

		res = append(res,max)
		if len(newNodes) > 0 {
			helper(newNodes)
		}
	}

	nodes := make([]*TreeNode,0)
	if root != nil {
		res = append(res,root.Val)
		if root.Left != nil {
			nodes = append(nodes,root.Left)
		}
		if root.Right != nil {
			nodes = append(nodes,root.Right)
		}
		if len(nodes) > 0 {
			helper(nodes)
		}
	}

	return res
}

func main()  {
	root := &TreeNode{
		Val:0,
		Left:&TreeNode{
			Val:-1,
			Left:nil,
			Right:nil,
		},
		Right:nil,
	}
	fmt.Println(largestValues(root))
}
