package list

import (
	"fmt"
	"reflect"
)

type CircleLink struct {
	head   *singleNode
	tail   *singleNode
	length int
}

func NewCircleLink() *CircleLink {
	return &CircleLink{}
}

func (l *CircleLink) String() (desc string) {
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

func (l *CircleLink) Len() int {
	return l.length
}

func (l *CircleLink) Clear() {
	*l = CircleLink{}
}

func (l *CircleLink) Value(i int) (v interface{}, ok bool) {
	if i < 0 || i >= l.length {
		return
	}
	j, p := 0, l.head
	for p != nil && j < i {
		j += 1
		p = p.next
	}
	if p != nil && j == i {
		v = p.value
		ok = true
	}
	return
}

func (l *CircleLink) Locate(v interface{}, op int) (locations []int) {
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

func (l *CircleLink) Insert(i int, v interface{}, op int) bool {
	if i < 0 || i > l.length {
		return false
	}
	if op <= 0 {
		i -= 1
	}
	j, p := 0, l.head
	for p != nil && j < i {
		j += 1
		p = p.next
	}
	switch {
	case i == -1:
		ele := &singleNode{value: v}
		if l.head == nil {
			ele.next = ele
			l.head = ele
			l.tail = ele
		} else {
			ele.next = l.head
			l.head = ele
			l.tail.next = l.head
		}
		l.length += 1
		return true
	case j == i:
		ele := &singleNode{value: v}
		if p == nil {
			ele.next = ele
			l.head = ele
			l.tail = ele
		} else {
			ele.next = p.next
			p.next = ele
		}
		l.length += 1
		return true
	default:
		return false
	}
}

func (l *CircleLink) RemoveByLocate(i int) (v interface{}, ok bool) {
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
	case i == -1:
		v = l.head.value
		ok = true
		if l.head == l.tail {
			l.head = nil
			l.tail = nil
		} else {
			l.head = l.head.next
			l.tail.next = l.head
		}
		l.length -= 1
		return
	case p != nil && j == i:
		v = p.next.value
		ok = true
		p.next.value = nil
		p.next = p.next.next
		p.next.next = nil
		l.length -= 1
		return
	default:
		return
	}
}

func (l *CircleLink) RemoveByValue(v interface{}, op int) (count int) {
	if op == 0 {
		op = 1
	}
	var (
		p   = l.head
		pre *singleNode
	)
	for p != nil {
		if count == op {
			break
		}
		if !reflect.DeepEqual(p.value, v) {
			pre = p
			p = p.next
			goto step
		}
		switch {
		case pre == nil:
			l.head = p.next
			l.tail.next = l.head
			l.length -= 1
			p.value = nil
			p.next = nil
			p = l.head
			count += 1
		default:
			pre.next = p.next
			p.next = nil
			p.value = nil
			p = pre.next
			l.length -= 1
			count += 1
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
			l.tail = l.head
			l.head = p
		}
		next = p.next
		p.next = pre
		pre = p
		p = next
	}
}

func (l *CircleLink) Range(fn func(i int, value interface{}) bool) {
	j, p := 0, l.head
	for p != nil && j < l.length {
		if fn(j, p.value) {
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
		l.tail = list.tail
		return
	}

	l.tail.next = list.head
	list.tail.next = l.head
	l.tail = list.tail
}
