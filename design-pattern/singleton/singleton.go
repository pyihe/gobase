package main

import (
	"fmt"
	"sync"
)

// 单例模式
// 只能有一个实例
// 1、在内存里只有一个实例，减少了内存的开销，尤其是频繁的创建和销毁实例（比如管理学院首页页面缓存）
// 2、避免对资源的多重占用（比如写文件操作）

type Worker interface {
	Work()
}

type human struct{}

func (h *human) Work() {
	fmt.Println("human working...")
}

var (
	defaultWorker Worker
	one           sync.Once
)

func NewWorker() Worker {
	one.Do(func() {
		defaultWorker = &human{}
	})
	return defaultWorker
}

func main() {
	// 两次调用返回的都是同一个实例，即单例
	NewWorker()
	NewWorker()
}
