package main

import "testing"

func Test1_4(t *testing.T) {
	t.Log(containsDuplicate([]int{1, 3, 4, 5, 1}))
}

func containsDuplicate(nums []int) bool {
	var m = make(map[int]struct{})
	for _, n := range nums {
		if _, ok := m[n]; ok {
			return true
		}
		m[n] = struct{}{}
	}
	return false
}
