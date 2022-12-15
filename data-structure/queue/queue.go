package queue

type Queue interface {
	// String 用于打印，按照顺序打印元素
	String() string
	// Len 获取队列长度
	Len() int
	// Clear 清空列表
	Clear()
	// Head 返回队列头元素
	Head() (interface{}, bool)
	// EnQueue 入队列
	EnQueue(interface{}) error
	// DeQueue 出队列
	DeQueue() (interface{}, bool)
}
