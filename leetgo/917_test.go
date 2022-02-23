package main

import (
	"testing"
)

/**
给你一个字符串 s ，根据下述规则反转字符串：

所有非英文字母保留在原有位置。
所有英文字母（小写或大写）位置反转。
返回反转后的 s 。

示例 1：

输入：s = "ab-cd"
输出："dc-ba"
示例 2：

输入：s = "a-bC-dEf-ghIj"
输出："j-Ih-gfE-dCba"
示例 3：

输入：s = "Test1ng-Leet=code-Q!"
输出："Qedo1ct-eeLg=ntse-T!"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-only-letters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test917(t *testing.T) {
	t.Log(reverseOnlyLetters("ab-cd"))
	t.Log(reverseOnlyLetters("Test1ng-Leet=code-Q!"))
	t.Log(reverseOnlyLetters("7_28]"))
}

func reverseOnlyLetters(s string) string {
	var l, r = 0, len(s) - 1
	var sb = []byte(s)
	for l < r {
		if !isAlph(s[l]) {
			l++
			continue
		}
		if !isAlph(s[r]) {
			r--
			continue
		}
		sb[l], sb[r] = sb[r], sb[l]
		l++
		r--
	}
	return string(sb)
}

func isAlph(b byte) bool {
	if b >= 'A' && b <= 'Z' {
		return true
	}
	if b >= 'a' && b <= 'z' {
		return true
	}
	return false
}
