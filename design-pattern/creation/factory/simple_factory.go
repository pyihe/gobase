package main

import "fmt"

// 简单工厂模式:
// 定义一个创建对象的接口，传递类型参数，由工厂接口根据类型创建出具体的实例

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

type InstanceType string

const (
	HumanInstance   InstanceType = "HUMAN"
	AnimalInstance  InstanceType = "ANIMAL"
	MachineInstance InstanceType = "MACHINE"
)

// NewInstance 工厂模式，得到的实例由传入的传入的实例类型来决定，每当有新增的实例类型时，NewInstance也需要添加相应的实例类型
// 优点: 代码易于管理
// 缺点: 每添加一个类型需要改动原有的代码
func NewInstance(instanceType InstanceType) Worker {
	switch instanceType {
	case HumanInstance:
		return &human{}
	case AnimalInstance:
		return &animal{}
	case MachineInstance:
		return &machine{}
	default:
		return nil
	}
}

func main() {
	NewInstance(HumanInstance).Run()
}
