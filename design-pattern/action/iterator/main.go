package main

/*
	迭代器模式: 提供一种方法顺序访问一个聚合对象中的各个元素，却又不暴露该对象的内部表示
*/

// Iterator 迭代器接口, 提供迭代需要的相应接口
type Iterator interface {
	First() interface{}
	Next() interface{}
	IsDone() bool
	CurrentItem() interface{}
}

// Aggregate 聚集接口, 用于创建迭代某一聚集集合的迭代器
type Aggregate interface {
	CreateIterator() Iterator
}

// ConcreteIterator 迭代器的实现类
type ConcreteIterator struct {
	aggregate *ConcreteAggregate
	current   int
}

func newConcreteIterator(aggregate *ConcreteAggregate) Iterator {
	return &ConcreteIterator{
		aggregate: aggregate,
		current:   0,
	}
}

func (it *ConcreteIterator) First() interface{} {
	return it.aggregate.items[0]
}

func (it *ConcreteIterator) Next() interface{} {
	it.current += 1
	if it.current < it.aggregate.Len() {
		return it.aggregate.items[it.current]
	}
	return nil
}
func (it *ConcreteIterator) IsDone() bool {
	return it.current >= it.aggregate.Len()
}
func (it *ConcreteIterator) CurrentItem() interface{} {
	return it.aggregate.items[it.current]
}

// ConcreteAggregate 聚集集合的实现类
type ConcreteAggregate struct {
	items []interface{}
}

func (ca *ConcreteAggregate) CreateIterator() Iterator {
	return newConcreteIterator(ca)
}

func (ca *ConcreteAggregate) Len() int {
	return len(ca.items)
}

func (ca *ConcreteAggregate) Index(idx int) interface{} {
	return ca.items[idx]
}
