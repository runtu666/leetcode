package main

import (
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/count-the-number-of-consistent-strings/
func Test1684(t *testing.T) {
	allowed := "cad"
	words := []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}
	fmt.Println(countConsistentStrings(allowed, words))
}

func countConsistentStrings(allowed string, words []string) int {
	var m = make(map[int32]struct{})
	for _, v := range allowed {
		m[v] = struct{}{}
	}
	var l = len(words)
	for _, word := range words {
		for _, w := range word {
			if _, ok := m[w]; !ok {
				l--
				break
			}
		}
	}
	return l
}
