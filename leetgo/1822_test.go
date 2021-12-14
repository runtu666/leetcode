package main

import (
	"fmt"
	"testing"
)

/**
已知函数signFunc(x) 将会根据 x 的正负返回特定值：

如果 x 是正数，返回 1 。
如果 x 是负数，返回 -1 。
如果 x 是等于 0 ，返回 0 。
给你一个整数数组 nums 。令 product 为数组 nums 中所有元素值的乘积。

返回 signFunc(product) 。

示例 1：

输入：nums = [-1,-2,-3,-4,3,2,1]
输出：1
解释：数组中所有值的乘积是 144 ，且 signFunc(144) = 1
示例 2：

输入：nums = [1,5,0,2,-3]
输出：0
解释：数组中所有值的乘积是 0 ，且 signFunc(0) = 0
示例 3：

输入：nums = [-1,1,-1,1,-1]
输出：-1
解释：数组中所有值的乘积是 -1 ，且 signFunc(-1) = -1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sign-of-the-product-of-an-array
*/
func arraySign(nums []int) int {
	var lt int
	for _, n := range nums {
		if n == 0 {
			return 0
		}
		if n < 0 {
			lt++
		}
	}
	i := lt % 2
	if i > 0 {
		return -1
	}
	return 1
}

func Test1822(t *testing.T) {
	nums := []int{-1, -2, -3, -4, 3, 2, 1}
	fmt.Println(arraySign(nums))
}
