package space

import (
	"fmt"
	"testing"
)

func TestSIWalkGeneral(t *testing.T) {
	var ns = []int{0, 1, 2, 3, 5, 10, 100}
	var dims = []int{1, 2, 3, 10, 50, 100}
	for _, dim := range dims {
		walker := CreateSIWalker(dim)
		t.Run(fmt.Sprintf("Generalized algorithm %v dimensions", dim), func(t *testing.T) {
			for _, n := range ns {
				walker(n)
			}
		})
	}
}

func BenchmarkSIWalkGeneral(b *testing.B) {
	var dims = []int{1, 2, 3, 10, 50, 100}

	n := 10000
	for _, dim := range dims {
		walker := CreateSIWalker(dim)
		b.Run(fmt.Sprintf("Generalized algorithm %vD", dim), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				walker(n)
			}
		})
	}

}
