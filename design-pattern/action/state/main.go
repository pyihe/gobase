package main

import "fmt"

/*
	状态模式: 类的行为根据其状态的改变而改变
*/

type State interface {
	Exec(*Ctx)
}

type stateA struct {
}

func (s *stateA) Exec(ctx *Ctx) {
	fmt.Printf("状态A Exec\n")
	ctx.SetState(s)
}

type stateB struct {
}

func (s *stateB) Exec(ctx *Ctx) {
	fmt.Printf("状态B Exec\n")
	ctx.SetState(s)
}

type Ctx struct {
	state State
}

func (c *Ctx) Exec() {
	c.state.Exec(c)
}

func (c *Ctx) SetState(state State) {
	c.state = state
}

func main() {
	ctx := &Ctx{}
	a := &stateA{}
	a.Exec(ctx)
	b := &stateB{}
	b.Exec(ctx)
}
