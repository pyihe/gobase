package stack

type Stack interface {
	// Len 获取栈长度
	Len() int
	// Clear 清空栈
	Clear()
	// Top 获取栈顶元素
	Top() (interface{}, bool)
	// Push 入栈
	Push(interface{})
	// Pop 出栈
	Pop() (interface{}, bool)
}
