package grid

import (
	"fmt"
	"math/rand"
	. "randomWalk/walk"
	"randomWalk/walk/space"
)

type VectorInt []int64

func ToFloatVector(v VectorInt) space.Vector {
	res := make(space.Vector, len(v))
	for i, val := range v {
		res[i] = float64(val)
	}
	return res
}

func add(a VectorInt, b VectorInt) VectorInt {
	if len(a) != len(b) {
		panic(fmt.Sprintf("Mismatched dimensions %v and %v.", len(a), len(b)))
	}
	for i := 0; i < len(a); i++ {
		a[i] = a[i] + b[i]
	}
	return a
}

func equals(a VectorInt, b VectorInt) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isZero(a VectorInt) bool {
	for _, val := range a {
		if val != 0 {
			return false
		}
	}
	return true
}

func IsPointIntersecting(walk []VectorInt) bool {
	for i, p1 := range walk {
		for _, p2 := range walk[:i] {
			if equals(p1, p2) {
				return true
			}
		}
	}
	return false
}

func RandomDirectionN(n int, rand rand.Rand) VectorInt {
	res := make(VectorInt, n)
	index := rand.Intn(n)
	sign := rand.Intn(2)*2 - 1
	res[index] = int64(sign)

	return res
}

func NonreturningRandomDirection(prev VectorInt, r rand.Rand) VectorInt {
	n := len(prev)
	for true {
		test := RandomDirectionN(n, r)
		if !isZero(add(prev, test)) {
			return test
		}
	}
	panic("This should be completely unreachable, it's literally after a 'while true' loop")
}

func CreateRandomDirectionFunctionInt(n int) RandomFunction[VectorInt] {
	return func(r *rand.Rand) VectorInt {
		return RandomDirectionN(n, *r)
	}
}
