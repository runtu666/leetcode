package main

import (
	"fmt"
)

/**
2. 两数相加
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
链接：https://leetcode-cn.com/problems/add-two-numbers/
*/

type ListNode struct {
	Next *ListNode
	Val  int
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var tail, res *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if res == nil {
			res = &ListNode{
				Val: sum,
			}
			tail = res
		} else {
			tail.Next = &ListNode{
				Val: sum,
			}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{
			Val: carry,
		}
	}
	return res
}

func main() {
	l1 := ListNode{
		Next: &ListNode{
			Next: &ListNode{
				Next: nil,
				Val:  4,
			},
			Val: 6,
		},
		Val: 5,
	}

	l2 := ListNode{
		Next: &ListNode{
			Next: &ListNode{
				Next: nil,
				Val:  5,
			},
			Val: 4,
		},
		Val: 2,
	}

	fmt.Println(addTwoNumbers(&l1, &l2))
}
