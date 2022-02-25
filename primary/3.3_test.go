package main

import (
	. "leetcode/common"
	"testing"
)

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnnhm6/
func Test3_3(t *testing.T) {
	t.Log(reverseList(GetListNode([]int{1, 2, 3})))
}

func reverseList(head *ListNode) *ListNode {
	var curr = head
	var pre *ListNode
	for curr != nil {
		curr.Next, pre, curr = pre, curr, curr.Next
	}

	return pre
}
