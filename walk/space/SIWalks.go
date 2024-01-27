package space

import (
	"fmt"
	"math/rand"
	. "randomWalk/walk"
)

func CreateSIWalker(dim int) Walker[Vector] {
	if dim < 1 {
		panic(fmt.Sprintf("Invalid dimension %v", dim))
	}
	randFun := CreateRandomDirectionFunction(dim)

	return func(n int) []Vector {
		res := make([]Vector, n+1)
		res[0] = make(Vector, dim)
		r := rand.New(rand.NewSource(rand.Int63()))

		for i := 1; i < n+1; i++ {
			res[i] = add(res[i-1], randFun(r))
		}
		return res
	}
}
