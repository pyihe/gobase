package heap

type IntHeap []int

func (ih IntHeap) Len() int {
	return len(ih)
}

func (ih IntHeap) Swap(i, j int) {
	ih[i], ih[j] = ih[j], ih[i]
}

func (ih IntHeap) Less(i, j int) bool {
	// 如果是>，则构建出来的是小顶堆
	return ih[i] < ih[j]
}

func (ih *IntHeap) Pop() (x interface{}) {
	n := len(*ih)
	if n == 0 {
		return
	}
	x = (*ih)[n-1]
	*ih = (*ih)[:n-1]
	return
}

func (ih *IntHeap) Push(x interface{}) {
	ix, ok := x.(int)
	if !ok {
		return
	}
	*ih = append(*ih, ix)
}
