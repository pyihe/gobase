package tree

import (
	"fmt"
	"sort"
	"strings"
)

type nodeList []*huffmanNode

func (w nodeList) Len() int {
	return len(w)
}

func (w nodeList) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w nodeList) Less(i, j int) bool {
	switch {
	case w[i].weight < w[j].weight:
		return true
	case w[i].weight > w[j].weight:
		return false
	case w[i].weight == w[j].weight:
		return w[i].data < w[j].data
	}
	return false
}

type huffmanNode struct {
	weight     int          // 数据对应的权重
	data       rune         // 节点存储的数据
	parent     *huffmanNode // 父节点
	leftChild  *huffmanNode // 左孩子
	rightChild *huffmanNode // 右孩子
}

func (node *huffmanNode) String() string {
	if node == nil {
		return "<nil>"
	}
	return fmt.Sprintf("(data: %v, weight: %d)", node.data, node.weight)
}

func (node *huffmanNode) print() (desc string) {
	if node == nil {
		return
	}
	desc = fmt.Sprintf("%s", string(node.data))
	if leftDesc := node.leftChild.print(); leftDesc != "" {
		desc = fmt.Sprintf("%s->%s", desc, leftDesc)
	}
	if rightDesc := node.rightChild.print(); rightDesc != "" {
		desc = fmt.Sprintf("%s->%s", desc, rightDesc)
	}
	return
}

func newHuffmanNode(data rune, weight int) *huffmanNode {
	return &huffmanNode{
		weight: weight,
		data:   data,
	}
}

type HuffmanTree struct {
	root    *huffmanNode    // 树的根结点
	codeSet map[rune]string // 每个原数据对应的编码
}

func NewHuffmanTree(source map[rune]int) *HuffmanTree {
	t := &HuffmanTree{
		codeSet: make(map[rune]string),
	}
	t.build(source)
	return t
}

func (h *HuffmanTree) String() (desc string) {
	if h.root == nil {
		return ""
	}
	return h.root.print()
}

func (h *HuffmanTree) build(source map[rune]int) {
	// 1. 根据原集合生成所有的叶子节点
	leafSet := make(nodeList, 0, len(source))
	for data, weight := range source {
		leafSet = append(leafSet, newHuffmanNode(data, weight))
	}

	regroup := func(list nodeList, extra *huffmanNode) nodeList {
		if extra != nil {
			list = append(list, extra)
		}
		sort.Sort(list)
		return list
	}

	leaf := leafSet
	sort.Sort(leaf)
	// 2. 用叶子节点来构造霍夫曼树
	for leaf.Len() > 1 {
		left, right := leaf[0], leaf[1]
		// 非叶子节点不需要存储数据
		parent := newHuffmanNode(0, left.weight+right.weight)
		parent.leftChild, parent.rightChild = left, right
		left.parent, right.parent = parent, parent

		// 将父节点插入集合，重新排序
		leaf = regroup(leaf[2:], parent)
	}
	h.root = leaf[0]

	// 3. 解析每个叶子节点的编码
	for _, node := range leafSet {
		p := node
		for p.parent != nil {
			if p == p.parent.leftChild {
				// 左孩子路径的编码为0
				h.codeSet[node.data] = fmt.Sprintf("0%s", h.codeSet[node.data])
			} else {
				// 右孩子路径的编码为1
				h.codeSet[node.data] = fmt.Sprintf("1%s", h.codeSet[node.data])
			}
			p = p.parent
		}
	}
}

func (h *HuffmanTree) Encode(char string) string {
	builder := &strings.Builder{}
	for _, c := range char {
		builder.WriteString(h.codeSet[c])
	}
	return builder.String()
}

func (h *HuffmanTree) Decode(code string) string {
	p, builder := h.root, &strings.Builder{}
	for _, c := range code {
		switch c {
		case '1':
			p = p.rightChild
		case '0':
			p = p.leftChild
		default:
			panic("unknown code!")
		}
		if p.leftChild == nil && p.rightChild == nil {
			builder.WriteRune(p.data)
			p = h.root
		}
	}
	return builder.String()
}
