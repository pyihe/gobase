package list

import (
	"fmt"
)

type CircleLink struct {
	head   *singleNode
	rear   *singleNode
	length int
}

func NewCircleLink() *CircleLink {
	return &CircleLink{}
}

func (l *CircleLink) String() (desc string) {
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

func (l *CircleLink) Len() int {
	return l.length
}

func (l *CircleLink) Clear() {
	*l = CircleLink{}
}

func (l *CircleLink) Get(i int) (e *Element) {
	if i < 0 || i >= l.length {
		return
	}
	j, p := 0, l.head
	for p != nil && j < i {
		j += 1
		p = p.next
	}
	if p != nil && j == i {
		e = p.element
	}
	return
}

func (l *CircleLink) Insert(i int, v interface{}, op int) *Element {
	if i < 0 {
		i = 0
	}
	if i >= l.length {
		i = l.length - 1
	}
	if op <= 0 {
		i -= 1
	}
	j, p := 0, l.head
	for p != nil && j < i {
		j += 1
		p = p.next
	}

	e := &Element{v}
	switch {
	case i == -1: // i == -1表示插入位置在头节点前面
		ele := &singleNode{element: e}
		if l.head == nil { // 插入第一个节点, 头尾指针都指向一个节点，next也指向自己
			l.head = ele
			l.rear = ele
			ele.next = ele
		} else {
			ele.next = l.head
			l.head = ele
			l.rear.next = l.head
		}
		l.length += 1
		return e
	case j == i:
		ele := &singleNode{element: e}
		if p == nil { // 插入的是第一个节点
			l.head = ele
			l.rear = ele
			ele.next = ele
		} else {
			ele.next = p.next
			p.next = ele
			if ele.next == l.head {
				l.rear = ele
			}
		}
		l.length += 1
		return e
	default:
		return nil
	}
}

func (l *CircleLink) RemoveByLocate(i int) (e *Element) {
	if i < 0 || i >= l.length {
		return
	}

	i = i - 1
	j, p := 0, l.head
	for p != nil && j < i {
		j += 1
		p = p.next
	}
	switch {
	case i == -1: // 删除的是头节点
		e = l.head.element
		if l.head == l.rear {
			l.head = nil
			l.rear = nil
		} else {
			l.head = l.head.next
			l.rear.next = l.head
		}
		l.length -= 1
		return
	case p != nil && j == i:
		e = p.next.element
		p.next.element = nil
		p.next = p.next.next
		p.next.next = nil
		l.length -= 1
		return
	default:
		return
	}
}

func (l *CircleLink) Remove(v *Element) {
	var (
		p   = l.head
		pre *singleNode
	)
	for p != nil {
		if p.element != v {
			pre = p
			p = p.next
			goto step
		}
		switch {
		case pre == nil:
			l.head = p.next
			l.rear.next = l.head
			l.length -= 1
			p.element = nil
			p.next = nil
			p = l.head
		default:
			pre.next = p.next
			p.next = nil
			p.element = nil
			p = pre.next
			l.length -= 1
		}
	step:
		if p == l.head {
			break
		}
	}
	return
}

func (l *CircleLink) Reverse() {
	var (
		j    = 0
		p    = l.head
		pre  *singleNode
		next *singleNode
	)
	for p != nil && j < l.length {
		if p.next == l.head {
			l.rear = l.head
			l.head = p
		}
		next = p.next
		p.next = pre
		pre = p
		p = next
	}
}

func (l *CircleLink) Range(fn func(i int, e *Element) bool) {
	j, p := 0, l.head
	for p != nil && j < l.length {
		if fn(j, p.element) {
			break
		}
		j += 1
		p = p.next
	}
}

// Merge 合并循环链表
func (l *CircleLink) Merge(list *CircleLink) {
	if list == nil {
		return
	}

	if l.length == 0 {
		l.head = list.head
		l.rear = list.rear
		return
	}

	l.rear.next = list.head
	list.rear.next = l.head
	l.rear = list.rear
}
