package main

import (
	"testing"
)

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnmav1/
func Test2_9(t *testing.T) {
	t.Log(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}
	if length == 1 {
		return strs[0]
	}
	var getAns = func(a, b string) string {
		var la, lb = len(a), len(b)
		if len(a) == 0 || len(b) == 0 {
			return ""
		}
		if la > lb {
			a, b = b, a
		}
		for i := range a {
			if a[i] != b[i] {
				return a[:i]
			}
		}
		return a
	}

	var l = 1
	var ans = strs[0]
	for l < length {
		ans = getAns(ans, strs[l])
		l++
	}

	return ans
}
