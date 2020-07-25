package main

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

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return getAns(1, n)
}

func getAns(start int, end int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if start > end {
		res = append(res, nil)
		return res
	}
	for i := start; i <= end; i++ {
		leftTrees := getAns(start, i-1)
		rightTrees := getAns(i+1, end)
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				res = append(res, &TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				})
			}
		}
	}
	return res
}
