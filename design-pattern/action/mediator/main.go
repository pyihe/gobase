package main

import "fmt"

/*
	中介者模式: 中介者模式中用一个中介者对象来封装一系列的对象交互。中介者使各对象不需要显示的相互引用，从而使其耦合松散，而且可以独立地改变他们之间的交互

	比如联合国与世界各国之间的关系可视为中介者模式，其中联合国充当中介者，其他各国则为相互独立的对象
*/

// Mediator 中介者
type Mediator interface {
	SendMsg(msg string, colleague Colleague)
	AddColleague(cs ...Colleague)
}

type Colleague interface {
	GetMsg(msg string)
	SendMsg(msg string)
}

type UnitedNation struct {
	colleagues []Colleague
}

func (un *UnitedNation) SendMsg(msg string, colleague Colleague) {
	fmt.Printf("UnitedNation 收到消息: %v\n", msg)
	for _, c := range un.colleagues {
		if c != colleague {
			c.GetMsg(msg)
		}
	}
}

func (un *UnitedNation) AddColleague(cs ...Colleague) {
	un.colleagues = append(un.colleagues, cs...)
}

type Country1 struct {
	m Mediator
}

func newCountry1(m Mediator) Colleague {
	return &Country1{m: m}
}

func (c *Country1) GetMsg(m string) {
	fmt.Printf("Country1 收到消息: %s\n", m)
}

func (c *Country1) SendMsg(m string) {
	c.m.SendMsg(m, c)
}

type Country2 struct {
	m Mediator
}

func (c *Country2) GetMsg(m string) {
	fmt.Printf("Country2 收到消息: %s\n", m)
}

func (c *Country2) SendMsg(m string) {
	c.m.SendMsg(m, c)
}

func main() {
	var (
		m  Mediator = &UnitedNation{}
		c1          = &Country1{m: m}
		c2          = &Country2{m: m}
	)
	m.AddColleague(c1, c2)

	c1.SendMsg("hello c2")
	c2.SendMsg("hi c1")
}
