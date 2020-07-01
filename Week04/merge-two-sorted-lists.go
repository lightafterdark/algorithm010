package main

import (
	"fmt"
)

//21. 合并两个有序链表
type ListNode struct {
	Val int
	Next *ListNode
}

/**
递归，
谁的第一个元素小，后面的元素就往谁的next里面添加
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}else if l2 == nil {
		return l1
	}else if l1.Val < l2.Val{
		l1.Next = mergeTwoLists(l1.Next,l2)
		return l1
	}else{
		l2.Next = mergeTwoLists(l1,l2.Next)
		return l2
	}
}


func main()  {
	l1 := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:2,
			Next:&ListNode{
				Val:4,
				Next:nil,
			},
		},
	}
	l2 := &ListNode{
		Val:1,
		Next:&ListNode{
			Val:3,
			Next:&ListNode{
				Val:4,
				Next:nil,
			},
		},
	}
	fmt.Println(mergeTwoLists(l1,l2))
}
