package list

import (
	"fmt"
)

const (
	defaultInitSize = 32
)

// 切片构成的线性表(顺序存储)

type LinearList struct {
	initSize int        // 列表初始容量
	elements []*Element // 列表元素
}

func NewLinearList(size int) *LinearList {
	if size <= 0 {
		size = defaultInitSize
	}
	return &LinearList{
		elements: make([]*Element, 0, size),
		initSize: size,
	}
}

func (l *LinearList) String() (desc string) {
	desc = fmt.Sprintf("{initSize: %d, elements: [", l.initSize)
	for i, e := range l.elements {
		switch {
		case i == 0:
			desc = fmt.Sprintf("%s%v", desc, e.value)
		default:
			desc = fmt.Sprintf("%s, %v", desc, e.value)
		}
	}
	return fmt.Sprintf("%s]", desc)
}

// Len 获取列表长度
func (l *LinearList) Len() int {
	return len(l.elements)
}

// Clear 清空列表
func (l *LinearList) Clear() {
	l.elements = make([]*Element, 0, l.initSize)
}

// Get 获取位置i处的元素
func (l *LinearList) Get(i int) (e *Element) {
	n := len(l.elements)
	if i < 0 || i >= n {
		return
	}
	e = l.elements[i]
	return
}

// Insert 插入元素
// op 为0表示在位置i插入元素
// op为正数表示在i后面插入元素
// op为负数表示在i前面插入元素
func (l *LinearList) Insert(i int, v interface{}, op int) *Element {
	n := len(l.elements)
	if i < 0 {
		i = 0
	}
	if i >= n {
		i = n - 1
	}

	ele := &Element{v}

	// 需要扩容
	l.expansion()
	l.elements = l.elements[0 : n+1]

	switch {
	case op <= 0: // 在i前面插入
		copy(l.elements[i+1:], l.elements[i:])
		l.elements[i] = ele
	case op > 0: // 在i后面插入
		if n == 0 || i == n {
			l.elements[n] = ele
		} else {
			copy(l.elements[i+2:], l.elements[i+1:])
			l.elements[i+1] = ele
		}
	}

	return ele
}

// RemoveByLocate 删除位置i处的元素
func (l *LinearList) RemoveByLocate(i int) *Element {
	n := len(l.elements)
	if i < 0 || i >= n {
		return nil
	}

	// 是否需要缩容
	l.shrink()

	// 删除位置i的元素
	// 获取i对应的值，用于返回
	v := l.elements[i]
	// 将i之后的所有元素向前移动一个位置
	copy(l.elements[i:], l.elements[i+1:])
	// 将最后一个元素置为nil
	l.elements[n-1] = nil
	// 去掉最后一个元素
	l.elements = l.elements[:n-1]
	return v
}

// Remove 删除与元素v相等的元素
// 默认删除一次
// op小于0表示删除所有与v相等的元素
// op大于0表示删除与v相等的元素op次
// 返回实际删除的次数
func (l *LinearList) Remove(e *Element) {
	n := len(l.elements)
	l.shrink()

	for i := 0; i < n; i++ {
		element := l.elements[i]
		if element != e {
			continue
		}
		copy(l.elements[i:], l.elements[i+1:])
		l.elements[n-1] = nil
		l.elements = l.elements[:n-1]
		break
	}
}

// Reverse 列表反转
func (l *LinearList) Reverse() {
	n := len(l.elements)
	for i := n/2 - 1; i >= 0; i-- {
		opp := n - i - 1
		l.elements[i], l.elements[opp] = l.elements[opp], l.elements[i]
	}
}

// Range 遍历列表
func (l *LinearList) Range(fn func(i int, e *Element) bool) {
	for i, v := range l.elements {
		if fn(i, v) {
			break
		}
	}
}

func (l *LinearList) shrink() {
	c := cap(l.elements)
	n := len(l.elements)
	if n < (c/2) && c >= 2*l.initSize {
		nList := make([]*Element, n, c/2)
		copy(nList, l.elements)
		l.elements = nList
	}
}

func (l *LinearList) expansion() {
	n := len(l.elements)
	c := cap(l.elements)
	if n > 0 && n+1 > c {
		nList := make([]*Element, n, c*2)
		copy(nList, l.elements)
		l.elements = nList
	}
}
