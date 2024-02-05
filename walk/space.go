package walk

import (
	"math"
	"math/rand"
)

func (v VectorFloat) mul(b float64) VectorFloat {
	c := make(VectorFloat, len(v))
	for i := 0; i < len(v); i++ {
		c[i] = b * v[i]
	}
	return c
}

func IsBallIntersecting(d float64, walk []Vector[float64]) bool {
	dSquare := d * d
	for i, p1 := range walk {
		for _, p2 := range walk[:i] {
			if SquareDist(p1, p2) < dSquare {
				return true
			}
		}
	}
	return false
}

func IsLatestBallIntersecting(d float64, point Vector[float64], walk []Vector[float64]) bool {
	dSquare := d * d
	for _, p1 := range walk {
		if SquareDist(p1, point) < dSquare {
			return true
		}
	}
	return false
}

func RandomSpaceDirection(dim int, rand *rand.Rand) Vector[float64] {
	res := make(VectorFloat, dim)
	sn := float64(0)

newAttempt:
	sn = 0
	for i := 0; i < dim; i++ {
		res[i] = rand.Float64()*2 - 1
		sn += res[i] * res[i]

		if sn > 1 {
			goto newAttempt
		}
	}

	res = res.mul(1 / math.Sqrt(sn))
	return res
}

func NonreturningSpaceDirection(d float64, dim int, prev Vector[float64], rand *rand.Rand) Vector[float64] {
	dSquare := d * d
newAttempt:
	res := RandomSpaceDirection(dim, rand)
	if SquareNorm(prev.Add(res)) < dSquare {
		goto newAttempt
	}
	return res
}
