package main

import "testing"

func Test1450(t *testing.T) {

}

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	var cnt int
	for i, start := range startTime {
		if start <= 4 && endTime[i] >= 4 {
			cnt++
		}
	}
	return cnt
}
