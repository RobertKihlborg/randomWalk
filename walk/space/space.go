package space

import (
	"fmt"
	"math"
	"math/rand"
	. "randomWalk/walk"
)

type VectorFloat []float64

func dot(a, b VectorFloat) float64 {
	if len(a) != len(b) {
		panic(fmt.Sprintf("Mismatched dimensions %v and %v in scalar product", len(a), len(b)))
	}
	sum := float64(0)
	for i := 0; i < len(a); i++ {
		sum += a[i] * b[i]
	}
	return sum
}

func squareNorm(a VectorFloat) float64 {
	return dot(a, a)
}

func norm(a VectorFloat) float64 {
	return math.Sqrt(squareNorm(a))
}

func add(a, b VectorFloat) VectorFloat {
	if len(a) != len(b) {
		panic(fmt.Sprintf("Mismatched dimensions %v and %v in scalar product", len(a), len(b)))
	}
	c := make(VectorFloat, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i] + b[i]
	}
	return c
}

func minus(a VectorFloat) VectorFloat {
	c := make(VectorFloat, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = -a[i]
	}
	return c
}

func sub(a, b VectorFloat) VectorFloat {
	return add(a, minus(b))
}

func mul(a VectorFloat, b float64) VectorFloat {
	c := make(VectorFloat, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = b * a[i]
	}
	return c
}

func squareDist(a, b VectorFloat) float64 {
	return squareNorm(sub(a, b))
}

func dist(a, b VectorFloat) float64 {
	return norm(sub(a, b))
}

func IsBallIntersecting(d float64, walk []VectorFloat) bool {
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

func CreateIntersectionFunction(d float64) CollisionChecker[VectorFloat] {
	return func(walk []VectorFloat) bool {
		return IsBallIntersecting(d, walk)
	}
}

func randomDirectionRS(dim int, rand *rand.Rand) VectorFloat {
	res := make(VectorFloat, dim)
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

func randomDirectionAngles(dim int, r *rand.Rand) VectorFloat {
	if dim == 1 {
		return VectorFloat{float64(r.Intn(2)*2 - 1)}
	}
	res := make(VectorFloat, dim)
	angles := make([]float64, dim-1)
	sins := make([]float64, dim-1)
	coss := make([]float64, dim-1)

	angles[0] = r.Float64() * 2 * math.Pi
	for i := 1; i < len(angles); i++ {
		angles[i] = math.Acos(r.Float64()*2 - 1)
	}
	for i := 0; i < len(angles); i++ {
		sins[i] = math.Sin(angles[i])
		coss[i] = math.Cos(angles[i])
	}
	var fac float64 = 1
	for i := 0; i < dim-1; i++ {
		res[i] = coss[i] * fac
		fac *= sins[i]
	}
	return res
}

func CreateRandomDirectionFunction(dim int) RandomFunction[VectorFloat] {
	if dim < 6 {
		return func(r *rand.Rand) VectorFloat {
			return randomDirectionRS(dim, r)
		}
	}
	return func(r *rand.Rand) VectorFloat {
		return randomDirectionAngles(dim, r)
	}
}

func CreateNonreturningRandomDirection(d float64, dim int) NonreturningRandomFunction[VectorFloat] {
	dSquare := d * d
	randDirFun := CreateRandomDirectionFunction(dim)
	return func(prev VectorFloat, r *rand.Rand) VectorFloat {
		for true {
			test := randDirFun(r)
			if squareNorm(add(prev, test)) > dSquare {
				return test
			}
		}
		panic("This should be completely unreachable, it's literally after a 'while true' loop")
	}
}
