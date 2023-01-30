package main

import (
	"testing"

	. "leetcode/common"
)

func Test1669(t *testing.T) {

}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	curr := list1
	i := 0
	for {
		if i == b {
			curr2 := list2
			for curr2.Next != nil {
				curr2 = curr2.Next
			}
			curr2.Next = curr.Next
			break
		}
		curr = curr.Next
		i++
	}
	i = 0
	curr = list1
	for {
		if i == a-1 {
			curr.Next = list2
			break
		}
		curr = curr.Next
		i++
	}
	return list1
}
