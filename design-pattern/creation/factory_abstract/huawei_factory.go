package main

type HuaweiFactory struct {
}

func (hf *HuaweiFactory) GetPhone() Phone {
	return &HuaWeiPhone{}
}

func (hf *HuaweiFactory) GetComputer() Computer {
	return &HuaweiComputer{}
}
