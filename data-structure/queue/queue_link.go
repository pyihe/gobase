package queue

import "fmt"

type element struct {
	value interface{}
	next  *element
}

type LinkQueue struct {
	length int
	front  *element // 队头指针
	rear   *element // 队尾指针
}

func NewLinkQueue() *LinkQueue {
	// 链表的头节点，初始状态时头尾指针都指向头节点
	head := &element{}
	return &LinkQueue{
		length: 0,
		front:  head,
		rear:   head,
	}
}

func (queue *LinkQueue) String() (desc string) {
	p := queue.front.next
	for p != nil {
		if desc == "" {
			desc = fmt.Sprintf("%v", p.value)
		} else {
			desc = fmt.Sprintf("%s->%v", desc, p.value)
		}
		p = p.next
	}
	return
}

func (queue *LinkQueue) Len() int {
	return queue.length
}

func (queue *LinkQueue) Clear() {
	head := &element{}
	queue.length = 0
	queue.front = head
	queue.rear = head
}

func (queue *LinkQueue) Head() (v interface{}, ok bool) {
	head := queue.front.next
	if head != nil {
		v = head.value
		ok = true
	}
	return
}

func (queue *LinkQueue) EnQueue(e interface{}) error {
	// 加入的始终是链表中的最后一个节点，所以ele的next为nil
	ele := &element{value: e}
	// 将尾节点的next赋值为ele
	queue.rear.next = ele
	// ele设为尾节点
	queue.rear = ele
	// 长度加1
	queue.length += 1
	return nil
}

func (queue *LinkQueue) DeQueue() (v interface{}, ok bool) {
	// 如果是空队列
	if queue.front == queue.rear {
		return
	}
	// 出队列的节点
	queue.length -= 1
	p := queue.front.next
	v = p.value
	ok = true

	// 节点向前移动一个
	queue.front.next = p.next
	// 如果出队列的是尾节点（即最后一个节点），
	// 此时需要将front赋值给rear，表明队列空了
	if queue.rear == p {
		queue.rear = queue.front
	}
	return
}
