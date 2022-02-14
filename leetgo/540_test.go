package main

import "testing"

func Test540(t *testing.T) {
	t.Log(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	t.Log(singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))
}

func singleNonDuplicate(nums []int) int {
	var ans = 0
	for _, num := range nums {
		ans = ans ^ num
	}
	return ans
}
