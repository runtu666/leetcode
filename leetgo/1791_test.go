package main

import (
	"testing"
)

func Test1791(t *testing.T) {
	t.Log(findCenter([][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}))
}

func findCenter(edges [][]int) int {
	var m = make(map[int]struct{})
	for _, edge := range edges[:2] {
		for _, i := range edge {
			if _, ok := m[i]; ok {
				return i
			}
			m[i] = struct{}{}
		}
	}
	return 0
}
