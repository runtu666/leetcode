package main

import (
	"testing"
)

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2ba4i/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func Test1_7(t *testing.T) {
	moveZeroes([]int{0, 1, 0, 3, 12})
}

func moveZeroes(nums []int) {
	var i, j = 0, 0
	for l := len(nums); j < l; j++ {
		if nums[j] == 0 {
			i++
			continue
		}
		if i == 0 {
			continue
		}
		nums[j-i] = nums[j]
		nums[j] = 0
	}
}
