package walk

import (
	"math/rand"
)

func CreateSIWalker[T any](zero Vector[T], randFun RandomFunction[T]) Walker[T] {
	return func(n int) []Vector[T] {
		res := make([]Vector[T], n+1)
		res[0] = zero
		r := rand.New(rand.NewSource(rand.Int63()))

		for i := 1; i < n+1; i++ {
			res[i] = res[i-1].Add(randFun(r))
		}
		return res
	}
}

func CreateNaivestSAWalker[T any](zero Vector[T], randFun RandomFunction[T], collFun CollisionChecker[T]) Walker[T] {
	baseWalker := CreateSIWalker(zero, randFun)
	return func(n int) []Vector[T] {
	newAttempt:
		walk := baseWalker(n)
		if collFun(walk) {
			goto newAttempt
		}
		return walk
	}
}

func CreateNaiveSAWalker[T any](zero Vector[T], randFun RandomFunction[T], collFun PointCollisionChecker[T]) Walker[T] {
	return func(n int) []Vector[T] {
		res := make([]Vector[T], n+1)
		res[0] = zero
		r := rand.New(rand.NewSource(rand.Int63()))

	newAttempt:
		for i := 1; i < n+1; i++ {
			res[i] = res[i-1].Add(randFun(r))
			if collFun(res[i], res[:i]) {
				goto newAttempt
			}
		}
		return res
	}
}

func CreateBasicSAWalker[T any](zero Vector[T], randFun NonreturningRandomFunction[T], collFun PointCollisionChecker[T]) Walker[T] {
	return func(n int) []Vector[T] {
		res := make([]Vector[T], n+1)
		res[0] = zero
		r := rand.New(rand.NewSource(rand.Int63()))

	newAttempt:
		for i := 1; i < n+1; i++ {
			res[i] = res[i-1].Add(randFun(res[i-1], r))
			if collFun(res[i], res[:i]) {
				goto newAttempt
			}
		}
		return res
	}
}
