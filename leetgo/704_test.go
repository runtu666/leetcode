package main

import (
	"testing"
)

func Test704(t *testing.T) {
	t.Log(search([]int{-1, 0, 3, 5, 9, 12}, 2))
	t.Log(search([]int{5}, 5))
}

func search(nums []int, target int) int {
	var left, right = 0, len(nums) - 1
	for right >= left {
		mid := (right-left)/2 + left
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func firstBadVersion(n int) int {
	var left, right = 0, n
	var minVersion = n
	for right > left {
		ver := (right-left)/2 + left
		if isBadVersion(ver) {
			minVersion = ver
			left = ver + 1
		} else {
			right = ver - 1
		}
	}
	return minVersion
}

func isBadVersion(version int) bool {
	return false
}
