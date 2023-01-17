package main

import "testing"

type (
	Trie struct {
		Root *Node
	}
	Node struct {
		IsEnd    bool
		Children map[rune]*Node
	}
)

func Constructor() Trie {
	return Trie{
		Root: newNode(),
	}
}

func newNode() *Node {
	return &Node{
		Children: map[rune]*Node{},
	}
}

func (this *Trie) Insert(word string) {
	chars := []rune(word)
	var curr = this.Root
	for _, char := range chars {
		if _, ok := curr.Children[char]; !ok {
			curr.Children[char] = newNode()
		}
		curr = curr.Children[char]
	}
	curr.IsEnd = true
}

func (this *Trie) Search(word string) bool {
	chars := []rune(word)
	var curr = this.Root
	for _, char := range chars {
		if _, ok := curr.Children[char]; !ok {
			return false
		}
		curr = curr.Children[char]
	}
	return curr.IsEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	chars := []rune(prefix)
	var curr = this.Root
	for _, char := range chars {
		if _, ok := curr.Children[char]; !ok {
			return false
		}
		curr = curr.Children[char]
	}
	return true
}

func Test208(t *testing.T) {

}
