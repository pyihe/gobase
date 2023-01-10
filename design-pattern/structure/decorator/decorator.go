package main

import "fmt"

/*
	装饰器模式:
	动态的将新功能(附加功能)附加到对象(核心功能)上

	比如唱歌(核心功能), 在唱歌前可以有不同的准备(附加功能)
*/

type Singer interface {
	Sing(song string)
}

type Decorator func(Singer) Singer

type singer struct {
	m map[interface{}]interface{}
}

func (s *singer) Sing(song string) {
	fmt.Println("singer is singing ", song)
}

func Decorate(sgr Singer, ds ...Decorator) {
	s := sgr
	for _, d := range ds {
		d(s)
	}
}

func SadDecorator() Decorator {
	return func(s Singer) Singer {
		fmt.Println("singer is sad")
		return s
	}
}

func HappyDecorator() Decorator {
	return func(s Singer) Singer {
		fmt.Println("singer is happy")
		return s
	}
}

func PopMusic() Decorator {
	return func(s Singer) Singer {
		fmt.Println("singer will sing pop music")
		return s
	}
}

func main() {
	s := &singer{}
	Decorate(s, HappyDecorator(), PopMusic())
	s.Sing("《东风破》")
}
