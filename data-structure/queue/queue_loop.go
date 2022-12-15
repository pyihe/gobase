package queue

import (
	"errors"
	"fmt"
)

var (
	ErrFullQueue = errors.New("full queue")
)

type LoopQueue struct {
	front    int // 头指针
	rear     int // 尾指针
	size     int
	elements []interface{} // 队列元素
}

func NewLoopQueue(size int) *LoopQueue {
	// 为了区分队列满和空，空一个不用
	size += 1
	return &LoopQueue{
		front:    0,
		rear:     0,
		size:     size,
		elements: make([]interface{}, size, size),
	}
}

func (queue *LoopQueue) isEmpty() bool {
	return queue.front == queue.rear
}

func (queue *LoopQueue) isFull() bool {
	return (queue.rear+1)%queue.size == queue.front
}

func (queue *LoopQueue) String() (desc string) {
	p := queue.front
	for p != queue.rear {
		if desc == "" {
			desc = fmt.Sprintf("[%v", queue.elements[p])
		} else {
			desc = fmt.Sprintf("%s %v", desc, queue.elements[p])
		}
		p = (p + 1) % queue.size
	}
	desc = fmt.Sprintf("%s]", desc)
	return
}

func (queue *LoopQueue) Len() int {
	return (queue.rear - queue.front + queue.size) % queue.size
}

func (queue *LoopQueue) Clear() {
	queue.front = 0
	queue.rear = 0
	queue.elements = make([]interface{}, queue.size, queue.size)
}

func (queue *LoopQueue) Head() (v interface{}, ok bool) {
	if queue.isEmpty() {
		return
	}
	v = queue.elements[queue.front]
	ok = true
	return
}

func (queue *LoopQueue) EnQueue(e interface{}) error {
	if queue.isFull() {
		return ErrFullQueue
	}
	queue.elements[queue.rear] = e
	queue.rear = (queue.rear + 1) % queue.size
	return nil
}

func (queue *LoopQueue) DeQueue() (v interface{}, ok bool) {
	if queue.isEmpty() {
		return
	}
	v = queue.elements[queue.front]
	queue.front = (queue.front + 1) % queue.size
	return v, true
}
