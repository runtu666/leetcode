package main

import (
	"testing"
)

func Test2319(t *testing.T) {
	t.Log(checkXMatrix([][]int{
		{2, 0, 0, 1},
		{0, 3, 1, 0},
		{0, 5, 2, 0},
		{4, 0, 0, 2},
	}))
}

func checkXMatrix(grid [][]int) bool {
	var x = len(grid)
	for i, list := range grid {
		for j, num := range list {
			if (i == j || i+j == x-1) && num == 0 {
				return false
			}
			if i != j && i+j != x-1 && num != 0 {
				return false
			}
		}
	}
	return true
}
