package main

import (
	"fmt"
)

/*
	外观模式: 为了隐藏系统的复杂性，外观模式提供一个统一的可访问接口为外部提供服务
	比如: 现在医院为了方便行动不便的病人看诊，推出专门的服务，该服务指定一名工作人员陪同病人，帮助病人进行各项预约以及排队等，而病人只需要去进行检查或者就诊即可
	这个特殊的陪同工作人员即可看作是对外的统一接口，病人不需要面对复杂的各项检查外的其他事物
*/

type Sicker struct {
	Name string
}

// Checker 每个部门提供的检查服务
type Checker interface {
	Queue(*Sicker) // 检查前需要排队
	Check(*Sicker) // 检查
}

// Escort 陪护人员提供协助
type Escort interface {
	Assist(*Sicker)
}

type UltrasoundDepartment struct {
}

func (u *UltrasoundDepartment) Check(sicker *Sicker) {
	fmt.Println(sicker.Name, "进行B超检查")
}

func (u *UltrasoundDepartment) Queue(sicker *Sicker) {
	fmt.Println(sicker.Name, "B超检查前排队")
}

type CtDepartment struct {
}

func (ct *CtDepartment) Check(sicker *Sicker) {
	fmt.Println(sicker.Name, "进行CT检查")
}

func (ct *CtDepartment) Queue(sicker *Sicker) {
	fmt.Println(sicker.Name, "CT检查前排队")
}

type Escorter struct {
	service []Checker
}

func newEscorter(services ...Checker) Escort {
	e := &Escorter{}
	e.service = append(e.service, services...)
	return e
}

func (e *Escorter) Assist(sicker *Sicker) {
	for _, service := range e.service {
		service.Queue(sicker)
		service.Check(sicker)
	}
}

func main() {
	escorter := newEscorter(&UltrasoundDepartment{}, &CtDepartment{})
	sicker := &Sicker{Name: "Joy Boy"}
	escorter.Assist(sicker)
}
