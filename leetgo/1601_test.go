package main

import (
	"testing"
)

// https://leetcode-cn.com/problems/maximum-number-of-achievable-transfer-requests/
func Test1601(t *testing.T) {
	t.Log(maximumRequests(5, [][]int{{0, 1}, {1, 0}, {0, 1}, {1, 2}, {2, 0}, {3, 4}}))
}

func maximumRequests(n int, requests [][]int) (ans int) {
	delta := make([]int, n)
	cnt, zero := 0, n
	var dfs func(int)
	dfs = func(pos int) {
		if pos == len(requests) {
			if zero == n && cnt > ans {
				ans = cnt
			}
			return
		}

		// 不选 requests[pos]
		dfs(pos + 1)

		// 选 requests[pos]
		z := zero
		cnt++
		r := requests[pos]
		x, y := r[0], r[1]
		if delta[x] == 0 {
			zero--
		}
		delta[x]--
		if delta[x] == 0 {
			zero++
		}
		if delta[y] == 0 {
			zero--
		}
		delta[y]++
		if delta[y] == 0 {
			zero++
		}
		dfs(pos + 1)
		delta[y]--
		delta[x]++
		cnt--
		zero = z
	}
	dfs(0)
	return
}
