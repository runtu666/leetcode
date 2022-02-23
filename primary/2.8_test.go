package main

import (
	"fmt"
	"testing"
)

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xnpvdm/
func Test2_8(t *testing.T) {
	t.Log(countAndSay(3))
}

func countAndSay(n int) string {
	var str = "1"
	if n == 1 {
		return str
	}
	for i := 1; i < n; i++ {
		var tmp = ""
		var curr = string(str[0])
		var cnt = 0
		for _, b := range str {
			s := string(b)
			if s != curr {
				tmp += fmt.Sprintf("%d%s", cnt, curr)
				cnt = 1
				curr = s
				continue
			}
			cnt++
		}
		tmp += fmt.Sprintf("%d%s", cnt, curr)
		str = tmp
	}
	return str
}
