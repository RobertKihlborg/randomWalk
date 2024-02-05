package walk

import (
	"math/rand"
)

func isZero(v Vector[int64]) bool {
	for i := 0; i < v.Size(); i++ {
		if v.Get(i) != 0 {
			return false
		}
	}
	return true

}

func IsPointIntersecting(walk []VectorInt) bool {
	visited := make(map[string]interface{})
	for _, p := range walk {
		key := FormatVector[int64](p)
		_, alreadyVisited := visited[key]
		visited[key] = nil
		if alreadyVisited {
			return true
		}
	}
	return false
}

func RandomGridDirection(n int, rand *rand.Rand) Vector[int64] {
	res := make(VectorInt, n)
	index := rand.Intn(n)
	sign := rand.Intn(2)*2 - 1
	res[index] = int64(sign)

	return res
}

func NonreturningGridDirection(dim int, prev Vector[int64], r *rand.Rand) Vector[int64] {
newAttempt:
	test := RandomGridDirection(dim, r)
	if isZero(prev.Add(test)) {
		goto newAttempt
	}
	return test
}
