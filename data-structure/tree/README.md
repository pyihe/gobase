### 接口(API)

```go
const (
    NoColor Color = iota // 没有颜色
    Red                  // 红黑树：红色
    Black                // 红黑树：黑色
)

// Color 节点颜色
type Color uint8

// Element 树节点存储的元素
type Element interface {
    Value() interface{}  // 元素值
    Compare(Element) int // 两个元素相比，大于返回>0, 小于返回<0, 等于返回=0
}

// Node 树节点
type Node interface {
    String() string // String
    Data() Element  // 节点存储的数据
    Root() Node      // 返回根节点
    LeftChild() Node // 左孩子
    RightChild() Node  // 右孩子
    LeftSibling() Node // 左兄弟
    RightSibling() Node // 右兄弟
    Parent() Node       // 父节点
    Depth() int   // 返回自己所处的深度
    Color() Color // 返回节点颜色
}

// Tree 树
type Tree interface {
    // Root 返回树的根节点
    Root() Node
    // Depth 返回树的深度
    Depth() int
    // Insert 插入新节点
    Insert(Element) Node
    // Remove 移除节点
    Remove(Element) bool
    // Find 查找节点
    Find(Element) Node
    // Update 更新节点
    Update(Element, Element) bool
}
```

### 实现列表(Implement List)

| 名称(Name)      | 代码(Code)                                                                           |
|---------------|------------------------------------------------------------------------------------|
| 树遍历(Traverse) | [tree.go](https://github.com/pyihe/gobase/blob/master/data-structure/tree/tree.go) |
| 二叉搜索树(BST)    | [bst.go](https://github.com/pyihe/gobase/blob/master/data-structure/tree/bst.go)   |
