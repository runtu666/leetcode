package main

import (
	"testing"
)

//https://leetcode-cn.com/problems/where-will-the-ball-fall/
func Test1706(t *testing.T) {
	t.Log(findBall([][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}}))
	t.Log(findBall([][]int{{-1}}))
	t.Log(findBall([][]int{{1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}, {1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}}))
}

func findBall(grid [][]int) []int {
	var xl = len(grid)
	var yl = len(grid[0])
	var ans = make([]int, yl)
	for y := range ans {
		var x = 0
		var z = y
		for x < xl && z < yl {
			if grid[x][z] > 0 {
				if z+1 >= yl {
					break
				}
				if grid[x][z+1] < 0 {
					break
				}
				z++
			}
			if grid[x][z] < 0 {
				if z-1 < 0 {
					break
				}
				if grid[x][z-1] > 0 {
					break
				}
				z--
			}
			x++
		}
		if x == xl {
			ans[y] = z
			continue
		}
		ans[y] = -1
	}

	return ans
}
