package common

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetListNode(arr []int) *ListNode {
	var h = new(ListNode)
	curr := h
	for _, a := range arr {
		curr.Next = &ListNode{Val: a}
		curr = curr.Next
	}
	return h.Next
}

func (l *ListNode) ToArr() []int {
	var c = l
	var r []int
	for c != nil {
		r = append(r, c.Val)
		c = c.Next
	}
	return r
}
