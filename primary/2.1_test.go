package main

import (
	"testing"
)

/*
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。

不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

示例 1：

输入：s = ["h","e","l","l","o"]
输出：["o","l","l","e","h"]
示例 2：

输入：s = ["H","a","n","n","a","h"]
输出：["h","a","n","n","a","H"]

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnhbqj/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func Test2_1(t *testing.T) {
	//reverseString([]byte{'h', 'e', 'l', 'l', 'o'})
	//reverseString([]byte{'H', 'a', 'n', 'n', 'a', 'h'})
	reverseString([]string{"h", "e", "l", "l", "o"})
	reverseString([]string{"H", "a", "n", "n", "a", "h"})
}

func reverseString(s []string) {
	l := 0
	r := len(s) - 1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
