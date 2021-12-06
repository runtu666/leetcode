package main

import (
	"fmt"
	"strings"
)

/*
句子 是一个单词列表，列表中的单词之间用单个空格隔开，且不存在前导或尾随空格。每个单词仅由大小写英文字母组成（不含标点符号）。

例如，"Hello World"、"HELLO" 和 "hello world hello world" 都是句子。
给你一个句子 s​​​​​​ 和一个整数 k​​​​​​ ，请你将 s​​ 截断 ​，​​​使截断后的句子仅含 前 k​​​​​​ 个单词。返回 截断 s​​​​​​ 后得到的句子。



示例 1：

输入：s = "Hello how are you Contestant", k = 4
输出："Hello how are you"
解释：
s 中的单词为 ["Hello", "how" "are", "you", "Contestant"]
前 4 个单词为 ["Hello", "how", "are", "you"]
因此，应当返回 "Hello how are you"
示例 2：

输入：s = "What is the solution to this problem", k = 4
输出："What is the solution"
解释：
s 中的单词为 ["What", "is" "the", "solution", "to", "this", "problem"]
前 4 个单词为 ["What", "is", "the", "solution"]
因此，应当返回 "What is the solution"
示例 3：

输入：s = "chopper is not a tanuki", k = 5
输出："chopper is not a tanuki"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/truncate-sentence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	s := "What is the solution to this problem"
	k := 4
	fmt.Println(truncateSentence(s, k))
}

func truncateSentence2(s string, k int) string {
	return strings.Join(strings.Split(s, " ")[:k], " ")
}

func truncateSentence(s string, k int) string {
	var space = ' '
	var words = make([]string, 0)
	var cur = 0
	for i, v := range s {
		if v == space {
			words = append(words, s[cur:i])
			cur = i
		}
	}
	words = append(words, s[cur:])
	var str = ""
	for i, w := range words {
		if i == k {
			break
		}
		str += w
	}
	return str
}
