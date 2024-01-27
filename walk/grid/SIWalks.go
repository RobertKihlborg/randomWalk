package grid

import (
	"fmt"
	"math/rand"
	"randomWalk/walk"
)

func SIWalk2D(n int) []Vec2 {
	res := make([]Vec2, n+1)
	r := rand.New(rand.NewSource(rand.Int63()))

	for i := 1; i < n+1; i++ {
		res[i] = res[i-1] + UnitVectors2D[r.Intn(4)]
	}
	return res
}

func SIWalk3D(n int) []Vec3 {
	res := make([]Vec3, n+1)
	r := rand.New(rand.NewSource(rand.Int63()))

	for i := 1; i < n+1; i++ {
		res[i] = res[i-1] + UnitVectors3D[r.Intn(6)]
	}
	return res
}

func SIWalk4D(n int) []Vec4 {
	res := make([]Vec4, n+1)
	r := rand.New(rand.NewSource(rand.Int63()))

	for i := 1; i < n+1; i++ {
		res[i] = res[i-1] + UnitVectors4D[r.Intn(8)]
	}
	return res
}

func CreateSIWalker(dim int) walk.Walker[VectorInt] {
	if dim < 1 {
		panic(fmt.Sprintf("Invalid dimension %v", dim))
	}
	randFun := CreateRandomDirectionFunctionInt(dim)

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
