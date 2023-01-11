package main

import "fmt"

/*
	命令模式: 将一个请求封装为一个对象，从而可以用不同的请求对客户端进行参数化，对请求排队或者记录请求日志以及支持可撤销的操作
	命令模式中包含三个角色:
	1. Command: (被执行的)命令
	2. Invoker: (传递命令)服务员
	3. Receiver: (执行命令)命令接收者
*/

// Command 命令
type Command struct {
	Receiver *Receiver // 命令接受者
}

func (c *Command) Exec() {
	c.Receiver.Action()
}

// Receiver 命令接受者，执行命令的角色
type Receiver struct {
}

func (r *Receiver) Action() {
	fmt.Printf("执行命令!\n")
}

// Invoker 传递命令的角色
type Invoker struct {
	command *Command
}

func (i *Invoker) SetCommand(c *Command) {
	i.command = c
}

func (i *Invoker) ExecCommand() {
	i.command.Exec()
}

func main() {
	var (
		receiver = &Receiver{}
		command  = &Command{receiver}
		invoker  = &Invoker{}
	)

	invoker.SetCommand(command)
	invoker.ExecCommand()
}
