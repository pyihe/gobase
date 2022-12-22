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

	// output: linear after init: {initSize: 5, elements: [8, 5, 3], len: 3
	fmt.Printf("linear after init: %v, len: %d\n", list, list.Len())

	// 在列表头部插入元素：
	list.Insert(0, 11, -1)
	list.Insert(0, 12, 0)
	list.Insert(0, 13, 1)
	// output: linear insert front: {initSize: 5, elements: [12, 13, 11, 8, 5, 3], len: 6
	fmt.Printf("linear insert front: %v, len: %d\n", list, list.Len())

	// 在列表中间插入元素
	mid := list.Len() / 2
	list.Insert(mid, 21, -1)
	list.Insert(mid, 22, 0)
	list.Insert(mid, 23, 1)
	// output: linear insert mid: {initSize: 5, elements: [12, 13, 11, 22, 23, 21, 8, 5, 3], 9
	fmt.Printf("linear insert mid: %v, %d\n", list, list.Len())

	// 在列表尾部插入元素
	list.Insert(list.Len(), 31, -1)
	list.Insert(list.Len(), 32, 0)
	list.Insert(list.Len(), 33, 1)
	// output: linear insert back: {initSize: 5, elements: [12, 13, 11, 22, 23, 21, 8, 5, 31, 32, 3, 33], len: 12
	fmt.Printf("linear insert back: %v, len: %d\n", list, list.Len())

	// 获取元素
	// output: linear get: front=12, mid=8, back=33
	fmt.Printf("linear get: front=%v, mid=%v, back=%v\n", list.Get(0), list.Get(list.Len()/2), list.Get(list.Len()-1))

	// 根据位置删除元素
	// output: linear removeByLocate: front=12, mid=8, back=33
	fmt.Printf("linear removeByLocate: front=%v, mid=%v, back=%v\n", list.RemoveByLocate(0), list.RemoveByLocate(list.Len()/2), list.RemoveByLocate(list.Len()-1))
	// output: linear after removeByLocate: {initSize: 5, elements: [13, 11, 22, 23, 21, 5, 31, 32, 3], 9
	fmt.Printf("linear after removeByLocate: %v, %d\n", list, list.Len())
	// 删除元素
	list.Remove(list.Get(0))
	list.Remove(list.Get(list.Len() / 2))
	list.Remove(list.Get(list.Len() - 1))
	// output: linear after remove: {initSize: 5, elements: [11, 22, 23, 21, 31, 32], 6
	fmt.Printf("linear after remove: %v, %d\n", list, list.Len())

	list.Reverse()
	// output: linear after reverse: {initSize: 5, elements: [32, 31, 21, 23, 22, 11], 6
	fmt.Printf("linear after reverse: %v, %d\n", list, list.Len())

	// output:
	// linear range: 0=32
	// linear range: 1=31
	// linear range: 2=21
	// linear range: 3=23
	// linear range: 4=22
	// linear range: 5=11
	list.Range(func(i int, element *Element) bool {
		fmt.Printf("linear range: %d=%v\n", i, element.Value())
		return false
	})
	fmt.Println()
	fmt.Println()
	fmt.Println()
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
	fmt.Println()
}

