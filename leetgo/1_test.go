package main

import (
	"fmt"
	"testing"
)

/**
1. 两数之和
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

链接：https://leetcode-cn.com/problems/two-sum/
*/
func twoSum(nums []int, target int) []int {
	var res = make([]int, 0)
	for i, m := range nums {
		for j, n := range nums {
			if j == i {
				continue
			}
			if m+n == target {
				res = append(res, i, j)
				return res
			}
		}
	}
	return res
}

func Test1(t *testing.T) {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
