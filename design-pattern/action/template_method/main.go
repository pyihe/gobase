package main

import "fmt"

/*
	模板方法模式: 模板方法模式的应用主要是由父类定义一些列的方法，子类根据自己的需要来实现自己的方法内容，没有实现的则继承自父类
*/

/*
	假设完成一件事(或一个算法)需要3个步骤，分别是Step1, Step2, Step3
	父类定义好模板方法并定义好执行方法执行顺序，然后自己匿名组合父类并根据自己的需要来实现自己的方法内容，没有覆盖的方法将按照父类的算法来执行
	需要注意的是，Golang中采用组合的方式来变相的实现类似继承的功能
	模板方法模式中需要用到的是匿名继承，匿名继承中子类没有实现的方法在执行时将执行父类的方法
	但是模板方法需要方法由子类到父类的方向执行，所以需要额外在父类中包含一个由子类实现的模板方法
*/

type Template interface {
	Step1()
	Step2()
	Step3()
	Do()
}

type Super struct {
	template Template
}

func newSuper(template Template) *Super {
	return &Super{template: template}
}

func (s *Super) Step1() {
	fmt.Printf("执行父类的Step1\n")
}

func (s *Super) Step2() {
	fmt.Printf("执行父类的Step2\n")
}

func (s *Super) Step3() {
	fmt.Printf("执行父类的Step3\n")
}

func (s *Super) Do() {
	s.template.Step1()
	s.template.Step2()
	s.template.Step3()
}

type Sub struct {
	*Super
}

func (s *Sub) Step1() {
	fmt.Printf("执行子类的Step1\n")
}

// func (s *Sub)Step2() {
//
// }
// func (s *Sub)Step3() {
//
// }

// func (s *Sub) Do() {
//
// }

func main() {
	var (
		super = &Super{}
		sub   = &Sub{}
	)
	super.template = sub
	sub.Super = super
	sub.Do()
}
