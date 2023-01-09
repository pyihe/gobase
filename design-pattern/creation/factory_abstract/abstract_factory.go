package main

// 抽象工厂模式:
// 抽象工厂模式提供创建一些列相关或者相互依赖对象的接口，无需指定具体的类型
// 适用场景: 创建同一产品族不同产品的对象需要大量重复的代码
// 优点: 具体产品在应用层的代码隔离，无需关心创建的细节
// 缺点: 规定了所有可能被创建的产品类型; 增加来系统的抽象性和理解难度
//
// 抽象工厂模式从代码看起来与工厂方法模式差不多，但实际两者背后代表的思想是不同的:
// 1. 工厂方法模式代表的是每种产品一个工厂，当需要不同的产品实现时，创建不同的工厂
// 2. 抽象工厂模式代表的是每个产品族一个工厂，

type ProductFactory interface {
	GetPhone() Phone
	GetComputer() Computer
}

type Phone interface {
	Name() string
	Start()
	Shutdown()
	Call()
	SendSms()
}

type Computer interface {
	Name() string
	Start()
	Shutdown()
	Program()
}

func main() {
	var pf ProductFactory = &XiaomiFactory{}
	pf.GetPhone().Start()
	pf.GetComputer().Start()

	pf = &HuaweiFactory{}
	pf.GetPhone().Start()
	pf.GetComputer().Start()
}
