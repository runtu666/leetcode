package main

import (
	"sort"
	"testing"
)

func Test1984(t *testing.T) {
	t.Log(minimumDifference([]int{90}, 1))
	t.Log(minimumDifference([]int{9, 4, 1, 7}, 2))
}

func minimumDifference(nums []int, k int) int {
	var l = len(nums)
	if k <= 1 || l <= 1 {
		return 0
	}
	sort.Ints(nums)
	var ans = nums[l-1]
	for i := range nums {
		if i+k > l {
			break
		}
		a := nums[i+k-1] - nums[i]
		if a < ans {
			ans = a
		}
		if ans == 0 {
			return ans
		}
	}
	return ans
}
