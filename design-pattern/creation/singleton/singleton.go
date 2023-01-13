package main

import (
	"fmt"
	"sync"
)

// 单例模式
// 只能有一个实例
// 1、在内存里只有一个实例，减少了内存的开销，尤其是频繁的创建和销毁实例（比如管理学院首页页面缓存）
// 2、避免对资源的多重占用（比如写文件操作）

/*
	单例模式: 保证一个类仅有一个实例, 并提供一个访问它的全局访问点。目的是减少内存开销
	单例模式分为饿汉式单例类和懒汉式单例类
	1. 饿汉式单例类: 实例在被加载时就将自己初始化
	2. 懒汉式单例类: 实例在第一次被引用时才进行初始化
*/

type Instance struct{}

var (
	// 饿汉式加载
	// instance = &Instance{}
	instance *Instance
	once     sync.Once
)

func GetInstance() *Instance {
	// 懒汉式加载
	once.Do(func() {
		if instance == nil {
			instance = &Instance{}
		}
	})
	return instance
}

func main() {
	// 两次调用返回的都是同一个实例，即单例
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			ins := GetInstance()
			fmt.Printf("%v, %p\n", ins == nil, ins)
			wg.Done()
		}()
	}
	wg.Wait()
}
