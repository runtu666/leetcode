package main

import (
	. "leetcode/common"
	"testing"
)

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnnbp2/
func Test3_4(t *testing.T) {
	l1 := GetListNode([]int{1, 2, 4, 5})
	l2 := GetListNode([]int{1, 3, 4})
	t.Log(mergeTwoLists(l1, l2).ToArr())
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var node = new(ListNode)
	curr := node
	for {
		if list1 == nil {
			curr.Next = list2
			break
		}
		if list2 == nil {
			curr.Next = list1
			break
		}
		if list1.Val > list2.Val {
			curr.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		} else {
			curr.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		}
		curr = curr.Next

	}
	return node.Next
}
