package main

import "fmt"

func main() {
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
