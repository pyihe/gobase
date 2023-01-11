package main

import "fmt"

/*
	观察者模式: 在观察者模式中，氛围观察者和被观察者，其中被观察者维持这观察者列表，在被观察者发生变动时会主动通知观察者自己发生的变化
	比较常见的场景为发布与订阅，被观察者为消息发布者，观察者为消息订阅者
*/

// Observer 观察者接口
type Observer interface {
	Update(string)
}

type Publisher struct {
	obs []Observer
}

func (pub *Publisher) Notify() {
	for _, ob := range pub.obs {
		ob.Update("publish")
	}
}

type SubscriberA struct {
}

func (s *SubscriberA) Update(content string) {
	fmt.Println("SubscriberA receive update: ", content)
}

type SubscriberB struct {
}

func (s *SubscriberB) Update(content string) {
	fmt.Println("SubscriberA receive update: ", content)
}

func main() {
	sa := &SubscriberA{}
	sb := &SubscriberA{}
	pub := &Publisher{}
	pub.obs = append(pub.obs, sa, sb)

	pub.Notify()
}
