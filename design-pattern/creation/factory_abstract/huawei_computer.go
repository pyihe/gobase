package main

import "fmt"

type HuaweiComputer struct{}

func (hw *HuaweiComputer) Name() string {
	return "Huawei Computer"
}
func (hw *HuaweiComputer) Start() {
	fmt.Println("Start Huawei Computer")
}
func (hw *HuaweiComputer) Shutdown() {
	fmt.Println(" Shutdown Huawei Computer")
}
func (hw *HuaweiComputer) Program() {
	fmt.Println("Huawei Computer Programming")
}
