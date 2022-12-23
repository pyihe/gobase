package list

import (
	"fmt"
)

/*单向链表构成的列表(链式存储)*/

type singleNode struct {
	element *Element    // 数据域
	next    *singleNode // 指针域
}

type SingleLink struct {
	head   *singleNode // 头指针
	length int         // 列表长度
}

func NewSingleLink() *SingleLink {
	return &SingleLink{}
}

func (l *SingleLink) String() (desc string) {
	p := l.head
	for p != nil {
		if desc == "" {
			desc = fmt.Sprintf("%v", p.element.value)
		} else {
			desc = fmt.Sprintf("%s->%v", desc, p.element.value)
		}
		p = p.next
	}
	return
}

// Len 获取列表长度
func (l *SingleLink) Len() int {
	return l.length
}

// Clear 清空列表
func (l *SingleLink) Clear() {
	l.length = 0
	l.head = nil
}

// Get 获取位置i处的元素
func (l *SingleLink) Get(i int) (e *Element) {
	if i < 0 || i >= l.length {
		return nil
	}

	j, p := 0, l.head

	for p != nil {
		if j == i {
			e = p.element
			break
		}
		j += 1
		p = p.next
	}
	return
}

func (l *SingleLink) Insert(i int, v interface{}, op int) *Element {
	if i < 0 {
		i = 0
	}
	if i >= l.length {
		i = l.length - 1
	}

	// 根据插入位置和插入方向获取新节点的前驱节点所处的位置
	if op <= 0 {
		// 在位置i前面插入，则需要找到的前驱节点为i-1
		// 在位置i后面插入，前驱节点即为i
		i -= 1
	}
	// j 当前遍历的节点序号
	// p 当前遍历的节点
	j, p := 0, l.head
	for p != nil && j < i {
		j += 1
		p = p.next
	}

	e := &Element{v}
	switch {
	case i == -1: // 插入的是头节点
		ele := &singleNode{element: e}
		ele.next = l.head
		l.head = ele
		l.length += 1
		return e
	case j == i: // 找到了前驱节点
		ele := &singleNode{element: e}
		if p == nil { // 插入的是头节点
			ele.next = l.head
			l.head = ele
		} else {
			ele.next = p.next
			p.next = ele
		}
		l.length += 1
		return e
	default:
		return nil
	}
}

// RemoveByLocate 根据位置删除元素
func (l *SingleLink) RemoveByLocate(i int) (e *Element) {
	if i < 0 || i >= l.length {
		return
	}
	// 需要找到前驱节点才能删除
	i = i - 1
	j, p := 0, l.head
	for p != nil && j < i {
		j += 1
		p = p.next
	}
	switch {
	case i == -1: // 删除的是头节点
		e = l.head.element
		l.head = l.head.next
		l.length -= 1
		return
	case p != nil && j == i: // 找到了前驱节点
		e = p.next.element
		p.next.element = nil
		p.next = p.next.next
		l.length -= 1
		return
	default:
		return
	}
}

// Remove 删除与元素v相等的元素
// 默认删除一次
// op小于0表示删除所有与v相等的元素
// op大于0表示删除与v相等的元素op次
// 返回实际删除的次数
func (l *SingleLink) Remove(element *Element) {
	var (
		p   = l.head
		pre *singleNode
	)

	for p != nil {
		if p.element != element {
			pre = p
			p = p.next
			continue
		}
		switch {
		case pre == nil: // 没有前驱节点，证明是头节点
			p = p.next
			l.head.next = nil
			l.head = p
			l.length -= 1

		default:
			pre.next = p.next
			p.next = nil
			p.element = nil
			p = pre.next
			l.length -= 1
		}
		break
	}
	return
}

// Reverse 链表翻转
func (l *SingleLink) Reverse() {
	var (
		p   = l.head    // 当前正在遍历的节点
		pre *singleNode // 当前节点的前驱节点
	)

	for p != nil {
		if p.next == nil {
			l.head = p
		}
		next := p.next // 当前节点的后继节点
		p.next = pre
		pre = p
		p = next
	}
}

// Range 遍历链表
func (l *SingleLink) Range(fn func(i int, e *Element) bool) {
	i, p := 0, l.head
	for p != nil {
		if fn(i, p.element) {
			break
		}
		i += 1
		p = p.next
	}
}
