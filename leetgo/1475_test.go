package main

import "testing"

// https://leetcode.cn/problems/final-prices-with-a-special-discount-in-a-shop/
func Test1475(t *testing.T) {
	t.Log(finalPrices([]int{1}))
}

func finalPrices(prices []int) []int {
	l := len(prices)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if prices[i] >= prices[j] {
				prices[i] = prices[i] - prices[j]
				break
			}
		}
	}
	return prices
}
