package main

type XiaomiFactory struct {
}

func (xf *XiaomiFactory) GetPhone() Phone {
	return &XiaomiPhone{}
}

func (xf *XiaomiFactory) GetComputer() Computer {
	return &XiaomiComputer{}
}
