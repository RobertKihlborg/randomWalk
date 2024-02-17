package pkg

import (
	"fmt"
	"strings"
)

type VectorFloat []float64

func (v VectorFloat) String() string {
	elems := make([]string, len(v))
	for i, p := range v {
		elems[i] = fmt.Sprintf("%.3f", p)
	}
	return strings.Join(elems, ", ")
}

type SpaceWalk []VectorFloat

func (w SpaceWalk) String() string {
	b := new(strings.Builder)
	for _, vec := range w {
		b.WriteString(vec.String())
		b.WriteRune('\n')
	}
	return b.String()
}
