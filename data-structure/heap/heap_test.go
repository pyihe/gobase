package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestIntHeap_Push(t *testing.T) {
	h := IntHeap{100, 20, 29, 55, 44}
	heap.Init(&h)
	fmt.Println(heap.Pop(&h))
	heap.Push(&h, 22)
	fmt.Println(heap.Pop(&h))
	h[0] = 15
	heap.Fix(&h, 0)
	fmt.Println(heap.Pop(&h))
}
