package tree

import "strings"

type trieNode struct {
	ending   bool               // 是否结尾
	value    rune               // 节点存储的数据
	children map[rune]*trieNode // 孩子节点
}

func newTrieNode(ending bool, value rune, children ...*trieNode) *trieNode {
	node := &trieNode{
		ending:   ending,
		value:    value,
		children: make(map[rune]*trieNode),
	}
	for _, c := range children {
		node.children[c.value] = c
	}
	return node
}

type Trie struct {
	root *trieNode // 根节点
}

func NewTrie() *Trie {
	return &Trie{
		root: &trieNode{
			children: make(map[rune]*trieNode),
		},
	}
}

func (t *Trie) Insert(word string) {
	if word == "" {
		return
	}
	p := t.root
	for i, char := range word {
		ending := i == len(word)-1
		next := p.children[char]
		if next == nil {
			next = newTrieNode(ending, char)
			p.children[char] = next
		}
		p = next
	}
	p.ending = true
}

// Find 完全匹配word字符串
func (t *Trie) Find(word string) (exist bool) {
	p := t.root
	for i, char := range word {
		if p.children == nil {
			break
		}
		next := p.children[char]
		if next == nil {
			break
		}
		if i == len(word)-1 {
			exist = next.ending
			break
		}
		p = next
	}
	return
}

// StartWith 获取所有前缀为prefix的字符串
func (t *Trie) StartWith(prefix string) (collections []string) {
	p := t.root
	for _, char := range prefix {
		// 前缀遍历完之前就没有孩子节点了，直接返回
		if len(p.children) == 0 {
			return
		}
		if p = p.children[char]; p == nil {
			return
		}
	}
	if p.ending {
		collections = []string{prefix}
	}
	cs := collect(p)
	for _, s := range cs {
		builder := &strings.Builder{}
		builder.WriteString(prefix)
		builder.WriteString(s)
		collections = append(collections, builder.String())
	}
	return
}

func collect(root *trieNode) (collection []string) {
	for _, child := range root.children {
		builder := &strings.Builder{}
		builder.WriteRune(child.value)
		if child.ending {
			collection = append(collection, builder.String())
		} else {
			cs := collect(child)
			for _, s := range cs {
				builder.WriteString(s)
				collection = append(collection, builder.String())
			}
		}
	}
	return
}
