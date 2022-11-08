package main

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/count-unique-characters-of-all-substrings-of-a-given-string/
func Test828(t *testing.T) {
	t.Log(uniqueLetterString("count"))
}

func uniqueLetterString(s string) int {
	var m = make(map[string]int)
	l := len(s)
	for i := 1; i <= l; i++ {
		for j := 0; j < l; j++ {
			if j+i > l {
				break
			}
			fmt.Println(s[j : j+i])
			m[s[j:j+i]]++
		}
	}
	var ans int
	for k, v := range m {
		if v > 1 {
			//fmt.Println(k)
			continue
		}
		ans += len(k)
	}
	//fmt.Println(m)
	//fmt.Println(len(m))
	return ans
}