// func TestDoubleLink_Insert(t *testing.T) {
// 	var (
// 		list List = NewDoubleLink()
// 		data      = []int{8, 5, 3}
// 	)
//
// 	for _, v := range data {
// 		list.Insert(list.Len(), v, 1)
// 	}
//
// 	// 初始化后
// 	fmt.Printf("double after init: %v, len: %d\n", list, list.Len())
//
// 	// 在列表头部插入元素：
// 	list.Insert(0, 11, -1)
// 	list.Insert(0, 12, 0)
// 	list.Insert(0, 13, 1)
// 	fmt.Printf("double insert front: %v, len: %d\n", list, list.Len())
//
// 	// 在列表中间插入元素
// 	mid := list.Len() / 2
// 	list.Insert(mid, 21, -1)
// 	list.Insert(mid, 22, 0)
// 	list.Insert(mid, 23, 1)
// 	fmt.Printf("double insert mid: %v, %d\n", list, list.Len())
//
// 	// 在列表尾部插入元素
// 	list.Insert(list.Len(), 31, -1)
// 	list.Insert(list.Len(), 32, 0)
// 	list.Insert(list.Len(), 33, 1)
// 	fmt.Printf("double insert back: %v, len: %d\n", list, list.Len())
//
// 	// 获取元素
// 	fmt.Printf("double get: front=%v, mid=%v, back=%v\n", list.Get(0), list.Get(list.Len()/2), list.Get(list.Len()-1))
//
// 	// 根据位置删除元素
// 	fmt.Printf("double removeByLocate: front=%v, mid=%v, back=%v\n", list.RemoveByLocate(0), list.RemoveByLocate(list.Len()/2), list.RemoveByLocate(list.Len()-1))
// 	fmt.Printf("double after removeByLocate: %v, %d\n", list, list.Len())
// 	// 删除元素
// 	list.Remove(list.Get(0))
// 	list.Remove(list.Get(list.Len() / 2))
// 	list.Remove(list.Get(list.Len() - 1))
// 	fmt.Printf("double after remove: %v, %d\n", list, list.Len())
//
// 	list.Reverse()
// 	fmt.Printf("double after reverse: %v, %d\n", list, list.Len())
//
// 	list.Range(func(i int, element *Element) bool {
// 		fmt.Printf("double range: %d=%v\n", i, element.Value())
// 		return false
// 	})
// 	fmt.Println()
// 	fmt.Println()
// 	fmt.Println()
// }
//
// func TestSingleLink_Insert(t *testing.T) {
// 	var (
// 		list List = NewSingleLink()
// 		data      = []int{8, 5, 3}
// 	)
//
// 	for _, v := range data {
// 		list.Insert(list.Len(), v, 1)
// 	}
//
// 	// 初始化后
// 	fmt.Printf("single after init: %v, len: %d\n", list, list.Len())
//
// 	// 在列表头部插入元素：
// 	list.Insert(0, 11, -1)
// 	list.Insert(0, 12, 0)
// 	list.Insert(0, 13, 1)
// 	fmt.Printf("single insert front: %v, len: %d\n", list, list.Len())
//
// 	// 在列表中间插入元素
// 	mid := list.Len() / 2
// 	list.Insert(mid, 21, -1)
// 	list.Insert(mid, 22, 0)
// 	list.Insert(mid, 23, 1)
// 	fmt.Printf("single insert mid: %v, %d\n", list, list.Len())
//
// 	// 在列表尾部插入元素
// 	list.Insert(list.Len(), 31, -1)
// 	list.Insert(list.Len(), 32, 0)
// 	list.Insert(list.Len(), 33, 1)
// 	fmt.Printf("single insert back: %v, len: %d\n", list, list.Len())
//
// 	// 获取元素
// 	fmt.Printf("single get: front=%v, mid=%v, back=%v\n", list.Get(0), list.Get(list.Len()/2), list.Get(list.Len()-1))
//
// 	// 根据位置删除元素
// 	fmt.Printf("single removeByLocate: front=%v, mid=%v, back=%v\n", list.RemoveByLocate(0), list.RemoveByLocate(list.Len()/2), list.RemoveByLocate(list.Len()-1))
// 	fmt.Printf("single after removeByLocate: %v, %d\n", list, list.Len())
// 	// 删除元素
// 	list.Remove(list.Get(0))
// 	list.Remove(list.Get(list.Len() / 2))
// 	list.Remove(list.Get(list.Len() - 1))
// 	fmt.Printf("single after remove: %v, %d\n", list, list.Len())
//
// 	list.Reverse()
// 	fmt.Printf("single after reverse: %v, %d\n", list, list.Len())
//
// 	list.Range(func(i int, element *Element) bool {
// 		fmt.Printf("single range: %d=%v\n", i, element.Value())
// 		return false
// 	})
// 	fmt.Println()
// 	fmt.Println()
// 	fmt.Println()
// }
//
// func TestStaticLink_Insert(t *testing.T) {
// 	var (
// 		list List = NewStaticLink(5)
// 		data      = []int{8, 5, 3}
// 	)
//
// 	for _, v := range data {
// 		list.Insert(list.Len(), v, 1)
// 	}
//
// 	// 初始化后
// 	fmt.Printf("static after init: %v, len: %d\n", list, list.Len())
//
// 	// 在列表头部插入元素：
// 	list.Insert(0, 11, -1)
// 	list.Insert(0, 12, 0)
// 	list.Insert(0, 13, 1)
// 	fmt.Printf("static insert front: %v, len: %d\n", list, list.Len())
//
// 	// 在列表中间插入元素
// 	mid := list.Len() / 2
// 	list.Insert(mid, 21, -1)
// 	list.Insert(mid, 22, 0)
// 	list.Insert(mid, 23, 1)
// 	fmt.Printf("static insert mid: %v, %d\n", list, list.Len())
//
// 	// 在列表尾部插入元素
// 	list.Insert(list.Len(), 31, -1)
// 	list.Insert(list.Len(), 32, 0)
// 	list.Insert(list.Len(), 33, 1)
// 	fmt.Printf("static insert back: %v, len: %d\n", list, list.Len())
//
// 	// 获取元素
// 	fmt.Printf("static get: front=%v, mid=%v, back=%v\n", list.Get(0), list.Get(list.Len()/2), list.Get(list.Len()-1))
//
// 	// 根据位置删除元素
// 	fmt.Printf("static removeByLocate: front=%v, mid=%v, back=%v\n", list.RemoveByLocate(0), list.RemoveByLocate(list.Len()/2), list.RemoveByLocate(list.Len()-1))
// 	fmt.Printf("static after removeByLocate: %v, %d\n", list, list.Len())
// 	// 删除元素
// 	list.Remove(list.Get(0))
// 	list.Remove(list.Get(list.Len() / 2))
// 	list.Remove(list.Get(list.Len() - 1))
// 	fmt.Printf("static after remove: %v, %d\n", list, list.Len())
//
// 	list.Reverse()
// 	fmt.Printf("static after reverse: %v, %d\n", list, list.Len())
//
// 	list.Range(func(i int, element *Element) bool {
// 		fmt.Printf("static range: %d=%v\n", i, element.Value())
// 		return false
// 	})
// 	fmt.Println()
// 	fmt.Println()
// 	fmt.Println()
// }
