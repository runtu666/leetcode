package main

import (
	. "leetcode/common"
	"testing"
)

func Test3_1(t *testing.T) {

}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
