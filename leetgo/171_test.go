package main

import (
	"fmt"
	"testing"
)

/**
给你一个字符串columnTitle ，表示 Excel 表格中的列名称。返回该列名称对应的列序号。



例如，

    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28
    ...

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/excel-sheet-column-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test171(t *testing.T) {
	columnTitle := "FXSHRXW"
	fmt.Println(titleToNumber(columnTitle))
}

func titleToNumber(columnTitle string) int {
	var num int
	l := len(columnTitle)
	for i, j := range columnTitle {
		num += int(j-64) * x(26, l-i-1)
	}
	return num
}

func x(num int, x int) int {
	if x == 0 {
		return 1
	}
	var n = 1
	for {
		n *= num
		x--
		if x == 0 {
			break
		}
	}
	return n
}
