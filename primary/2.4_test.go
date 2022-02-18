package main

import (
	"testing"
)

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若s 和 t中每个字符出现的次数都相同，则称s 和 t互为字母异位词。

示例1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn96us/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func Test2_4(t *testing.T) {
	t.Log(isAnagram("anagram", "nagaram"))
	t.Log(isAnagram("rat", "car"))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var getMap = func(s string) map[int32]int {
		var m = make(map[int32]int)
		for _, v := range s {
			if _, ok := m[v]; ok {
				m[v]++
				continue
			}
			m[v] = 1
		}
		return m
	}
	m1 := getMap(s)
	m2 := getMap(t)
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}
