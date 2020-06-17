package main

import "fmt"

//145. 二叉树的后序遍历

/**
递归
 */
func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	rest := append(postorderTraversal(root.Left),postorderTraversal(root.Right)...)
	rest = append(rest,root.Val)
	return rest
}


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

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
左右根
*/
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	rest := []int{}
	stack := Stack([]*TreeNode{root})
	for len(stack) > 0 {
		cur := stack.Pop()
		if cur != nil {
			stack.Push(cur)
			stack.Push(nil)
			if cur.Right != nil {
				stack.Push(cur.Right)
			}
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
	fmt.Println(postorderTraversal(root))
}
