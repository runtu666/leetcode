package main

import (
	"sort"
	"testing"

	. "leetcode/common"
)

func Test654(t *testing.T) {
	constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
}
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var m = make(map[int]int)
	for i, num := range nums {
		m[num] = i
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	root := new(TreeNode)
	curr := root
	curr.Val = nums[0]
	for _, num := range nums[1:] {
		insert(curr, &TreeNode{
			Val: num,
		}, m)
	}
	return root
}

func insert(curr, newNode *TreeNode, m map[int]int) {
	if m[newNode.Val] > m[curr.Val] {
		InsertRight(curr, newNode, m)
	} else {
		InsertLeft(curr, newNode, m)
	}
}

func InsertLeft(curr, newNode *TreeNode, m map[int]int) {
	if curr.Left == nil {
		curr.Left = newNode
	} else {
		insert(curr.Left, newNode, m)
	}
}

func InsertRight(curr, newNode *TreeNode, m map[int]int) {
	if curr.Right == nil {
		curr.Right = newNode
	} else {
		insert(curr.Right, newNode, m)
	}
}
