package main

import (
	"strconv"
	"testing"
)

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
