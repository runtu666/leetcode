package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

/**
给定一个二叉树的根节点 root ，返回它的 中序 遍历。
*/

var rs = make([]int, 0)

func Test94(t *testing.T) {
	nodeG := TreeNode{Val: 7, Left: nil, Right: nil}
	nodeF := TreeNode{Val: 6, Left: &nodeG, Right: nil}
	nodeE := TreeNode{Val: 5, Left: nil, Right: nil}
	nodeD := TreeNode{Val: 4, Left: &nodeE, Right: nil}
	nodeC := TreeNode{Val: 3, Left: nil, Right: nil}
	nodeB := TreeNode{Val: 2, Left: &nodeD, Right: &nodeF}
	nodeA := &TreeNode{Val: 1, Left: &nodeB, Right: &nodeC}

	f(nodeA)
	fmt.Println(rs)
	fmt.Println(inorderTraversal(nodeA))
}

func f(root *TreeNode) {
	if root != nil {
		f(root.Left)
		rs = append(rs, root.Val)
		f(root.Right)
	}
}

func inorderTraversal(root *TreeNode) []int {
	var s = make([]*TreeNode, 0)
	var rs = make([]int, 0)
	c := root
	for c != nil || len(s) != 0 {
		if c != nil {
			s = append(s, c)
			c = c.Left
		} else {
			t := s[len(s)-1]
			s = s[:len(s)-1]
			rs = append(rs, t.Val)

			c = t.Right
		}
	}

	return rs
}
