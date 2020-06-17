package main
//589. N叉树的前序遍历

type Node struct {
	Val int
	Children []*Node
}

/**
递归
 */
func preorder1(root *Node) []int {
	rest := []int{}
	if root == nil {
		return rest
	}

	rest = append(rest,root.Val)
	for _,n := range root.Children{
		rest = append(rest,preorder1(n)...)
	}
	return rest
}

/**
迭代
 */
func preorder2(root *Node) []int {
	if root == nil {
		return []int{}
	}
	rest := []int{}
	stack := make([]*Node,0)
	stack = append(stack,root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		rest = append(rest,node.Val)
		for i:=len(node.Children)-1;i>=0 ;i--  {
			stack = append(stack,node.Children[i])
		}
	}
	return rest
}





