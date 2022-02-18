package main

import (
	"testing"
)

/**
给你两个整数数组nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。



示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]
示例 2:

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2y0c2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func Test1_6(t *testing.T) {
	t.Log(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}

func intersect(nums1 []int, nums2 []int) []int {
	m1 := getMap(nums1)
	m2 := getMap(nums2)
	var ans []int
	for k, c := range m1 {
		if _, ok := m2[k]; !ok {
			continue
		}
		var count = c
		if c > m2[k] {
			count = m2[k]
		}
		for i := 0; i < count; i++ {
			ans = append(ans, k)
		}
	}
	return ans
}

func getMap(nums []int) map[int]int {
	var m = make(map[int]int)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			m[num]++
			continue
		}
		m[num] = 1
	}
	return m
}
