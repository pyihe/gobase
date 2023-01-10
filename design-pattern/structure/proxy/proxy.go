package main

import (
	"fmt"
	"sync"
)

/*
	代理模式: 为某对象提供一种代理以控制对该对象的访问。即客户端通过代理间接地访问该对象，从而限制、增强或修改该对象的一些特性
*/

// Proxy 代理接口
type Proxy interface {
	HandleRequest()
}

// RealObject 实际的对象
type RealObject struct {
}

func (ro *RealObject) HandleRequest() {
	fmt.Println("real object handle request")
}

// PxyObject 代理对象
type PxyObject struct {
	once  sync.Once
	proxy Proxy // 代理
}

func (po *PxyObject) newRealObject() {
	// 这里创建一个实际的object
	// 可能是同一个进城内的其他类型的服务；也可能是跨进程或者跨主机的其他服务
	po.once.Do(func() {
		po.proxy = &RealObject{}
	})
}

func (po *PxyObject) HandleRequest() {
	fmt.Println("proxy object handle request")
	if po.proxy == nil {
		po.newRealObject()
	}
	po.proxy.HandleRequest()
}

func main() {
	pxy := &PxyObject{}
	pxy.HandleRequest()
}
