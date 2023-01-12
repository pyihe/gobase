package main

import "fmt"

/*
	访问者模式: 访问者模式表示一个作用于某对象结构中的各元素的操作。它使你可以在不改变各元素的类的前提下定义作用于这些元素的新操作。

	访问者模式适用于数据结构相对稳定的系统，该模式将作用于结构上的操作和数据结构的耦合解脱开，使得操作集合可以相对自由的演化，
	访问者模式的目的就是要把操作从数据结构分离出来

	优点: 增加新的操作很容易(新的操作就是新的访问者，只需要增加新的访问者实现类即可)
	缺点: 对于数据结构的变更支持不友好

	访问者模式中有三个角色:
	1. Visitor(访问者类)
	2. Element(被访问者): 接受访问者的访问
	3. ObjectStructure: 对象结构
	比如：
	动物在面对不同场景时的表现是不一样的，这里的动物就是被访问者，而不同的场景则是访问者。
	当你需要增加不同的场景时只需要实现对应的访问者接口即可

	但是当你需要增加动物能够被访问的属性时，此时代码维护不是很方便
*/

type Visitor interface {
	GetTigerEmotion(*tiger)
	GetDogEmotion(*dog)
}

type Animal interface {
	Accept(visitor Visitor)
}

type dog struct{}

func (d *dog) Accept(visitor Visitor) {
	visitor.GetDogEmotion(d)
}

type tiger struct{}

func (t *tiger) Accept(visitor Visitor) {
	visitor.GetTigerEmotion(t)
}

type dangerous struct{}

func (d *dangerous) GetTigerEmotion(tiger2 *tiger) {
	fmt.Printf("老虎接受危险的挑战!\n")
}

func (d *dangerous) GetDogEmotion(dog2 *dog) {
	fmt.Printf("小狗接受危险的挑战\n")
}

type food struct {
}

func (h *food) GetTigerEmotion(tiger2 *tiger) {
	fmt.Printf("老虎面对食物时的表现\n")
}

func (h *food) GetDogEmotion(dog2 *dog) {
	fmt.Printf("小狗面对食物时的表现\n")
}

type ObjectStructure struct {
	animals []Animal
}

func (os *ObjectStructure) Add(animal ...Animal) {
	os.animals = append(os.animals, animal...)
}

func (os *ObjectStructure) Del(animal Animal) {
	for i, a := range os.animals {
		if a == animal {
			os.animals = append(os.animals[:i], os.animals[i+1:]...)
			break
		}
	}
}

func (os *ObjectStructure) Visit(visitor Visitor) {
	for _, animal := range os.animals {
		animal.Accept(visitor)
	}
}

func main() {
	objs := &ObjectStructure{}
	tiger1 := &tiger{}
	dog1 := &dog{}

	objs.Add(tiger1, dog1)

	danger := &dangerous{}
	objs.Visit(danger)

	meet := &food{}
	objs.Visit(meet)
}
