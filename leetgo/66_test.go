package main

import (
	"fmt"
	"testing"
)

/*
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。

示例1：
输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。

示例2：
输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。

示例 3：
输入：digits = [0]
输出：[1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/plus-one
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test66(t *testing.T) {
	fmt.Println(plusOne([]int{9}))
}

func plusOne(digits []int) []int {
	l := len(digits) - 1
	for {
		if digits[l] == 9 {
			digits[l] = 0
			if l == 0 {
				digits = append([]int{1}, digits...)
				break
			}
			l--
			continue
		}
		digits[l] += 1
		break
	}
	return digits
}
