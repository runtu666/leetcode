package main

import (
	"fmt"
	"strings"
)

/**
290. 单词规律
给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。

示例1:

输入: pattern = "abba", str = "dog cat cat dog"
输出: true
示例 2:

输入:pattern = "abba", str = "dog cat cat fish"
输出: false
示例 3:

输入: pattern = "aaaa", str = "dog cat cat dog"
输出: false
示例 4:

输入: pattern = "abba", str = "dog dog dog dog"
输出: false
说明:
你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。
*/

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
}

func wordPattern(pattern string, s string) bool {
	sList := strings.Split(s, " ")
	patternList := strings.Split(pattern, "")

	if len(sList) != len(patternList) {
		return false
	}
	var m1 = make(map[string]string)
	var m2 = make(map[string]string)
	for i, p := range patternList {

		if v, ok := m1[p]; ok && v != sList[i] {
			return false
		}
		if v, ok := m2[sList[i]]; ok && v != p {
			return false
		}
		m1[p] = sList[i]
		m2[sList[i]] = p
	}
	return true
}
