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
	for p != nil {
		if j == i {
			break
		}
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

	ele.pre = pre
	ele.next = next
	if pre != nil {
		pre.next = ele
	} else {
		l.head = ele
	}
	if next != nil {
		next.pre = ele
	}
	l.length += 1
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
	pre := p.pre
	next := p.next
	if pre != nil {
		pre.next = next
	} else {
		l.head = next
	}
	if next != nil {
		next.pre = pre
	}
	l.length -= 1
	return p.element
}

func (l *DoubleLink) Remove(v *Element) {
	p := l.head
	for p != nil {
		if p.element != v {
			p = p.next
			continue
		}
		pre := p.pre
		next := p.next
		if pre != nil {
			pre.next = next
		} else {
			l.head = next
		}
		if next != nil {
			next.pre = pre
		}
		l.length -= 1
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
