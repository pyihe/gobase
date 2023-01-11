package main

import "fmt"

/*
	责任链模式: 责任链模式主要用于处理链式流程，比如流程审批；用户提交审批申请后，一个环节审批完成后传递给下一环节进行审批
*/

type Chain interface {
	Approve()
	Next() Chain
}

type subChain struct {
	name string
	next *subChain
}

func newSubChain(name string, next *subChain) *subChain {
	return &subChain{
		name: name,
		next: next,
	}
}

func (sub *subChain) Approve() {
	fmt.Printf("%s进行审批\n", sub.name)
	if sub.next != nil {
		sub.next.Approve()
	}
}

func (sub *subChain) Next() Chain {
	if sub.next != nil {
		return sub.next
	}
	return nil
}

func main() {
	thirdChain := newSubChain("third", nil)
	secondChain := newSubChain("second", thirdChain)
	firstChain := newSubChain("first", secondChain)

	firstChain.Approve()
}
