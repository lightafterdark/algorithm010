package main

import (
	"fmt"
)

//面试题 02.01. 移除重复节点
//type ListNode struct {
//	Val int
//	Next *ListNode
//}

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	res := make(map[int]bool)
	res[head.Val] = true
	pos := head

	for pos.Next != nil {
		cur := pos.Next
		if !res[cur.Val] {
			res[cur.Val] = true
			fmt.Println(pos)
			pos = pos.Next
			fmt.Println(pos)
		}else{
			pos.Next = pos.Next.Next
		}
	}
	pos.Next =  nil
	return head
}

func main()  {
	head := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:1,
			Next:&ListNode{
				Val:2,
				Next:nil,
			},
		},
	}
	fmt.Println(removeDuplicateNodes(head))
}

