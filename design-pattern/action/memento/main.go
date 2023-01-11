package main

/*
	备忘录模式: 在不破坏封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态
	备忘录模式中有三个角色
	1. Originator 发起备忘的人
	2. Memento 备忘录
	3. Caretaker 管理备忘录

	备忘录的特点是将需要备忘的属性封装在备忘录中，而不是直接透过发起者暴露给外部
	备忘录模式比较适合于功能比较复杂，但需要维护或者记录属性历史的场景，比如撤销操作

	如果需要记录的属性太多，则备忘录模式的备忘录对象相对而言会很耗内存
*/

// Originator 发起人，备忘录需要记录的事物
type Originator struct {
	state string // 发起人的属性
}

// CreateMemento 发起者创建备忘录
func (o *Originator) CreateMemento() *Memento {
	return &Memento{state: o.state}
}

// ApplyMemento 应用备忘录，将相关属性恢复到备忘录记录的状态
func (o *Originator) ApplyMemento(memento *Memento) {
	if memento == nil {
		return
	}
	o.state = memento.state
}

type Memento struct {
	state string // 需要备忘的属性
}

type Caretaker struct {
	Memento *Memento // 管理者管理的备忘录
}

func main() {
	orig := &Originator{
		state: "init",
	}
	// 生成备忘录
	taker := &Caretaker{
		Memento: orig.CreateMemento(),
	}
	// 状态变更
	orig.state = "mid"

	// 恢复备忘录
	orig.ApplyMemento(taker.Memento)
}
