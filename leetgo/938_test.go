package main

import (
	. "leetcode/common"
	"testing"
)

/**
给定二叉搜索树的根结点root，返回值位于范围 [low, high] 之间的所有结点的值的和。

输入：root = [10,5,15,3,7,null,18], low = 7, high = 15
输出：32
示例 2：


输入：root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
输出：23

提示：

树中节点数目在范围 [1, 2 * 104] 内
1 <= Node.val <= 105
1 <= low <= high <= 105
所有 Node.val 互不相同

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/range-sum-of-bst
*/

func Test938(t *testing.T) {
	tree := BuildTree([]int{10, 5, 15, 3, 7, 0, 18})
	rangeSumBST(tree, 7, 15)
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	var tra func(node *TreeNode, low int, high int)
	var vals []int
	tra = func(node *TreeNode, low int, high int) {
		if node != nil {
			if node.Val >= low {
				tra(node.Left, low, high)
			}
			if node.Val >= low && node.Val <= high {
				vals = append(vals, node.Val)
			}
			if node.Val <= high {
				tra(node.Right, low, high)
			}
		}
	}
	tra(root, low, high)
	var val int
	for _, v := range vals {
		val += v
	}
	return val
}
