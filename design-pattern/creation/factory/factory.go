package main

import "fmt"

// 工厂方法模式:
// 创建一个工厂接口，每个类型都有自己的工厂实例，最终由工厂来创建类型的实例
// 比如汽车，每个品牌的汽车都有自己的工厂，该工厂为一系列方法构成的接口

// Car 通过工厂最终获取的类型
type Car interface {
	Name() string
	Drive()
	Park()
}

/**********************************************************************************************************************/

// CarFactory 汽车工厂
type CarFactory interface {
	Name() string
	BuyCar() Car
}

/**********************************************************************************************************************/

type Focus struct{}

func (f *Focus) Name() string {
	return "Ford Focus"
}

func (f *Focus) Drive() {
	fmt.Println("focus is driving...")
}

func (f *Focus) Park() {
	fmt.Println("focus is parking...")
}

/**********************************************************************************************************************/

type FocusFactory struct{}

func (f *FocusFactory) Name() string {
	return "Ford Focus Factory"
}

func (f *FocusFactory) BuyCar() Car {
	return &Focus{}
}

/**********************************************************************************************************************/

type Peugeot struct{}

func (p *Peugeot) Name() string {
	return "Peugeot"
}

func (p *Peugeot) Drive() {
	fmt.Println("peugeot is driving...")
}

func (p *Peugeot) Park() {
	fmt.Println("peugeot is parking...")
}

/**********************************************************************************************************************/

type PeugeotFactory struct{}

func (pf *PeugeotFactory) Name() string {
	return "Peugeot Factory"
}

func (pf *PeugeotFactory) BuyCar() Car {
	return &Peugeot{}
}

func main() {
	var (
		focusFactory   CarFactory = &FocusFactory{}
		peugeotFactory CarFactory = &PeugeotFactory{}
		car            Car
	)

	car = focusFactory.BuyCar()
	car.Drive()
	car = peugeotFactory.BuyCar()
	car.Drive()
}
