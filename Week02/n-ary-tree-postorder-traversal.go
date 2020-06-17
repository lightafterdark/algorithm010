package main
//590. N叉树的后序遍历

func postorder1(root *Node) []int {
	rest := []int{}
	if root == nil {
		return rest
	}

	for _,n := range root.Children {
		rest = append(rest,postorder1(n)...)
	}
	rest = append(rest,root.Val)
	return rest
}


func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	rest := []int{}
	stack := make([]*Node,0)
	stack = append(stack,root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		rest = append([]int{node.Val},rest...)
		for i:=0;i<len(node.Children) ;i++  {
			stack = append(stack,node.Children[i])
		}
	}
	return rest
}
