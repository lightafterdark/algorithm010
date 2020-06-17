package main

import (
	"fmt"
	"strconv"
	"strings"
)

//297. 二叉树的序列化与反序列化
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

type Codec struct {
	l []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize1(root *TreeNode) string {
	if root == nil {
		return "null"
	}
	return strconv.Itoa(root.Val)+","+this.serialize(root.Left)+","+this.serialize(root.Right)
}

func (this *Codec) serialize(root *TreeNode) string {
	return this.serializeHelper(root,"")
}

func (this *Codec)  serializeHelper(root *TreeNode,str string) string{
	if root == nil {
		str += "null,"
		return str
	}else{
		str += strconv.Itoa(root.Val) + ","
		str = this.serializeHelper(root.Left,str)
		str = this.serializeHelper(root.Right,str)
	}
	return str
}


func (this *Codec) deserializeHelper() *TreeNode{
	if this.l[0] == "null" {
		this.l = this.l[1:]
		return nil
	}

	val,_ := strconv.Atoi(this.l[0])
	root := &TreeNode{Val:val}
	this.l = this.l[1:]
	root.Left = this.deserializeHelper()
	root.Right = this.deserializeHelper()
	return root
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	l := strings.Split(data,",")
	for i:=0;i<len(l) ;i++  {
		if l[i] != "" {
			this.l = append(this.l,l[i])
		}
	}
	return this.deserializeHelper()
}

func main()  {
	root := &TreeNode{
		Val:1,
		Left:&TreeNode{
			Val:2,
			Right:nil,
			Left:nil,
		},
		Right:&TreeNode{
			Val:3,
			Left:&TreeNode{
				Val:4,
				Left:nil,
				Right:nil,
			},
			Right:&TreeNode{
				Val:5,
				Left:nil,
				Right:nil,
			},
		},
	}
	obj := Constructor();
	data := obj.serialize(root)
	fmt.Println(data)
	//ans := obj.deserialize(data)
	//fmt.Println(ans)
}