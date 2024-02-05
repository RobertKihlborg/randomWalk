package walk

import (
	"fmt"
	"math"
	"strings"
)

type Number interface {
	int64 | int32 | int8 | float64 | float32
}
type VectorInt []int64

func (v VectorInt) Get(i int) int64 {
	return v[i]
}

func (v VectorInt) Size() int {
	return len(v)
}

func (v VectorInt) Add(v2 Vector[int64]) Vector[int64] {
	res := make(VectorInt, len(v))
	for i := 0; i < len(v); i++ {
		res[i] = v[i] + v2.Get(i)
	}
	return res
}

type VectorFloat []float64

func (v VectorFloat) Get(i int) float64 {
	return v[i]
}

func (v VectorFloat) Size() int {
	return len(v)
}

func (v VectorFloat) Add(v2 Vector[float64]) Vector[float64] {
	res := make(VectorFloat, len(v))
	for i := 0; i < len(v); i++ {
		res[i] = v[i] + v2.Get(i)
	}
	return res
}

func SquareDist[T Number](a, b Vector[T]) float64 {
	var res float64
	size := a.Size()
	for i := 0; i < size; i++ {
		x := float64(a.Get(i) - b.Get(i))
		res += x * x
	}
	return res
}

func Dist[T Number](a, b Vector[T]) float64 {
	return math.Sqrt(SquareDist(a, b))
}

func SquareNorm[T Number](a Vector[T]) float64 {
	var res float64
	size := a.Size()
	for i := 0; i < size; i++ {
		x := float64(a.Get(i))
		res += x * x
	}
	return res
}

func Norm[T Number](a Vector[T]) float64 {
	return math.Sqrt(SquareNorm(a))
}

func FormatVector[T Number](vector Vector[T]) string {
	sb := strings.Builder{}
	sb.WriteString("(")
	N := vector.Size()
	if N > 0 {
		sb.WriteString(fmt.Sprint(vector.Get(0)))
		for i := 1; i < vector.Size(); i++ {
			sb.WriteString(fmt.Sprintf(", %v", vector.Get(i)))
		}
	}
	sb.WriteString(")")
	return sb.String()
}
