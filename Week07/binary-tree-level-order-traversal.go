package main

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	var helper func(root *TreeNode, level int)
	helper = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if len(res) == level {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)
		helper(root.Left, level+1)
		helper(root.Right, level+1)
	}
	helper(root, 0)
	return res
}
