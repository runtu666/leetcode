package main

import (
	"strconv"
	"testing"
)

/**
给你一个字符串date ，按 YYYY-MM-DD 格式表示一个 现行公元纪年法 日期。请你计算并返回该日期是当年的第几天。

通常情况下，我们认为 1 月 1 日是每年的第 1 天，1 月 2 日是每年的第 2 天，依此类推。每个月的天数与现行公元纪年法（格里高利历）一致。



示例 1：

输入：date = "2019-01-09"
输出：9
示例 2：

输入：date = "2019-02-10"
输出：41
示例 3：

输入：date = "2003-03-01"
输出：60
示例 4：

输入：date = "2004-03-01"
输出：61

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/day-of-the-year
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test1154(t *testing.T) {
	t.Log(dayOfYear("2004-03-01"))
}

func dayOfYear(date string) int {
	var days = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:])
	var res = day
	for i := 0; i < month-1; i++ {
		res += days[i]
	}
	if year%4 == 0 && month > 2 {
		res++
	}
	return res
}
