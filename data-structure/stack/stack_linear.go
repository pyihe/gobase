package stack

import "fmt"

const defaultStackSize = 32

type LinearStack struct {
	initSize int           // 初始容量
	elements []interface{} // 元素数组
}

func NewLinearStack(size int) *LinearStack {
	if size <= 0 {
		size = defaultStackSize
	}
	return &LinearStack{
		initSize: size,
		elements: make([]interface{}, 0, size),
	}
}

func (s *LinearStack) String() (desc string) {
	n := len(s.elements) - 1
	for i := n; i >= 0; i-- {
		if desc == "" {
			desc = fmt.Sprintf("[%v", s.elements[i])
		} else {
			desc = fmt.Sprintf("%s %v", desc, s.elements[i])
		}
	}
	desc = fmt.Sprintf("%s]", desc)
	return
}

func (s *LinearStack) Len() int {
	return len(s.elements)
}

func (s *LinearStack) Clear() {
	s.elements = make([]interface{}, s.initSize, s.initSize)
}

func (s *LinearStack) Top() (e interface{}, ok bool) {
	n := len(s.elements)
	if n == 0 {
		return
	}
	ok = true
	e = s.elements[n-1]
	return
}

func (s *LinearStack) Push(e interface{}) {
	n, c := len(s.elements), cap(s.elements)
	if n+1 > c {
		nElements := make([]interface{}, n, c*2)
		copy(nElements, s.elements)
		s.elements = nElements
	}
	s.elements = s.elements[:n+1]
	s.elements[n] = e
	return
}

func (s *LinearStack) Pop() (e interface{}, ok bool) {
	if len(s.elements) == 0 {
		return
	}

	n, c := len(s.elements), cap(s.elements)
	if n < (c/2) && c > 2*s.initSize {
		nElements := make([]interface{}, n, c/2)
		copy(nElements, s.elements)
		s.elements = nElements
	}
	ok = true
	e = s.elements[n-1]
	s.elements = s.elements[:n-1]
	return
}
