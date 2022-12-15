package list

import (
	"fmt"
	"reflect"
)

const (
	defaultInitSize = 32
)

// 切片构成的线性表(顺序存储)

type LinearList struct {
	initSize int           // 列表初始容量
	elements []interface{} // 列表元素
}

func NewLinearList(size int) *LinearList {
	if size <= 0 {
		size = defaultInitSize
	}
	return &LinearList{
		elements: make([]interface{}, 0, size),
		initSize: size,
	}
}

func (l *LinearList) String() (desc string) {
	return fmt.Sprintf("{initSize: %d, elements: %v", l.initSize, l.elements)
}

// Len 获取列表长度
func (l *LinearList) Len() int {
	return len(l.elements)
}

// Clear 清空列表
func (l *LinearList) Clear() {
	l.elements = make([]interface{}, 0, l.initSize)
}

// Value 获取位置i处的元素
func (l *LinearList) Value(i int) (v interface{}, ok bool) {
	n := len(l.elements)
	if i < 0 || i >= n {
		return
	}
	v = l.elements[i]
	ok = true
	return
}

// Locate 找出与v相等的元素
// op 默认找出一个相等的元素，为负数表示找出所有相等的元素，否则找出op对应个数的元素
func (l *LinearList) Locate(v interface{}, op int) (locations []int) {
	if op == 0 {
		op = 1
	}
	// 预分配
	locations = make([]int, 0, len(l.elements))

	for i, e := range l.elements {
		if !reflect.DeepEqual(e, v) {
			continue
		}
		locations = append(locations, i)
		if len(locations) == op {
			break
		}
	}
	return
}

// Insert 插入元素
// op 为0表示在位置i插入元素
// op为正数表示在i后面插入元素
// op为负数表示在i前面插入元素
func (l *LinearList) Insert(i int, v interface{}, op int) bool {
	n := len(l.elements)
	if i < 0 || i > n {
		return false
	}

	// 需要扩容
	l.expansion()
	l.elements = l.elements[0 : n+1]

	switch {
	case op <= 0: // 在i前面插入
		copy(l.elements[i+1:], l.elements[i:])
		l.elements[i] = v
	case op > 0: // 在i后面插入
		if n == 0 || i == n {
			l.elements[n] = v
		} else {
			copy(l.elements[i+2:], l.elements[i+1:])
			l.elements[i+1] = v
		}
	}

	return true
}

// RemoveByLocate 删除位置i处的元素
func (l *LinearList) RemoveByLocate(i int) (interface{}, bool) {
	n := len(l.elements)
	if i < 0 || i >= n {
		return nil, false
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
	return v, true
}

// RemoveByValue 删除与元素v相等的元素
// 默认删除一次
// op小于0表示删除所有与v相等的元素
// op大于0表示删除与v相等的元素op次
// 返回实际删除的次数
func (l *LinearList) RemoveByValue(v interface{}, op int) (count int) {
	if op == 0 {
		op = 1
	}
	n := len(l.elements)
	l.shrink()

	for i := 0; i < n; i++ {
		element := l.elements[i]
		if !reflect.DeepEqual(element, v) {
			continue
		}
		copy(l.elements[i:], l.elements[i+1:])
		l.elements[n-1] = nil
		l.elements = l.elements[:n-1]
		count += 1
		if op > 0 && count == op {
			break
		}
		// i被删除后，i+1以及之后的元素位置前移1，此时需要重新从i开始遍历，同时总的遍历次数减1
		i, n = i-1, n-1
	}
	return
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
func (l *LinearList) Range(fn func(i int, value interface{}) bool) {
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
		nList := make([]interface{}, n, c/2)
		copy(nList, l.elements)
		l.elements = nList
	}
}

func (l *LinearList) expansion() {
	n := len(l.elements)
	c := cap(l.elements)
	if n > 0 && n+1 > c {
		nList := make([]interface{}, n, c*2)
		copy(nList, l.elements)
		l.elements = nList
	}
}
