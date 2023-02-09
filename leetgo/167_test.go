package main

import (
	"fmt"
	"testing"
)

func Test167(t *testing.T) {
	t.Log(twoSum167([]int{5, 25, 75}, 100))
}

func twoSum167(numbers []int, target int) []int {
	var l, length = 0, len(numbers)
	var r = l + 1
	for l < length {
		fmt.Println(l, r)
		if numbers[l]+numbers[r] == target {
			break
		} else if numbers[l]+numbers[r] < target {
			r++
		}
		if r == length || numbers[l]+numbers[r] > target {
			l++
			r = l + 1
		}
	}

	return []int{l + 1, r + 1}
}
