package main

import (
	"testing"
)

func Test1464(t *testing.T) {
	t.Log(maxProduct([]int{3, 4, 5, 2}))
}

func maxProduct(nums []int) int {
	var a, b int
	for _, num := range nums {
		if num > b {
			if num > a {
				b, a = a, num
			} else {
				b = num
			}
		}
	}
	return (a - 1) * (b - 1)
}
