package adapter

/*
	Singer接口: 可以唱歌
	Jumper接口: 可以跳跃

	现在我想跳着唱歌，或者不管唱歌或者跳跃，每次都只需要通过一次调用就能达到效果
	这是可以用适配器在中间进行适配
*/

type Singer interface {
	Sing()
}

type Jumper interface {
	Jump()
}

type Adapter struct {
	Singer
	Jumper
}

func (a *Adapter) Act(sing, jump bool) {
	if sing {
		a.Sing()
	}
	if jump {
		a.Jump()
	}
}
