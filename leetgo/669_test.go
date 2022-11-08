package main

import (
	. "leetcode/common"
	"testing"
)

// https://leetcode.cn/problems/trim-a-binary-search-tree/
func Test669(t *testing.T) {

}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	if root == nil {
		return nil
	}
	for curr := root; curr.Left != nil; {
		if curr.Left.Val < low {
			curr.Left = curr.Left.Right
		} else {
			curr = curr.Left
		}
	}

	for curr := root; curr.Right != nil; {
		if curr.Right.Val > high {
			curr.Right = curr.Right.Left
		} else {
			curr = curr.Right
		}
	}
	return root
}
