package main

import (
	"fmt"
	"testing"
)

/**
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回-1 。

你可以认为每种硬币的数量是无限的。



示例1：

输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1
示例 2：

输入：coins = [2], amount = 3
输出：-1
示例 3：

输入：coins = [1], amount = 0
输出：0
示例 4：

输入：coins = [1], amount = 1
输出：1
示例 5：

输入：coins = [1], amount = 2
输出：2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/coin-change
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test332(t *testing.T) {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	for i := 0; i <= amount; i++ {
		for _, v := range coins {
			if i-v < 0 {
				continue
			}
			//fmt.Println(dp[i])
			if dp[i] > dp[i-v]+1 {
				dp[i] = dp[i-v] + 1
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
