package main

import (
	"fmt"
)

// 抽象工厂模式:
// 围绕一个超级工厂创建其他工厂。该超级工厂又称为其他工厂的工厂

type Worker interface {
	Run()
}

type human struct{}

func (h *human) Run() {
	fmt.Println("human is running...")
}

type animal struct{}

func (a *animal) Run() {
	fmt.Println("animal is running...")
}

type machine struct{}

func (m *machine) Run() {
	fmt.Println("machine is running...")
}

type workerFactory struct{}

func (w *workerFactory) GetColor(typo string) Color {
	return nil
}

func (w *workerFactory) GetWorker(typo string) Worker {
	switch typo {
	case "human":
		return &human{}
	case "animal":
		return &animal{}
	default:
		return &machine{}
	}
}

/****************************************************************************************************************/

type Color interface {
	Color()
}

type red struct{}

func (r *red) Color() {
	fmt.Println("this is red...")
}

type blue struct{}

func (b *blue) Color() {
	fmt.Println("this is blue...")
}

type green struct{}

func (g *green) Color() {
	fmt.Println("this is green...")
}

type colorFactory struct{}

func (c *colorFactory) GetColor(typo string) Color {
	switch typo {
	case "red":
		return &red{}
	case "blue":
		return &blue{}
	default:
		return &green{}
	}
}

func (c *colorFactory) GetWorker(typo string) Worker {
	return nil
}

/********************************************************************************************************/

// AbstractFactory 超级工厂，用于生产不同的子工厂，每个子工厂可以根据类型来生成不同的接口实例
type AbstractFactory interface {
	GetColor(string) Color
	GetWorker(string) Worker
}

type factory struct{}

func NewInstance(name string) AbstractFactory {
	switch name {
	case "color":
		return &colorFactory{}
	default:
		return &workerFactory{}
	}
}

func main() {
	NewInstance("color").GetColor("red").Color()
	NewInstance("worker").GetWorker("human").Run()
}
