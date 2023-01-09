package main

import "fmt"

/*
	两个需要没有关联但需要交互的类型之间通过一个适配器进行交互
	比如国内外的插座，需要一个转接器才能使电器正常工作
*/

// Adapter 适配器
type Adapter interface {
	Adapt()
}

type adapter struct {
	*AmericanStandard
}

func newAdapter(a *AmericanStandard) Adapter {
	return &adapter{a}
}

func (ad *adapter) Adapt() {
	ad.Work()
}

type AmericanStandard struct{}

func (a *AmericanStandard) Work() {
	fmt.Println("美标可以正常工作了")
}

type NationalStandard struct {
}

func (n *NationalStandard) Work(ad ...Adapter) {
	fmt.Println("国标可以工作了")
	if len(ad) > 0 {
		ad[0].Adapt()
	}
}

func main() {
	// 国标和美标必须同时工作电器才能正常使用
	ns := &NationalStandard{}
	as := &AmericanStandard{}
	adapter := newAdapter(as)
	ns.Work(adapter)
}
