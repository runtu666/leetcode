package main

import (
	"fmt"
	"testing"
)

/**
给你一个整数n，请你返回所有 0 到 1 之间（不包括 0 和 1）满足分母小于等于n的 最简分数。分数可以以 任意顺序返回。
示例 1：

输入：n = 2
输出：["1/2"]
解释："1/2" 是唯一一个分母小于等于 2 的最简分数。
示例 2：

输入：n = 3
输出：["1/2","1/3","2/3"]
示例 3：

输入：n = 4
输出：["1/2","1/3","1/4","2/3","3/4"]
解释："2/4" 不是最简分数，因为它可以化简为 "1/2" 。
示例 4：

输入：n = 1
输出：[]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/simplified-fractions
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test1447(t *testing.T) {
	t.Log(simplifiedFractions(1))
	t.Log(simplifiedFractions(2))
	t.Log(simplifiedFractions(3))
	t.Log(simplifiedFractions(4))
}

func simplifiedFractions(n int) []string {
	var ans = make([]string, 0)
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			if isOK(i, j) {
				ans = append(ans, fmt.Sprintf("%d/%d", j, i))
			}
		}
	}
	return ans
}

func isOK(x, y int) bool {
	for x*y > 0 {
		if x > y {
			x %= y
		} else if x < y {
			y %= x
		}
	}
	var d = y
	if x > y {
		d = x
	}
	return d == 1
}
