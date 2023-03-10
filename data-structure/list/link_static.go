package list

import (
	"fmt"
)

// element
type staticNode struct {
	element *Element // 数据域
	next    int      // 指针域
}

/*
	由静态链表构成的列表
	第一个和最后一个元素存放不带数据域的元素，其中：
	第一个元素存放备用链表的头节点
	最后一个元素存放静态链表第一个元素(即有数据的链表的头节点)
*/

type StaticLink struct {
	initSize int           // 初始大小
	length   int           // 链表长度
	nodes    []*staticNode // 链表本身
}

func NewStaticLink(size int) *StaticLink {
	list := &StaticLink{
		initSize: size,
	}

	list.init()

	return list
}

func (l *StaticLink) init() {
	l.length = 0
	l.nodes = make([]*staticNode, l.initSize, l.initSize)
	for i := 0; i < l.initSize; i++ {
		if l.nodes[i] == nil {
			l.nodes[i] = &staticNode{}
		}
		l.nodes[i].element = nil
		switch i {
		case l.initSize - 1: // 数据链表的头指针, 初始状态下指向空
			l.nodes[i].next = 0
		case l.initSize - 2: // 备用链表的尾节点, 指向空
			l.nodes[i].next = 0
		default: // 备用链表, 每个节点指向下一个空闲节点
			l.nodes[i].next = i + 1
		}
	}
}

// 获取下一个空闲节点，如果数组空间不够，则进行扩容
func (l *StaticLink) getNextSpare() (i int) {
	// 静态链表的可用空间由nodes[0]的备用链表索引
	// 所以获取第一个节点即可

	i = l.nodes[0].next

	// 备用头指针可用，返回第一个节点的同时删除第一个节点
	if i != 0 {
		l.nodes[0].next = l.nodes[i].next
		return
	}

	// 如果备用头指针的next不可用，证明数组空间用完了，需要扩容
	c := cap(l.nodes)
	n := len(l.nodes)
	nNode := make([]*staticNode, c*2, c*2)
	copy(nNode, l.nodes)

	// 赋值数据链表的头指针
	nNode[c*2-1] = &staticNode{
		next: l.nodes[n-1].next,
	}

	l.nodes = nNode

	// 重新构建新的备用链表
	for j := n - 1; j < len(l.nodes)-1; j++ {
		if l.nodes[j] == nil {
			l.nodes[j] = &staticNode{}
		}
		l.nodes[j].element = nil
		switch j {
		case len(l.nodes) - 2:
			l.nodes[j].next = 0
		default:
			l.nodes[j].next = j + 1
		}
	}
	i = n - 1
	l.nodes[0].next = n
	return
}

func (l *StaticLink) shrink() {
	c := cap(l.nodes)
	n := len(l.nodes)
	if n >= (c/2) || c < 2*l.initSize {
		return
	}
	// 达到收缩条件, 申请新空间
	// 达到缩容标准且元素数量超过初始容量，则容量减半
	if n >= l.initSize {
		c /= 2
	} else { // 否则回归到初始容量
		c = l.initSize
	}
	fmt.Println("xxx", c)
	nList := make([]*staticNode, c, c)

	// 将数据节点全部转移至新空间
	p := l.nodes[len(l.nodes)-1].next
	for i := 1; i < l.initSize-1; i++ {
		switch {
		case p != 0:
			nod := l.nodes[p]
			p = nod.next
			nList[i] = &staticNode{
				element: nod.element,
				next:    i + 1,
			}
			if i == l.initSize-2 {
				nList[i].next = 0
			}
		default:
			// 记录备用链表的头指针
			if nList[0] == nil {
				nList[0] = &staticNode{next: i}
			}
			nList[i] = &staticNode{}
			if i != l.initSize-2 {
				nList[i] = &staticNode{next: i + 1}
			}
		}
	}

	// 没有空闲空间
	if nList[0] == nil {
		nList[0] = &staticNode{}
	}
	// 数据头节点
	nList[l.initSize-1] = &staticNode{next: 1}
	l.nodes = nList
}

func (l *StaticLink) free(j int) {
	l.nodes[j].element = nil
	l.nodes[j].next = l.nodes[0].next
	l.nodes[0].next = j
}

