package main

import "fmt"

// 建造者模式
// 在软件开发过程中有时需要创建一个复杂的对象，这个复杂对象通常由多个子部件按一定的步骤组合而成。
// 例如，计算机是由 CPU、主板、内存、硬盘、显卡、机箱、显示器、键盘、鼠标等部件组装而成的，
// 采购员不可能自己去组装计算机，而是将计算机的配置要求告诉计算机销售公司，计算机销售公司安排技术人员
// 去组装计算机，然后再交给要买计算机的采购员。

// 建造者模式的主要优点如下：
// 1. 封装性好，构建和表示分离。
// 2. 扩展性好，各个具体的建造者相互独立，有利于系统的解耦。
// 3. 客户端不必知道产品内部组成的细节，建造者可以对创建过程逐步细化，而不对其它模块产生任何影响，便于控制细节风险。

// 缺点如下：
// 1. 产品的组成部分必须相同，这限制了其使用范围。
// 2. 如果产品的内部变化复杂，如果产品内部发生变化，则建造者也要同步修改，后期维护成本较大。

// Builder 建造者接口，每个建造者构建Product时需要实现的接口
type Builder interface {
	BuildA() interface{}
	BuildB() interface{}
	BuildC() interface{}
	Product() *Product
}

// Product 建造者模式最终建造的"产品"
type Product struct {
	partA interface{}
	partB interface{}
	partC interface{}
}

// ConcreteBuilder 具体的建造者，通过实现Builder接口来构建Product
type ConcreteBuilder struct{}

func (c *ConcreteBuilder) BuildA() interface{} {
	// 这里是构建某个部分的具体实现
	return fmt.Sprintf("build A")
}

func (c *ConcreteBuilder) BuildB() interface{} {
	// 这里是构建某个部分的具体实现
	return fmt.Sprintf("build B")
}
func (c *ConcreteBuilder) BuildC() interface{} {
	// 这里是构建某个部分的具体实现
	return fmt.Sprintf("build C")
}

func (c *ConcreteBuilder) Product() *Product {
	var p = &Product{
		partA: c.BuildA(),
		partB: c.BuildB(),
		partC: c.BuildC(),
	}
	return p
}

// Director 构建者模式中的指挥官，通过调用Builder的方法来构建最终的Product
type Director struct {
	builder Builder // 构建者
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) GetProduct() *Product {
	return d.builder.Product()
}

func main() {
	director := NewDirector(&ConcreteBuilder{})
	director.GetProduct()
}
