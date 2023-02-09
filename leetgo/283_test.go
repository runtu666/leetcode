package main

import "testing"

func Test283(t *testing.T) {
	var nums = []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	t.Log(nums)
}

func moveZeroes(nums []int) {
	var left, right, l = 0, 0, len(nums)
	for right < l {
		if nums[right] != 0 {
			nums[right], nums[left] = nums[left], nums[right]
			left++
		}
		right++
	}
}
