### 接口(API)
```go
type List interface {
    // String 用于打印
    String() string
    
    // Len 返回列表长度
    Len() int
    
    // Clear 清空列表
    Clear()
    
    // Value 返回位置i处的值
    Value(i int) (v interface{}, ok bool)
    
    // Locate 返回v的位置
    // op<0 返回全部v的位置
    // op=0 返回第一个v的位置
    // op>0 返回前op个v的位置
    Locate(v interface{}, op int) []int
    
    // Insert 在位置i处插入元素v
    // op<0 在位置i前面插入
    // op=0 在位置i处插入
    // op>0 在位置i后面插入
    Insert(i int, v interface{}, op int) bool
    
    // RemoveByLocate 删除位置i处的元素, 同时返回被删除元素的值
    RemoveByLocate(i int) (interface{}, bool)
    
    // RemoveByValue 删除与v相等的元素, 同时返回删除的元素个数
    // v 要删除的值
    // op<0 删除全部与v相等的元素
    // op=0 删除一个与v相等的元素
    // op>0 删除op个与v相等的元素
    RemoveByValue(v interface{}, op int) int
    
    // Reverse 列表反转
    Reverse()
    
    // Range 列表遍历， fn返回true终止遍历
    Range(fn func(i int, value interface{}) bool)
}
```

### 实现列表(Implement List)

| 名称(Name)          | 代码(Code)                                                                                         |
|-------------------|--------------------------------------------------------------------------------------------------|
| 线性表(Linear)       | [linear.go](https://github.com/pyihe/gobase/blob/master/data-structure/list/linear.go)           |
| 单链表(Link)         | [link_single.go](https://github.com/pyihe/gobase/blob/master/data-structure/list/link_single.go) |
| 循环链表(Circle Link) | [link_circle.go](https://github.com/pyihe/gobase/blob/master/data-structure/list/link_circle.go) |
| 双链表(Double Link)  | [link_double.go](https://github.com/pyihe/gobase/blob/master/data-structure/list/link_double.go) |
| 静态链表(Static Link) | [link_static.go](https://github.com/pyihe/gobase/blob/master/data-structure/list/link_static.go) |
