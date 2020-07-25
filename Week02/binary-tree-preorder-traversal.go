package main

import "fmt"

//144. 二叉树的前序遍历
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
递归
*/
func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	res = append(res, root.Val)
	res = append(res, preorderTraversal1(root.Left)...)
	res = append(res, preorderTraversal1(root.Right)...)
	return res
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
根左右
*/
func preorderTraversal(root *TreeNode) []int {
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
			if cur.Left != nil {
				stack.Push(cur.Left)
			}
			stack.Push(cur)
			stack.Push(nil)
		} else {
			rest = append(rest, stack.Pop().Val)
		}
	}
	return rest
}

func preorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	return res
}

func main() {
	root := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Println(preorderTraversal(root))
}
