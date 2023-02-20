### 接口(API)
```go
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
```

### 实现列表(Implement List)
| 名称(Name)         | 代码(Code)                                                                                        |
|------------------|-------------------------------------------------------------------------------------------------|
| 循环队列(Loop Queue) | [queue_loop.go](https://github.com/pyihe/gobase/blob/master/data-structure/queue/queue_loop.go) |
| 单链表(Link)        | [queue_link.go](https://github.com/pyihe/gobase/blob/master/data-structure/queue/queue_link.go) |
| 阻塞/通知队列          | [queue_cond.go](https://github.com/pyihe/gobase/blob/master/data-structure/queue/queue_cond.go) |
