package main

import "fmt"

type HuaWeiPhone struct {
}

func (hw *HuaWeiPhone) Name() string {
	return "Huawei Phone"
}
func (hw *HuaWeiPhone) Start() {
	fmt.Println("Start Huawei Phone")
}
func (hw *HuaWeiPhone) Shutdown() {
	fmt.Println("Shutdown Huawei Phone")
}
func (hw *HuaWeiPhone) Call() {
	fmt.Println("Call Huawei Phone")
}
func (hw *HuaWeiPhone) SendSms() {
	fmt.Println("SendSms Huawei Phone")
}
