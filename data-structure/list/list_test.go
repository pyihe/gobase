package list

import (
	"fmt"
	"testing"
)

func TestLinearList_Insert(t *testing.T) {
	var (
		list List = NewLinearList(5)
		data      = []int{8, 5, 3}
	)

	for _, v := range data {
		list.Insert(list.Len(), v, 1)
	}

	fmt.Printf("linear after init: %v, len: %d\n", list, list.Len())

	// 在列表头部插入元素：
	list.Insert(0, 11, -1)
	list.Insert(0, 12, 0)
	list.Insert(0, 13, 1)
	fmt.Printf("linear insert front: %v, len: %d\n", list, list.Len())

	// 在列表中间插入元素
	mid := list.Len() / 2
	list.Insert(mid, 21, -1)
	list.Insert(mid, 22, 0)
	list.Insert(mid, 23, 1)
	fmt.Printf("linear insert mid: %v, %d\n", list, list.Len())

	// 在列表尾部插入元素
	list.Insert(list.Len(), 31, -1)
	list.Insert(list.Len(), 32, 0)
	list.Insert(list.Len(), 33, 1)
	fmt.Printf("linear insert back: %v, len: %d\n", list, list.Len())

	// 获取元素
	fmt.Printf("linear get: front=%v, mid=%v, back=%v\n", list.Get(0), list.Get(list.Len()/2), list.Get(list.Len()-1))

	// 根据位置删除元素
	fmt.Printf("linear removeByLocate: front=%v, mid=%v, back=%v\n", list.RemoveByLocate(0), list.RemoveByLocate(list.Len()/2), list.RemoveByLocate(list.Len()-1))

	fmt.Printf("linear after removeByLocate: %v, %d\n", list, list.Len())
	// 删除元素
	list.Remove(list.Get(0))
	list.Remove(list.Get(list.Len() / 2))
	list.Remove(list.Get(list.Len() - 1))
	fmt.Printf("linear after remove: %v, %d\n", list, list.Len())

	list.Reverse()
	fmt.Printf("linear after reverse: %v, %d\n", list, list.Len())

	list.Range(func(i int, element *Element) bool {
		fmt.Printf("linear range: %d=%v\n", i, element.Value())
		return false
	})
	fmt.Println()
	fmt.Println()

	// output:
	// linear after init: {initSize: 5, elements: [8, 5, 3], len: 3
	// linear insert front: {initSize: 5, elements: [12, 13, 11, 8, 5, 3], len: 6
	// linear insert mid: {initSize: 5, elements: [12, 13, 11, 22, 23, 21, 8, 5, 3], 9
	// linear insert back: {initSize: 5, elements: [12, 13, 11, 22, 23, 21, 8, 5, 31, 32, 3, 33], len: 12
	// linear get: front=12, mid=8, back=33
	// linear removeByLocate: front=12, mid=8, back=33
	// linear after removeByLocate: {initSize: 5, elements: [13, 11, 22, 23, 21, 5, 31, 32, 3], 9
	// linear after remove: {initSize: 5, elements: [11, 22, 23, 21, 31, 32], 6
	// linear after reverse: {initSize: 5, elements: [32, 31, 21, 23, 22, 11], 6
	// linear range: 0=32
	// linear range: 1=31
	// linear range: 2=21
	// linear range: 3=23
	// linear range: 4=22
	// linear range: 5=11
}

func TestCircleLink_Insert(t *testing.T) {
	var (
		list List = NewCircleLink()
		data      = []int{8, 5, 3}
	)

	for _, v := range data {
		list.Insert(list.Len(), v, 1)
	}

	// 初始化后
	// output:
	fmt.Printf("circle after init: %v, len: %d\n", list, list.Len())

	// 在列表头部插入元素：
	list.Insert(0, 11, -1)
	list.Insert(0, 12, 0)
	list.Insert(0, 13, 1)
	fmt.Printf("circle insert front: %v, len: %d\n", list, list.Len())

	// 在列表中间插入元素
	mid := list.Len() / 2
	list.Insert(mid, 21, -1)
	list.Insert(mid, 22, 0)
	list.Insert(mid, 23, 1)
	fmt.Printf("circle insert mid: %v, %d\n", list, list.Len())

	// 在列表尾部插入元素
	list.Insert(list.Len(), 31, -1)
	list.Insert(list.Len(), 32, 0)
	list.Insert(list.Len(), 33, 1)
	fmt.Printf("circle insert back: %v, len: %d\n", list, list.Len())

	// 获取元素
	fmt.Printf("circle get: front=%v, mid=%v, back=%v\n", list.Get(0), list.Get(list.Len()/2), list.Get(list.Len()-1))

	// 根据位置删除元素
	fmt.Printf("circle removeByLocate: front=%v, mid=%v, back=%v\n", list.RemoveByLocate(0), list.RemoveByLocate(list.Len()/2), list.RemoveByLocate(list.Len()-1))
	fmt.Printf("circle after removeByLocate: %v, %d\n", list, list.Len())
	// 删除元素
	list.Remove(list.Get(0))
	list.Remove(list.Get(list.Len() / 2))
	list.Remove(list.Get(list.Len() - 1))
	fmt.Printf("circle after remove: %v, %d\n", list, list.Len())

	list.Reverse()
	fmt.Printf("circle after reverse: %v, %d\n", list, list.Len())

	list.Range(func(i int, element *Element) bool {
		fmt.Printf("circle range: %d=%v\n", i, element.Value())
		return false
	})
	fmt.Println()
	fmt.Println()

	// output:
	// circle after init: 8->5->3, len: 3
	// circle insert front: 12->13->11->8->5->3, len: 6
	// circle insert mid: 12->13->11->22->23->21->8->5->3, 9
	// circle insert back: 12->13->11->22->23->21->8->5->31->32->3->33, len: 12
	// circle get: front=12, mid=8, back=33
	// circle removeByLocate: front=12, mid=8, back=33
	// circle after removeByLocate: 13->11->22->23->21->5->31->32->3, 9
	// circle after remove: 11->22->23->21->31->32, 6
	// circle after reverse: 32->31->21->23->22->11, 6
	// circle range: 0=32
	// circle range: 1=31
	// circle range: 2=21
	// circle range: 3=23
	// circle range: 4=22
	// circle range: 5=11
}

