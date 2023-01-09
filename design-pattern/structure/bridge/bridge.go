package main

import "fmt"

/*
	桥接模式:
	将抽象部分与实现部分分离, 使两者都可以独立的变化。


												 电脑
							/					  ｜                      \
						  台式					 平板					笔记本
					|      |       |        |     |      |          |     |      |
				   联想   苹果     戴尔     联想   苹果    戴尔       联想   苹果    戴尔

	一台电脑由类型和品牌构成，如何让品牌和类型的维护相互独立，在增加品牌或者类型时不对既有代码以及另一方造成影响:
	使用桥接模式; golang中体现为组合
*/

// Brand 品牌
type Brand interface {
	BrandInfo() string
}

type Apple struct {
}

func (apple *Apple) BrandInfo() string {
	return "Apple"
}

type Lenovo struct {
}

func (l *Lenovo) BrandInfo() string {
	return "Lenovo"
}

// Computer 电脑
type Computer interface {
	Brand
	ComputerInfo() string
}

type Pad struct {
	brand Brand
}

func newPad(brand Brand) Computer {
	return &Pad{brand: brand}
}

func (p *Pad) BrandInfo() string {
	return p.brand.BrandInfo()
}

func (p *Pad) ComputerInfo() string {
	return p.brand.BrandInfo() + "平板"
}

type Desktop struct {
	brand Brand
}

func newDesktop(brand Brand) Computer {
	return &Desktop{
		brand: brand,
	}
}

func (d *Desktop) BrandInfo() string {
	return d.brand.BrandInfo()
}

func (d *Desktop) ComputerInfo() string {
	return d.brand.BrandInfo() + "台式机"
}

func main() {
	var c = newPad(&Apple{})
	fmt.Println(c.ComputerInfo())
	c = newDesktop(&Lenovo{})
	fmt.Println(c.ComputerInfo())
}
