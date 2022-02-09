package main

import (
	"testing"
)

/**
给你一个整数数组nums和一个整数k，请你返回数对(i, j)的数目，满足i < j且|nums[i] - nums[j]| == k。

|x|的值定义为：

如果x >= 0，那么值为x。
如果x < 0，那么值为-x。


示例 1：

输入：nums = [1,2,2,1], k = 1
输出：4
解释：差的绝对值为 1 的数对为：
- [1,2,2,1]
- [1,2,2,1]
- [1,2,2,1]
- [1,2,2,1]
示例 2：

输入：nums = [1,3], k = 3
输出：0
解释：没有任何数对差的绝对值为 3 。
示例 3：

输入：nums = [3,2,1,5,4], k = 2
输出：3
解释：差的绝对值为 2 的数对为：
- [3,2,1,5,4]
- [3,2,1,5,4]
- [3,2,1,5,4]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-number-of-pairs-with-absolute-difference-k
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test2006(t *testing.T) {
	t.Log(countKDifference([]int{1, 2, 2, 1}, 1))
	t.Log(countKDifference([]int{1, 3}, 3))
	t.Log(countKDifference([]int{3, 2, 1, 5, 4}, 2))
}

func countKDifference(nums []int, k int) int {
	if len(nums) < 2 {
		return 0
	}
	var ans = 0
	for i, l := 0, len(nums); i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]-nums[j] == k || nums[j]-nums[i] == k {
				ans++
			}
		}
	}
	return ans
}
