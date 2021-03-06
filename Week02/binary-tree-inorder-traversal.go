package main

import "fmt"

//94. 二叉树的中序遍历
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
//func inorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return []int{}
//	}
//	rest := append(inorderTraversal(root.Left),root.Val)
//	rest = append(rest,inorderTraversal(root.Right)...)
//	return rest
//}

type Stack []*TreeNode

func (s *Stack) Push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *Stack) Pop() *TreeNode {
	n := []*TreeNode(*s)[len(*s)-1]
	*s = []*TreeNode(*s)[:len(*s)-1]
	return n
}

/**
迭代
左根右
*/
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	rest := []int{}
	stack := Stack([]*TreeNode{root})
	for len(stack) > 0 {
		cur := stack.Pop()
		if cur != nil {
			if cur.Right != nil {
				stack.Push(cur.Right)
			}
			stack.Push(cur)
			stack.Push(nil)
			if cur.Left != nil {
				stack.Push(cur.Left)
			}
		}else{
			rest = append(rest,stack.Pop().Val)
		}
	}
	return rest
}
func main()  {
	root := &TreeNode{
		Val:1,
		Left:nil,
		Right:&TreeNode{
			Val:2,
			Left:&TreeNode{
				Val:3,
				Left:nil,
				Right:nil,
			},
			Right:nil,
		},
	}
	fmt.Println(inorderTraversal(root))
}