func TestDoubleLink_Insert(t *testing.T) {
	var (
		list List = NewDoubleLink()
		data      = []int{8, 5, 3}
	)

	for _, v := range data {
		list.Insert(list.Len(), v, 1)
	}

	// 初始化后
	fmt.Printf("double after init: %v, len: %d\n", list, list.Len())

	// 在列表头部插入元素：
	list.Insert(0, 11, -1)
	list.Insert(0, 12, 0)
	list.Insert(0, 13, 1)
	fmt.Printf("double insert front: %v, len: %d\n", list, list.Len())

	// 在列表中间插入元素
	mid := list.Len() / 2
	list.Insert(mid, 21, -1)
	list.Insert(mid, 22, 0)
	list.Insert(mid, 23, 1)
	fmt.Printf("double insert mid: %v, %d\n", list, list.Len())

	// 在列表尾部插入元素
	list.Insert(list.Len(), 31, -1)
	list.Insert(list.Len(), 32, 0)
	list.Insert(list.Len(), 33, 1)
	fmt.Printf("double insert back: %v, len: %d\n", list, list.Len())

	// 获取元素
	fmt.Printf("double get: front=%v, mid=%v, back=%v\n", list.Get(0), list.Get(list.Len()/2), list.Get(list.Len()-1))

	// 根据位置删除元素
	fmt.Printf("double removeByLocate: front=%v, mid=%v, back=%v\n", list.RemoveByLocate(0), list.RemoveByLocate(list.Len()/2), list.RemoveByLocate(list.Len()-1))
	fmt.Printf("double after removeByLocate: %v, %d\n", list, list.Len())
	// 删除元素
	list.Remove(list.Get(0))
	list.Remove(list.Get(list.Len() / 2))
	list.Remove(list.Get(list.Len() - 1))
	fmt.Printf("double after remove: %v, %d\n", list, list.Len())

	list.Reverse()
	fmt.Printf("double after reverse: %v, %d\n", list, list.Len())

	list.Range(func(i int, element *Element) bool {
		fmt.Printf("double range: %d=%v\n", i, element.Value())
		return false
	})
	fmt.Println()
	fmt.Println()

	// output:
	// double after init: 8->5->3, len: 3
	// double insert front: 12->13->11->8->5->3, len: 6
	// double insert mid: 12->13->11->22->23->21->8->5->3, 9
	// double insert back: 12->13->11->22->23->21->8->5->31->32->3->33, len: 12
	// double get: front=12, mid=8, back=33
	// double removeByLocate: front=12, mid=8, back=33
	// double after removeByLocate: 13->11->22->23->21->5->31->32->3, 9
	// double after remove: 11->22->23->21->31->32, 6
	// double after reverse: 32->31->21->23->22->11, 6
	// double range: 0=32
	// double range: 1=31
	// double range: 2=21
	// double range: 3=23
	// double range: 4=22
	// double range: 5=11
}

