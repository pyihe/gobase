package main

import "fmt"

// 工厂模式:
// 定义一个创建对象的接口，让其子类自己决定实例化哪一个工厂类，工厂模式使其创建过程延迟到子类进行

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
