package list

import (
	"fmt"
	"reflect"
)

type doubleNode struct {
	value interface{}
	pre   *doubleNode
	next  *doubleNode
}

type DoubleLink struct {
	head   *doubleNode
	length int
}

func NewDoubleLink() *DoubleLink {
	return &DoubleLink{}
}

func (l *DoubleLink) String() (desc string) {
	i, p := 0, l.head
	for p != nil && i < l.length {
		if desc == "" {
			desc = fmt.Sprintf("%v", p.value)
		} else {
			desc = fmt.Sprintf("%s->%v", desc, p.value)
		}
		i += 1
		p = p.next
	}
	return
}

func (l *DoubleLink) Len() int {
	return l.length
}

func (l *DoubleLink) Clear() {
	*l = DoubleLink{}
}

func (l *DoubleLink) Value(i int) (v interface{}, ok bool) {
	if i < 0 || i >= l.length {
		return
	}

	var (
		j int
		p *doubleNode
	)

	switch {
	case i > l.length/2: // 从尾节点开始遍历
		j, p = l.length, l.head
		for p != nil && j >= i {
			j -= 1
			p = p.pre
		}
	case i <= l.length/2: // 从头节点开始遍历
		j, p = 0, l.head
		for p != nil && j < i {
			j += 1
			p = p.next
		}
	}
	if p != nil && j == i {
		v = p.value
		ok = true
	}
	return
}

func (l *DoubleLink) Locate(v interface{}, op int) (locations []int) {
	if op == 0 {
		op = 1
	}
	locations = make([]int, 0, l.length)
	j, p := 0, l.head
	for p != nil && j < l.length {
		if !reflect.DeepEqual(p.value, v) {
			goto step
		}
		locations = append(locations, j)
		if len(locations) == op {
			break
		}
	step:
		j += 1
		p = p.next
	}
	return
}

func (l *DoubleLink) Insert(i int, v interface{}, op int) bool {
	if i < 0 || i > l.length {
		return false
	}
	var (
		j int
		p *doubleNode
	)
	switch {
	case i > l.length/2: // 向前遍历
		j, p = l.length, l.head
		for p != nil && j > i {
			j -= 1
			p = p.pre
		}
	case i <= l.length/2: // 向后遍历
		j, p = 0, l.head
		for p != nil && j < i {
			j += 1
			p = p.next
		}
	}

	var (
		ele  = &doubleNode{value: v}
		pre  *doubleNode
		next *doubleNode
	)

	// 插入第一个节点
	if p == nil {
		ele.next = ele
		ele.pre = ele
		l.length += 1
		l.head = ele
		return true
	}

	switch {
	case op == 0:
		pre = p.pre
		next = p
	case op > 0:
		pre = p
		next = p.next
	default:
		pre = p.pre
		next = p
	}

	ele.pre = pre
	ele.next = next
	pre.next = ele
	next.pre = ele
	l.length += 1
	if i == 0 && op <= 0 {
		l.head = ele
	}
	return true
}

func (l *DoubleLink) RemoveByLocate(i int) (v interface{}, ok bool) {
	if i < 0 || i >= l.length {
		return
	}
	var (
		j int
		p *doubleNode
	)
	switch {
	case i > l.length/2:
		j, p = l.length, l.head
		for p != nil && j > i {
			j -= 1
			p = p.pre
		}
	case i <= l.length/2:
		j, p = 0, l.head
		for p != nil && j < i {
			j += 1
			p = p.next
		}
	}
	pre := p.pre
	next := p.next
	pre.next = next
	next.pre = pre
	l.length -= 1
	return p.value, true
}

func (l *DoubleLink) RemoveByValue(v interface{}, op int) (count int) {
	if op == 0 {
		op = 1
	}
	p := l.head
	for p != nil {
		if count == op {
			break
		}
		var (
			pre  *doubleNode
			next *doubleNode
		)
		if !reflect.DeepEqual(p.value, v) {
			p = p.next
			goto step
		}
		// 如果被删除的是头节点
		if p == l.head {
			l.head = p.next
		}
		p = p.next
		pre = p.pre
		next = p.next
		pre.next = next
		next.pre = pre
		l.length -= 1
		count += 1
	step:
		if p == l.head {
			break
		}
	}
	return
}

func (l *DoubleLink) Reverse() {
	var (
		p = l.head
		// tail = l.head.pre
		pre  *doubleNode
		next *doubleNode
	)
	for p != nil {
		next = p.next
		p.next = pre
		p.pre = next
		pre = p
		if next == nil {
			l.head = p.next
		}
		p = next
	}
}

func (l *DoubleLink) Range(fn func(i int, value interface{}) bool) {
	j, p := 0, l.head
	for p != nil && j < l.length {
		if fn(j, p.value) {
			break
		}
		j += 1
		p = p.next
	}
}
