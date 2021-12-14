package main

import (
	. "leetcode/common"
	"testing"
)

func Test230(t *testing.T) {

}

func kthSmallest(root *TreeNode, k int) int {
	s := make([]*TreeNode, 0)
	if root == nil {
		return 0
	}
	c := root
	for c != nil || len(s) != 0 {
		if c != nil {
			s = append(s, c)
			c = c.Left
		} else {
			tmp := s[len(s)-1]
			s = s[:len(s)-1]
			k--
			if k == 0 {
				return tmp.Val
			}
			c = tmp.Right
		}
	}
	return 0
}
