package main

import "fmt"

/*
	策略模式: 在策略模式中，一个类的行为或者算法可以在运行时进行更改；
	比如同样是编写一个程序，程序员可以用Golang实现，也可以用Java实现亦或其他语言进行实现
*/

type Strategy interface {
	Programming(code string)
}

type Golang struct {
}

func (golang *Golang) Programming(code string) {
	fmt.Printf("我是Golang, 我正在编写代码: %s\n", code)
}

type Java struct {
}

func (java *Java) Programming(code string) {
	fmt.Printf("我是Java, 我正在编写代码: %s\n", code)
}

type Worker struct {
	strategy Strategy
}

func newWorker(programmer Strategy) *Worker {
	return &Worker{strategy: programmer}
}

func (w *Worker) Program(code string) {
	w.strategy.Programming(code)
}

func main() {
	worker := newWorker(&Golang{})
	worker.Program("Hello World!")
}
