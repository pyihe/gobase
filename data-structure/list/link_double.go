package list

import (
	"fmt"
)

type doubleNode struct {
	element *Element
	pre     *doubleNode
	next    *doubleNode
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
			desc = fmt.Sprintf("%v", p.element.value)
		} else {
			desc = fmt.Sprintf("%s->%v", desc, p.element.value)
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

func (l *DoubleLink) Get(i int) (e *Element) {
	if i < 0 || i >= l.length {
		return
	}

	j, p := 0, l.head
	for p != nil {
		if j == i {
			break
		}
		p = p.next
		j += 1
	}
	if p != nil && j == i {
		e = p.element
	}
	return
}

func (l *DoubleLink) Insert(i int, v interface{}, op int) *Element {
	if i < 0 {
		i = 0
	}
	if i >= l.length {
		i = l.length - 1
	}

	j, p := 0, l.head
	for p != nil && j < i {
		p = p.next
		j += 1
	}
	var (
		ele  = &doubleNode{element: &Element{v}}
		pre  *doubleNode
		next *doubleNode
	)

	// 插入第一个节点
	if p == nil {
		l.length += 1
		// 只有一个节点时，前后指针都指向自己
		ele.next = ele
		ele.pre = ele
		l.head = ele
		return ele.element
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

	l.length += 1
	ele.pre = pre
	ele.next = next
	pre.next = ele
	next.pre = ele
	if next == l.head && op <= 0 {
		l.head = ele
	}
	return ele.element
}

func (l *DoubleLink) RemoveByLocate(i int) (e *Element) {
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
	e = p.element
	pre := p.pre
	next := p.next
	l.length -= 1
	pre.next = next
	next.pre = pre
	if next == l.head.next {
		l.head = next
	}
	return p.element
}

func (l *DoubleLink) Remove(v *Element) {
	j, p := 0, l.head
	for p != nil && j < l.length {
		if p.element != v {
			p = p.next
			j += 1
			continue
		}
		l.length -= 1
		pre := p.pre
		next := p.next
		pre.next = next
		next.pre = pre
		// 被删除的是头节点
		if next == l.head.next {
			l.head = next
		}
		break
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

func (l *DoubleLink) Range(fn func(i int, value *Element) bool) {
	j, p := 0, l.head
	for p != nil && j < l.length {
		if fn(j, p.element) {
			break
		}
		j += 1
		p = p.next
	}
}
