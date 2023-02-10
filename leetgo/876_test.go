package main

import (
	"testing"

	. "leetcode/common"
)

func Test876(t *testing.T) {

}

func middleNode(head *ListNode) *ListNode {
	var cur = head
	var l int
	for cur != nil {
		cur = cur.Next
		l++
	}
	var ans = head
	for i := 0; i < l/2; i++ {
		ans = ans.Next
	}
	return ans
}