func TestSingleLink_Insert(t *testing.T) {
	var (
		list List = NewSingleLink()
		data      = []int{8, 5, 3}
	)

	for _, v := range data {
		list.Insert(list.Len(), v, 1)
	}

	// 初始化后
	fmt.Printf("single after init: %v, len: %d\n", list, list.Len())

	// 在列表头部插入元素：
	list.Insert(0, 11, -1)
	list.Insert(0, 12, 0)
	list.Insert(0, 13, 1)
	fmt.Printf("single insert front: %v, len: %d\n", list, list.Len())

	// 在列表中间插入元素
	mid := list.Len() / 2
	list.Insert(mid, 21, -1)
	list.Insert(mid, 22, 0)
	list.Insert(mid, 23, 1)
	fmt.Printf("single insert mid: %v, %d\n", list, list.Len())

	// 在列表尾部插入元素
	list.Insert(list.Len(), 31, -1)
	list.Insert(list.Len(), 32, 0)
	list.Insert(list.Len(), 33, 1)
	fmt.Printf("single insert back: %v, len: %d\n", list, list.Len())

	// 获取元素
	fmt.Printf("single get: front=%v, mid=%v, back=%v\n", list.Get(0), list.Get(list.Len()/2), list.Get(list.Len()-1))

	// 根据位置删除元素
	fmt.Printf("single removeByLocate: front=%v, mid=%v, back=%v\n", list.RemoveByLocate(0), list.RemoveByLocate(list.Len()/2), list.RemoveByLocate(list.Len()-1))
	fmt.Printf("single after removeByLocate: %v, %d\n", list, list.Len())
	// 删除元素
	list.Remove(list.Get(0))
	list.Remove(list.Get(list.Len() / 2))
	list.Remove(list.Get(list.Len() - 1))
	fmt.Printf("single after remove: %v, %d\n", list, list.Len())

	list.Reverse()
	fmt.Printf("single after reverse: %v, %d\n", list, list.Len())

	list.Range(func(i int, element *Element) bool {
		fmt.Printf("single range: %d=%v\n", i, element.Value())
		return false
	})
	fmt.Println()
	fmt.Println()

	// output:
	// single after init: 8->5->3, len: 3
	// single insert front: 12->13->11->8->5->3, len: 6
	// single insert mid: 12->13->11->22->23->21->8->5->3, 9
	// single insert back: 12->13->11->22->23->21->8->5->31->32->3->33, len: 12
	// single get: front=12, mid=8, back=33
	// single removeByLocate: front=12, mid=8, back=33
	// single after removeByLocate: 13->11->22->23->21->5->31->32->3, 9
	// single after remove: 11->22->23->21->31->32, 6
	// single after reverse: 32->31->21->23->22->11, 6
	// single range: 0=32
	// single range: 1=31
	// single range: 2=21
	// single range: 3=23
	// single range: 4=22
	// single range: 5=11
}

func TestStaticLink_Insert(t *testing.T) {
	var (
		list List = NewStaticLink(5)
		data      = []int{8, 5, 3}
	)

	for _, v := range data {
		list.Insert(list.Len(), v, 1)
	}

	// 初始化后
	fmt.Printf("static after init: %v, len: %d\n", list, list.Len())

	// 在列表头部插入元素：
	list.Insert(0, 11, -1)
	list.Insert(0, 12, 0)
	list.Insert(0, 13, 1)
	fmt.Printf("static insert front: %v, len: %d\n", list, list.Len())

	// 在列表中间插入元素
	mid := list.Len() / 2
	list.Insert(mid, 21, -1)
	list.Insert(mid, 22, 0)
	list.Insert(mid, 23, 1)
	fmt.Printf("static insert mid: %v, %d\n", list, list.Len())

	// 在列表尾部插入元素
	list.Insert(list.Len(), 31, -1)
	list.Insert(list.Len(), 32, 0)
	list.Insert(list.Len(), 33, 1)
	fmt.Printf("static insert back: %v, len: %d\n", list, list.Len())

	// 获取元素
	fmt.Printf("static get: front=%v, mid=%v, back=%v\n", list.Get(0), list.Get(list.Len()/2), list.Get(list.Len()-1))

	// 根据位置删除元素
	fmt.Printf("static removeByLocate: front=%v, mid=%v, back=%v\n", list.RemoveByLocate(0), list.RemoveByLocate(list.Len()/2), list.RemoveByLocate(list.Len()-1))
	fmt.Printf("static after removeByLocate: %v, %d\n", list, list.Len())
	// 删除元素
	list.Remove(list.Get(0))
	list.Remove(list.Get(list.Len() / 2))
	list.Remove(list.Get(list.Len() - 1))
	list.Remove(list.Get(0))
	list.Remove(list.Get(0))
	fmt.Printf("static after remove: %v, %d\n", list, list.Len())

	list.Reverse()
	fmt.Printf("static after reverse: %v, %d\n", list, list.Len())

	list.Range(func(i int, element *Element) bool {
		fmt.Printf("static range: %d=%v\n", i, element.Value())
		return false
	})
	fmt.Println()
	fmt.Println()

	// output:
	// static after init: {initSize: 5, nodes: 8->5->3}, len: 3
	// static insert front: {initSize: 5, nodes: 12->13->11->8->5->3}, len: 6
	// static insert mid: {initSize: 5, nodes: 12->13->11->22->23->21->8->5->3}, 9
	// static insert back: {initSize: 5, nodes: 12->13->11->22->23->21->8->5->31->32->3->33}, len: 12
	// static get: front=12, mid=8, back=33
	// static removeByLocate: front=12, mid=8, back=33
	// static after removeByLocate: {initSize: 5, nodes: 13->11->22->23->21->5->31->32->3}, 9
	// static after remove: {initSize: 5, nodes: 23->21->31->32}, 4
	// static after reverse: {initSize: 5, nodes: 32->31->21->23}, 4
	// static range: 0=32
	// static range: 1=31
	// static range: 2=21
	// static range: 3=23
}
