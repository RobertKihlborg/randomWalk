package grid

import (
	"fmt"
	"math/rand"
	"randomWalk/walk"
)

func CreateSIWalkerFix[T Vec2 | Vec3 | Vec4](unitVectors []T) walk.Walker[T] {
	unitVecCount := len(unitVectors)
	return func(n int) []T {
		res := make([]T, n+1)
		r := rand.New(rand.NewSource(rand.Int63()))

		for i := 1; i < n+1; i++ {
			res[i] = res[i-1] + unitVectors[r.Intn(unitVecCount)]
		}
		return res
	}
}

func CreateSIWalker(dim int) walk.Walker[VectorInt] {
	if dim < 1 {
		panic(fmt.Sprintf("Invalid dimension %v", dim))
	}
	randFun := CreateRandomDirectionFunctionDim(dim)

	return func(n int) []VectorInt {
		res := make([]VectorInt, n+1)
		res[0] = make(VectorInt, dim)
		r := rand.New(rand.NewSource(rand.Int63()))

		for i := 1; i < n+1; i++ {
			res[i] = add(res[i-1], randFun(r))
		}
		return res
	}
}
