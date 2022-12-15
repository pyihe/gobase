### 接口(API)
```go
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
```

### 实现列表(Implement List)
| 名称(Name)          | 代码(Code)                                                                                            |
|:------------------|:----------------------------------------------------------------------------------------------------|
| 线性栈(Linear Stack) | [stack_linear.go](https://github.com/pyihe/gobase/blob/master/data-structure/stack/stack_linear.go) |
| 链式栈(Link Stack)   | [stack_link.go](https://github.com/pyihe/gobase/blob/master/data-structure/stack/stack_link.go)     |
