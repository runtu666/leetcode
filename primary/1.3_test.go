package main

import (
	"fmt"
	"testing"
)

/*
给你一个数组，将数组中的元素向右轮转 k个位置，其中k是非负数。



示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
示例2:

输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2skh7/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func Test1(t *testing.T) {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}

func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	if k == 0 {
		fmt.Println(nums)
		return
	}

	for i := 0; i < l; i++ {
		index := i + k
		if i+k >= l {
			index = index - l
		}
		tmp := nums[i+1]
		nums[i] = tmp
	}
	fmt.Println(nums)
}
