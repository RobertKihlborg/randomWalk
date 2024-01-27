package algorithms

import . "randomWalk/saw"

func DefaultZombieMapless(n int) *[]Vec2 {
	return ZombieMapless(n, 30)
}

func DefaultMapless(n int, stopWalking *bool) *[]Vec2 {
	return Mapless(n, 30, stopWalking)
}

func DefaultParallel(n int) *[]Vec2 {
	return Parallel(n, 30)
}

func ZombifiedReactiveMapless(n int) *[]Vec2 {
	return Pool(n, DefaultMapless, 1)
}

var ZombieWalkers = map[string]Walker{
	"simple": func(n int) *[]Vec2 {
		res, _ := SimpleSAW(n)
		return res
	},

	"mapless": DefaultZombieMapless,

	"zombified": ZombifiedReactiveMapless,

	"parallel": DefaultParallel,
}

var ReactiveWalkers = map[string]ReactiveWalker{
	"mapless": DefaultMapless,
}
