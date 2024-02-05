package walk

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSIWalkers(t *testing.T) {
	ns := []int{0, 1, 2, 10, 100}
	dims := []int{1, 2, 10}

	for _, d := range dims {
		t.Run(fmt.Sprintf("Grid %vD", d), func(t *testing.T) {
			walker := CreateSIWalker[int64](make(VectorInt, d), func(rand *rand.Rand) Vector[int64] {
				return RandomGridDirection(d, rand)
			})
			for _, n := range ns {
				res := walker(n)
				if len(res) != n+1 {
					t.Fatalf("Expected %v, got %v points in walk", n+1, len(res))
				}
			}
		})

		t.Run(fmt.Sprintf("Space %vD", d), func(t *testing.T) {
			walker := CreateSIWalker[float64](make(VectorFloat, d), func(rand *rand.Rand) Vector[float64] {
				return RandomSpaceDirection(d, rand)
			})
			for _, n := range ns {
				res := walker(n)
				if len(res) != n+1 {
					t.Fatalf("Expected %v, got %v points in walk", n+1, len(res))
				}
			}
		})
	}
}
