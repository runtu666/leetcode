package main

import (
	"testing"
)

/**
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

示例 1：
输入：nums = [1,1,2]
输出：2, nums = [1,2]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
示例 2：
输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。


作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2gy9m/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func Test1_1(t *testing.T) {
	var nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	t.Log(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var l = 0
	for r := 1; r < len(nums); r++ {
		if nums[l] != nums[r] {
			l += 1
			nums[l] = nums[r]
		}
	}
	return l + 1
}
