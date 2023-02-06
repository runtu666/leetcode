package main

import (
	"testing"

	. "leetcode/common"
)

func Test2331(t *testing.T) {

}

func evaluateTree(root *TreeNode) bool {
	switch root.Val {
	case 0:
		return false
	case 1:
		return true
	case 2:
		return evaluateTree(root.Right) || evaluateTree(root.Left)
	case 3:
		return evaluateTree(root.Right) && evaluateTree(root.Left)
	}
	return false
}
