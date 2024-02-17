package pkg

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

type VectorInt []int

func (v VectorInt) String() string {
	elems := make([]string, len(v))
	for i, p := range v {
		elems[i] = fmt.Sprint(p)
	}
	return strings.Join(elems, ", ")
}
func (v VectorInt) add(v2 VectorInt) VectorInt {
	if len(v2) != len(v) {
		panic("Mismatched dimensions")
	}
	res := make(VectorInt, len(v))
	for i := 0; i < len(v); i++ {
		res[i] = v[i] + v2[i]
	}

	return res
}
func (v VectorInt) norm() float64 {
	sum := 0
	for _, p := range v {
		sum += p * p
	}
	return math.Sqrt(float64(sum))
}

type GridWalk []VectorInt

func (w GridWalk) Distance() float64 {
	return w[len(w)-1].norm()
}

func (w GridWalk) String() string {
	b := new(strings.Builder)
	for _, vec := range w {
		b.WriteString(vec.String())
		b.WriteRune('\n')
	}
	return b.String()
}

func basicIGridWalk(dim, n int) Walk {
	res := make(GridWalk, n+1)
	res[0] = make(VectorInt, dim)

	r := rand.New(rand.NewSource(rand.Int63()))

	for i := 1; i < n+1; i++ {
		res[i] = res[i-1].add(randomGridDirection(dim, r))
	}

	return res
}

func randomGridDirection(dim int, rand *rand.Rand) VectorInt {
	res := make(VectorInt, dim)
	index := rand.Intn(dim)
	sign := rand.Intn(2)*2 - 1
	res[index] = sign

	return res
}
