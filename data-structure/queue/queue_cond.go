package queue

import (
	"fmt"
	"sync"
)

type CondQueue struct {
	cond sync.Cond
	data []interface{}
	cap  int
}

func NewCondQueue(size int) *CondQueue {
	return &CondQueue{
		cond: sync.Cond{L: &sync.Mutex{}},
		data: make([]interface{}, 0, size),
		cap:  size,
	}
}

func (cq *CondQueue) String() string {
	desc := ""
	cq.cond.L.Lock()
	for i, v := range cq.data {
		switch i {
		case 0:
			desc = fmt.Sprintf("%v", v)
		default:
			desc = fmt.Sprintf("%s->%v", desc, v)
		}
	}
	cq.cond.L.Unlock()
	return desc
}

func (cq *CondQueue) Len() int {
	cq.cond.L.Lock()
	n := len(cq.data)
	cq.cond.L.Unlock()
	return n
}

func (cq *CondQueue) Clear() {
	cq.cond.L.Lock()
	cq.data = cq.data[0:0]
	cq.cond.L.Unlock()
}

func (cq *CondQueue) Head() (v interface{}, ok bool) {
	cq.cond.L.Lock()
	if len(cq.data) > 0 {
		v = cq.data[0]
		ok = true
	}
	cq.cond.L.Unlock()
	return
}

func (cq *CondQueue) EnQueue(v interface{}) error {
	cq.cond.L.Lock()
	for len(cq.data) == cq.cap {
		cq.cond.Wait()
	}
	cq.data = append(cq.data, v)
	cq.cond.Broadcast()
	cq.cond.L.Unlock()
	return nil
}

func (cq *CondQueue) DeQueue() (v interface{}, ok bool) {
	cq.cond.L.Lock()
	for len(cq.data) == 0 {
		cq.cond.Wait()
	}
	v = cq.data[0]
	ok = true
	cq.data = cq.data[1:]
	cq.cond.Broadcast()
	cq.cond.L.Unlock()
	return
}
