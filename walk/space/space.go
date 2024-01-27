package space

import (
	"fmt"
	"math"
	"math/rand"
	. "randomWalk/walk"
)

type Vector []float64

func dot(a, b Vector) float64 {
	if len(a) != len(b) {
		panic(fmt.Sprintf("Mismatched dimensions %v and %v in scalar product", len(a), len(b)))
	}
	sum := float64(0)
	for i := 0; i < len(a); i++ {
		sum += a[i] * b[i]
	}
	return sum
}

func squareNorm(a Vector) float64 {
	return dot(a, a)
}

func norm(a Vector) float64 {
	return math.Sqrt(squareNorm(a))
}

func add(a, b Vector) Vector {
	if len(a) != len(b) {
		panic(fmt.Sprintf("Mismatched dimensions %v and %v in scalar product", len(a), len(b)))
	}
	n := len(a)
	for i := 0; i < n; i++ {
		a[i] = a[i] + b[i]
	}
	return a
}

func minus(a Vector) Vector {
	for i := 0; i < len(a); i++ {
		a[i] = -a[i]
	}
	return a
}

func sub(a, b Vector) Vector {
	return add(a, minus(b))
}

func mul(a Vector, b float64) Vector {
	for i := 0; i < len(a); i++ {
		a[i] = b * a[i]
	}
	return a
}

func squareDist(a, b Vector) float64 {
	return squareNorm(sub(a, b))
}

func dist(a, b Vector) float64 {
	return norm(sub(a, b))
}

func IsBallIntersecting(d float64, walk []Vector) bool {
	dSquare := d * d
	for i, p1 := range walk {
		for _, p2 := range walk[:i] {
			if squareDist(p1, p2) < dSquare {
				return true
			}
		}
	}
	return false
}

func CreateIntersectionFunction(d float64) CollisionChecker[Vector] {
	return func(walk []Vector) bool {
		return IsBallIntersecting(d, walk)
	}
}

func RandomDirectionN(dim int, rand *rand.Rand) Vector {
	res := make(Vector, dim)
	//counter := 0
	for true {
		for i := 0; i < dim; i++ {
			res[i] = rand.Float64()*2 - 1
		}

		if squareNorm(res) > 1 {
			//counter++
			continue
		}
		res = mul(res, 1/norm(res))
		//log.Println(fmt.Sprintf("rejection sampling counter: %v", counter))
		return res
	}
	panic("This should be completely unreachable, it's literally after a 'while true' loop")
}
func CreateRandomDirectionFunction(n int) RandomFunction[Vector] {
	return func(r *rand.Rand) Vector {
		return RandomDirectionN(n, r)
	}
}

func CreateNonreturningRandomDirection(d float64) NonreturningRandomFunction[Vector] {
	dSquare := d * d
	return func(prev Vector, r *rand.Rand) Vector {
		n := len(prev)
		for true {
			test := RandomDirectionN(n, r)
			if squareNorm(add(prev, test)) > dSquare {
				return test
			}
		}
		panic("This should be completely unreachable, it's literally after a 'while true' loop")
	}
}
