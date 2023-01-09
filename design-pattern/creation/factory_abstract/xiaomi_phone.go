package main

import "fmt"

type XiaomiPhone struct {
}

func (xm *XiaomiPhone) Name() string {
	return "Xiaomi Phone"
}
func (xm *XiaomiPhone) Start() {
	fmt.Println("Start Xiaomi Phone")
}
func (xm *XiaomiPhone) Shutdown() {
	fmt.Println("Shutdown Xiaomi Phone")
}
func (xm *XiaomiPhone) Call() {
	fmt.Println("Call Xiaomi Phone")
}
func (xm *XiaomiPhone) SendSms() {
	fmt.Println("SendSms Xiaomi Phone")
}
