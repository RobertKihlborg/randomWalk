package grid

import (
	"fmt"
	"testing"
)

func TestSIWalkFix(t *testing.T) {
	var ns = []int{0, 1, 2, 3, 5, 10, 100, 1000, 1_000_000}

	w2d := CreateSIWalkerFix(UnitVectors2D)
	w3d := CreateSIWalkerFix(UnitVectors3D)
	w4d := CreateSIWalkerFix(UnitVectors4D)

	t.Run(fmt.Sprintf("Fix 2D"), func(t *testing.T) {
		for _, n := range ns {
			w2d(n)
		}
	})
	t.Run(fmt.Sprintf("Fix 3D"), func(t *testing.T) {
		for _, n := range ns {
			w3d(n)
		}
	})
	t.Run(fmt.Sprintf("Fix 4D"), func(t *testing.T) {
		for _, n := range ns {
			w4d(n)
		}
	})
}

func TestSIWalkGeneral(t *testing.T) {
	var dims = []int{1, 2, 3, 10, 100}
	var ns = []int{0, 1, 2, 3, 5, 10, 100, 1000, 1_000_000}

	for _, dim := range dims {
		walker := CreateSIWalker(dim)
		t.Run(fmt.Sprintf("Generalized algorithm %v dimensions", dim), func(t *testing.T) {
			for _, n := range ns {
				walker(n)
			}
		})
	}
}
func BenchmarkSIWalkFix(b *testing.B) {
	n := 10000
	w2d := CreateSIWalkerFix(UnitVectors2D)
	w3d := CreateSIWalkerFix(UnitVectors3D)
	w4d := CreateSIWalkerFix(UnitVectors4D)

	b.Run(fmt.Sprintf("Fix 2D"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			w2d(n)
		}
	})
	b.Run(fmt.Sprintf("Fix 3D"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			w3d(n)
		}
	})
	b.Run(fmt.Sprintf("Fix 4D"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			w4d(n)
		}
	})
}

func BenchmarkSIWalkGeneral(b *testing.B) {
	var dims = []int{1, 2, 3, 10, 100}
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