func (l *StaticLink) String() (desc string) {
	desc = fmt.Sprintf("{initSize: %d, nodes: ", l.initSize)
	i, p := 0, l.nodes[len(l.nodes)-1].next
	for p != 0 && i < l.length {
		nod := l.nodes[p]
		if i == 0 {
			desc = fmt.Sprintf("%s%v", desc, nod.element.value)
		} else {
			desc = fmt.Sprintf("%s->%v", desc, nod.element.value)
		}
		i += 1
		p = nod.next
	}
	desc = fmt.Sprintf("%s}", desc)
	return
}

func (l *StaticLink) Len() int {
	return l.length
}

func (l *StaticLink) Clear() {
	l.init()
}

func (l *StaticLink) Get(i int) (e *Element) {
	if i < 0 || i >= l.length {
		return
	}
	j, p := -1, len(l.nodes)-1
	for p != 0 {
		if j == i {
			e = l.nodes[p].element
			break
		}
		j += 1
		p = l.nodes[p].next
	}
	return
}

func (l *StaticLink) Insert(i int, v interface{}, op int) *Element {
	if i < 0 {
		i = 0
	}
	if i >= l.length {
		i = l.length - 1
	}

	if op <= 0 {
		i -= 1
	}

	// 从备用链表中获取空闲位置，如果空闲链表长度为0，需要扩容
	// 备用链表删除一个节点，该节点作为数据链表的节点插入到结尾
	// 更改新插入的节点的位置
	// 数据链表长度加1

	// 将v保存在位置pos处的节点中
	sparePos := l.getNextSpare()
	n := len(l.nodes)
	ele := &Element{v}
	l.nodes[sparePos].element = ele
	l.nodes[sparePos].next = 0

	switch {
	case i == -1: // 插入的是头节点
		l.nodes[sparePos].next = l.nodes[n-1].next
		l.nodes[n-1].next = sparePos
		l.length += 1

	default:
		k := l.nodes[n-1].next
		for j := 1; j <= i; j++ {
			k = l.nodes[k].next
		}
		// 如果k==0，证明还没有头节点，此时插入头节点
		if k == 0 {
			l.nodes[n-1].next = sparePos
		} else {
			l.nodes[sparePos].next = l.nodes[k].next
			l.nodes[k].next = sparePos
		}
		l.length += 1
	}
	return ele
}

func (l *StaticLink) RemoveByLocate(i int) (e *Element) {
	if i < 0 || i >= l.length {
		return nil
	}

	// 要删除节点i，则需要找到i的前驱节点
	i = i - 1
	target, k := -1, len(l.nodes)-1
	for j := 0; j <= i; j++ {
		k = l.nodes[k].next
	}

	target = l.nodes[k].next
	e = l.nodes[target].element
	l.nodes[k].next = l.nodes[target].next
	l.length -= 1
	l.free(target)
	l.shrink()
	return
}

func (l *StaticLink) Remove(v *Element) {
	var (
		n   = len(l.nodes)
		p   = l.nodes[n-1].next
		pre *staticNode
	)

	for p != 0 {
		nod := l.nodes[p]
		if nod.element != v {
			pre = nod
			p = nod.next
			continue
		}

		switch {
		case pre == nil:
			next := nod.next
			l.free(p)
			p = next
			l.nodes[n-1].next = p
			l.length -= 1
		default:
			next := nod.next
			l.free(p)
			p = next
			pre.next = next
			l.length -= 1
		}
	}

	l.shrink()

	return
}

func (l *StaticLink) Reverse() {
	var (
		n       = len(l.nodes)
		p       = l.nodes[n-1].next
		newTail int
	)

	for p != 0 {
		nod := l.nodes[p]
		if nod.next == 0 {
			l.nodes[n-1].next = p
		}
		next := nod.next
		nod.next = newTail
		newTail = p
		p = next
	}
}

func (l *StaticLink) Range(fn func(i int, value *Element) bool) {
	n := len(l.nodes)
	i, p := 0, l.nodes[n-1].next
	for p != 0 {
		nod := l.nodes[p]
		if fn(i, nod.element) {
			break
		}
		i += 1
		p = l.nodes[p].next
	}
}
