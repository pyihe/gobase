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

// 假设完成某些事需要A、B、C、D四个步骤，只是顺序不同，那么使用建造者模式

type Product interface {
	Value() string
}

/**********************************************************************************************************************/

type Builder interface {
	BuildA() Builder
	BuildB() Builder
	BuildC() Builder
	Product() Product
}

/**********************************************************************************************************************/

type aProduct struct {
	v string
}

func (p *aProduct) Value() string {
	return p.v
}

type aBuilder struct {
	product *aProduct
}

func (ab *aBuilder) BuildA() Builder {
	if ab.product == nil {
		ab.product = &aProduct{}
	}
	ab.product.v = fmt.Sprintf("%sa", ab.product.v)
	return ab
}
func (ab *aBuilder) BuildB() Builder {
	if ab.product == nil {
		ab.product = &aProduct{}
	}
	ab.product.v = fmt.Sprintf("%sb", ab.product.v)
	return ab
}
func (ab *aBuilder) BuildC() Builder {
	if ab.product == nil {
		ab.product = &aProduct{}
	}
	ab.product.v = fmt.Sprintf("%sc", ab.product.v)
	return ab
}
func (ab *aBuilder) Product() Product {
	p := ab.product
	ab.product = nil
	return p
}

/**********************************************************************************************************************/

type bProduct struct {
	v string
}

func (bp *bProduct) Value() string {
	return bp.v
}

type bBuilder struct {
	product *bProduct
}

func (bb *bBuilder) BuildA() Builder {
	if bb.product == nil {
		bb.product = &bProduct{}
	}
	bb.product.v = fmt.Sprintf("%s1", bb.product.v)
	return bb
}
func (bb *bBuilder) BuildB() Builder {
	if bb.product == nil {
		bb.product = &bProduct{}
	}
	bb.product.v = fmt.Sprintf("%s2", bb.product.v)
	return bb
}
func (bb *bBuilder) BuildC() Builder {
	if bb.product == nil {
		bb.product = &bProduct{}
	}
	bb.product.v = fmt.Sprintf("%s3", bb.product.v)
	return bb
}
func (bb *bBuilder) Product() Product {
	p := bb.product
	bb.product = nil
	return p
}

/**********************************************************************************************************************/

func main() {
	var builder Builder = &aBuilder{}
	fmt.Println(builder.BuildA().BuildB().BuildC().Product().Value())

	builder = &bBuilder{}
	fmt.Println(builder.BuildC().BuildA().BuildB().Product().Value())
}
