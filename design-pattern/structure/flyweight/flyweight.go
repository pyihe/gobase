package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/*
	享元模式: 运用共享技术来有效地支持大量细粒度对象的复用；享元模式尝试重用现有的同类对象（池化）
	主要用于减少创建对象的数量，以减少内存占用和提高性能；

	现实生活中，共享汽车、共享单车都可以看作是享元模式的场景
*/

var (
	bikePool sync.Pool
	counter  int32
)

type SharedBike struct {
	renting bool // 是否在租用中
}

func GetBike() (bike *SharedBike) {
	x := bikePool.Get()
	if x == nil {
		atomic.AddInt32(&counter, 1)
		bike = &SharedBike{renting: true}
	} else {
		bike = x.(*SharedBike)
		bike.renting = true
	}
	return
}

func ReturnBike(bike *SharedBike) {
	if bike != nil {
		bike.renting = false
		bikePool.Put(bike)
	}
}

func main() {
	var (
		rad = rand.New(rand.NewSource(time.Now().UnixNano()))
		wg  = sync.WaitGroup{}
	)
	for i := 0; i <= 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			bike := GetBike()
			time.Sleep(time.Duration(rad.Intn(3)) * time.Second)
			ReturnBike(bike)
		}()
	}

	wg.Wait()
	fmt.Printf("一共创建了%d次对象\n", atomic.LoadInt32(&counter))
}
