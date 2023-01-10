package main

import "fmt"

/*
	组合模式: 创建一个包含自己的对象组的类（树形结构）
*/

type Employee struct {
	Name         string
	Dept         string
	Salary       int
	Subordinates []*Employee // 下属
}

func (e *Employee) String() string {
	return e.Name
}

func (e *Employee) AddSub(employee ...*Employee) {
	e.Subordinates = append(e.Subordinates, employee...)
}

func (e *Employee) PrintSubordinates() {
	if len(e.Subordinates) == 0 {
		return
	}
	for _, sub := range e.Subordinates {
		fmt.Println(sub)
		sub.PrintSubordinates()
	}
}

func newEmployee(name, dept string) *Employee {
	return &Employee{
		Name: name,
		Dept: dept,
	}
}

func main() {
	//
	ceo := newEmployee("CEO", "CEO")
	marketDirector := newEmployee("Market Director", "Market Director")
	financialDirector := newEmployee("Financial Director", "Financial Director")

	ceo.AddSub(marketDirector, financialDirector)

	marketStaff1 := newEmployee("Market Staff1", "Market Staff")
	marketStaff2 := newEmployee("Market Staff2", "Market Staff")
	marketDirector.AddSub(marketStaff1, marketStaff2)

	financialStaff1 := newEmployee("Financial Staff1", "Financial Staff")
	financialStaff2 := newEmployee("Financial Staff2", "Financial Staff")
	financialDirector.AddSub(financialStaff1, financialStaff2)

	ceo.PrintSubordinates()
}
