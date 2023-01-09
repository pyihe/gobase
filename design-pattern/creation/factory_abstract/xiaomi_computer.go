package main

import "fmt"

type XiaomiComputer struct{}

func (xm *XiaomiComputer) Name() string {
	return "Xiaomi Computer"
}
func (xm *XiaomiComputer) Start() {
	fmt.Println("Start Xiaomi Computer")
}
func (xm *XiaomiComputer) Shutdown() {
	fmt.Println("Shutdown Xiaomi Computer")
}
func (xm *XiaomiComputer) Program() {
	fmt.Println("Xiaomi Computer Programming")
}
