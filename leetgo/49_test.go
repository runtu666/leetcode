package main

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

/**
49. 字母异位词分组
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
说明：

所有输入均为小写字母。
不考虑答案输出的顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/group-anagrams
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func groupAnagrams(strs []string) [][]string {
	var m = make(map[string][]string)
	var resp = make([][]string, 0)
	for _, str := range strs {
		sortStr := getStrSort(str)
		if _, ok := m[sortStr]; !ok {
			m[sortStr] = make([]string, 0)
		}
		m[sortStr] = append(m[sortStr], str)
	}
	for _, item := range m {
		resp = append(resp, item)
	}
	return resp
}

func getStrSort(str string) string {
	var newStrList = make([]string, 0)
	for _, s := range str {
		newStrList = append(newStrList, string(s))
	}
	sort.Slice(newStrList, func(i, j int) bool {
		return newStrList[i] > newStrList[j]
	})
	return strings.Join(newStrList, "")
}

func Test49(t *testing.T) {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
