package main
//429. N叉树的层序遍历

var res [][]int
func levelOrder(root *Node) [][]int {
	res = [][]int{}
	if root == nil {
		return res
	}
	dfs(root,0)
	return res
}

func dfs(root *Node,level int) {
	if root == nil {
		return
	}
	if len(res) == level{
		res = append(res,[]int{})
	}
	res[level] = append(res[level],root.Val)
	for _,n := range root.Children {
		dfs(n,level+1)
	}
}