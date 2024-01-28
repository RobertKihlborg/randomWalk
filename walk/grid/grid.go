package grid

import (
	"fmt"
	"math/rand"
	. "randomWalk/walk"
	"randomWalk/walk/space"
)

type VectorInt []int64

func ToFloatVector(v VectorInt) space.VectorFloat {
	res := make(space.VectorFloat, len(v))
	for i, val := range v {
		res[i] = float64(val)
	}
	return res
}

func add(a VectorInt, b VectorInt) VectorInt {
	if len(a) != len(b) {
		panic(fmt.Sprintf("Mismatched dimensions %v and %v.", len(a), len(b)))
	}
	c := make(VectorInt, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i] + b[i]
	}
	return c
}

func isZero(a VectorInt) bool {
	for _, val := range a {
		if val != 0 {
			return false
		}
	}
	return true
}

func toKey(v VectorInt) string {
	return fmt.Sprint(v)
}

func IsPointIntersecting(walk []VectorInt) bool {
	visited := make(map[string]interface{})
	for _, p := range walk {
		key := toKey(p)
		_, alreadyVisited := visited[key]
		visited[key] = nil
		if alreadyVisited {
			return true
		}
	}
	return false
}

func RandomDirectionN(n int, rand *rand.Rand) VectorInt {
	res := make(VectorInt, n)
	index := rand.Intn(n)
	sign := rand.Intn(2)*2 - 1
	res[index] = int64(sign)

	return res
}

func NonreturningRandomDirection(prev VectorInt, r *rand.Rand) VectorInt {
	n := len(prev)
	for true {
		test := RandomDirectionN(n, r)
		if !isZero(add(prev, test)) {
			return test
		}
	}
	panic("This should be completely unreachable, it's literally after a 'while true' loop")
}

func CreateRandomDirectionFunctionDim(dim int) RandomFunction[VectorInt] {
	return func(r *rand.Rand) VectorInt {
		return RandomDirectionN(dim, r)
	}
}
