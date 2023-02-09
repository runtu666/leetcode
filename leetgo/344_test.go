package main

import "testing"

func Test344(t *testing.T) {
	s := []string{"h", "e", "l", "l", "o"}
	reverseString(s)
	t.Log(s)
}

func reverseString(s []string) {
	var l, r = 0, len(s) - 1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
