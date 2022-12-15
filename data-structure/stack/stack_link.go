package stack

import "fmt"

type element struct {
	value interface{}
	next  *element
}

type LinkStack struct {
	top    *element // 头指针作为栈顶元素
	length int      // 栈长度
}

func NewLinkStack() *LinkStack {
	return &LinkStack{}
}

func (s *LinkStack) String() (desc string) {
	p := s.top
	for p != nil {
		if desc == "" {
			desc = fmt.Sprintf("[%v", p.value)
		} else {
			desc = fmt.Sprintf("%s %v", desc, p.value)
		}
		p = p.next
	}
	desc = fmt.Sprintf("%s]", desc)
	return
}

func (s *LinkStack) Len() int {
	return s.length
}

func (s *LinkStack) Clear() {
	s.length = 0
	s.top = nil
}

func (s *LinkStack) Top() (v interface{}, ok bool) {
	if s.top != nil {
		v = s.top.value
		ok = true
	}
	return
}

func (s *LinkStack) Push(e interface{}) {
	s.top = &element{
		value: e,
		next:  s.top,
	}
	s.length += 1
}

func (s *LinkStack) Pop() (e interface{}, ok bool) {
	if s.length == 0 {
		return
	}
	e = s.top.value
	ok = true
	s.top = s.top.next
	s.length -= 1
	return
}
