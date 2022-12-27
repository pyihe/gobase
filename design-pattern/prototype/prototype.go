package prototype

import "github.com/pyihe/gobase/pkg"

// 原型模式: 用一个已经创建的实例作为原型，通过复制该原型对象来创建一个和原型相同或相似的新对象

// 优点: 原型实例指定了要创建的对象的种类。用这种方式创建对象非常高效，根本无须知道对象创建的细节

// 缺点: 需要为每一个结构体都配置一个clone方法，

type Object struct {
	PropertyA string
	PropertyB int
	PropertyC *Object
}

func (o *Object) Clone() *Object {
	v := pkg.DeepClone(o)
	return v.(*Object)
}
