package main

import "fmt"

/*
给定一个大小为n的整数数组，找出其中所有出现超过⌊ n/3 ⌋次的元素。

示例1：
输入：[3,2,3]
输出：[3]
示例 2：
输入：nums = [1]
输出：[1]
示例 3：
输入：[1,1,1,3,3,2,2,2]
输出：[1,2]


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/majority-element-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Println(majorityElement([]int{1, 1, 1, 3, 3, 2, 2, 2}))
}

func majorityElement(nums []int) []int {
	l := len(nums)
	t := l / 3
	var m = make(map[int]int)
	var rs = make([]int, 0)
	for _, n := range nums {
		m[n] += 1
		if m[n] == t+1 {
			rs = append(rs, n)
		}
	}
	return rs
}
