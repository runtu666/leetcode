package main

import (
	"strings"
	"testing"
)

/*
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
解释："amanaplanacanalpanama" 是回文串
示例 2:

输入: "race a car"
输出: false
解释："raceacar" 不是回文串

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xne8id/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func Test2_5(t *testing.T) {
	t.Log(isPalindrome("A man, a plan, a canal: Panama"))
	t.Log(isPalindrome("race a car"))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var l, r = 0, len(s) - 1
	for l < r {
		if s[l] < '0' || s[l] > 'z' || (s[l] > '9' && s[l] < 'a') {
			l++
			continue
		}
		if s[r] < '0' || s[r] > 'z' || (s[r] > '9' && s[r] < 'a') {
			r--
			continue
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
