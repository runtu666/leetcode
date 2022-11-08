package main

import (
	. "leetcode/common"
	"testing"
)

// https://leetcode.cn/problems/maximum-width-of-binary-tree/
func Test662(t *testing.T) {

}

func widthOfBinaryTree(root *TreeNode) int {
	var fn func(node *TreeNode, index, level int64)
	type x struct {
		max int64
		min int64
	}
	var m = make(map[int64]*x)
	fn = func(node *TreeNode, index, level int64) {
		if node == nil {
			return
		}
		if _, ok := m[level]; !ok {
			m[level] = &x{
				max: 0,
				min: index,
			}
		}
		if index > m[level].max {
			m[level].max = index
		}
		if index < m[level].min {
			m[level].min = index
		}
		fn(node.Left, 1, level+1)
		fn(node.Right, 2, level+1)
	}
	fn(root, 1, 1)
	var max int64
	for _, v := range m {
		i := v.max - v.min
		if i > max {
			max = i
		}
	}
	return int(max + 1)
}
