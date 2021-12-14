package main

import (
	. "leetcode/common"
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func increasingBST(root *TreeNode) *TreeNode {
	var sort = new(TreeNode)
	var inOrderTraverse func(node *TreeNode)
	var vals []int
	inOrderTraverse = func(node *TreeNode) {
		if node != nil {
			inOrderTraverse(node.Left)
			vals = append(vals, node.Val)
			inOrderTraverse(node.Right)
		}
	}
	inOrderTraverse(root)
	curNode := sort
	for _, val := range vals {
		curNode.Right = &TreeNode{
			Val: val,
		}
		curNode = curNode.Right
	}
	return sort.Right
}

func buildTree() *TreeNode {
	var vals = []int{5, 3, 6, 2, 4, 0, 8, 1, 0, 0, 0, 7, 9}
	var root = new(TreeNode)
	root.Val = 5
	for _, val := range vals {
		if val == 0 {
			continue
		}
		Insert(root, &TreeNode{Val: val})
	}
	return root
}

func Insert(nd, newNode *TreeNode) {
	if newNode.Val == nd.Val {
		return
	}
	if newNode.Val > nd.Val {
		if nd.Right == nil {
			nd.Right = newNode
		} else {
			Insert(nd.Right, newNode)
		}
	} else { //3 小于 继续比较插入到 左孩子
		if nd.Left == nil {
			nd.Left = newNode
		} else {
			Insert(nd.Left, newNode)
		}
	}
}
func Test897(t *testing.T) {
	increasingBST(buildTree())
}
