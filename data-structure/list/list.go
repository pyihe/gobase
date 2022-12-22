package list

import "fmt"

type List interface {
	// String 用于打印
	String() string

	// Len 返回列表长度
	Len() int

	// Clear 清空列表
	Clear()

	// Get 返回位置i处的值
	Get(i int) *Element

	// Insert 在位置i处插入元素v
	// op<0 在位置i前面插入
	// op=0 在位置i处插入
	// op>0 在位置i后面插入
	Insert(i int, v interface{}, op int) *Element

	// RemoveByLocate 删除位置i处的元素, 同时返回被删除的元素
	RemoveByLocate(i int) *Element

	// Remove 删除元素ele
	Remove(ele *Element)

	// Reverse reverse the list(列表反转)
	Reverse()

	// Range range the list and break when fn returns true(列表遍历， fn返回true终止遍历)
	Range(fn func(i int, element *Element) bool)
}

type Element struct {
	value interface{}
}

func (e *Element) Value() interface{} {
	return e.value
}

func (e *Element) String() (desc string) {
	if e == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%v", e.value)
}
