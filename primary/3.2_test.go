package main

import (
	. "leetcode/common"
	"testing"
)

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn2925/
func Test3_2(t *testing.T) {
	h := GetListNode([]int{1, 2, 3, 4, 5})
	t.Log(removeNthFromEnd(h, 2))
	t.Log(removeNthFromEnd(GetListNode([]int{1, 2}), 1))
	t.Log(removeNthFromEnd(GetListNode([]int{1}), 1))
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var curr = head
	var x = 0
	for curr != nil {
		x++
		curr = curr.Next
	}
	var pos = x - n
	if pos == 0 {
		return head.Next
	}
	curr = head
	for i := 1; i < pos; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
	return head
}
