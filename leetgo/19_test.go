package main

import (
	"testing"

	. "leetcode/common"
)

func Test19(t *testing.T) {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var ans = &ListNode{}
	ans.Next = head
	var p1 = ans
	var p2 = ans
	for i := 0; i < n+1; i++ {
		p1 = p1.Next
	}
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	p2.Next = p2.Next.Next

	return ans.Next
}
