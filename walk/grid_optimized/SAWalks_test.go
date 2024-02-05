package grid_optimized

import (
	"fmt"
	"randomWalk/walk"
	"testing"
)

func TestCreateNaiveWalkerFix(t *testing.T) {
	var ns = []int{0, 1, 2, 3, 5, 10, 20, 30}

	w2d := CreateNaiveWalkerFix(UnitVectors2D)
	w3d := CreateNaiveWalkerFix(UnitVectors3D)
	w4d := CreateNaiveWalkerFix(UnitVectors4D)
	t.Run(fmt.Sprintf("Fix 2D"), func(t *testing.T) {
		for _, n := range ns {
			res := w2d(n)
			if len(res) != (n+1) || IsPointIntersectingFix(res) {
				t.Fatalf("n=%v", n)
			}
		}
	})
	t.Run(fmt.Sprintf("Fix 3D"), func(t *testing.T) {
		for _, n := range ns {
			res := w3d(n)
			if len(res) != (n+1) || IsPointIntersectingFix(res) {
				t.Fatalf("n=%v", n)
			}
		}
	})
	t.Run(fmt.Sprintf("Fix 4D"), func(t *testing.T) {
		for _, n := range ns {
			res := w4d(n)
			if len(res) != (n+1) || IsPointIntersectingFix(res) {
				t.Fatalf("n=%v", n)
			}
		}
	})
}

func TestCreateNaiveWalker(t *testing.T) {
	var dims = []int{4, 10, 100}
	var ns = []int{0, 1, 2, 3, 5, 10, 20, 30, 40, 50}

	for _, dim := range dims {
		walker := CreateNaiveWalker(dim)
		t.Run(fmt.Sprintf("SAW %vD", dim), func(t *testing.T) {
			for _, n := range ns {
				res := walker(n)
				if len(res) != (n+1) || walk.IsPointIntersecting(res) {
					t.Fatalf("n=%v", n)
				}
			}
		})
	}
}

func TestCreateWalkerFix(t *testing.T) {
	var ns = []int{0, 1, 2, 30, 50, 80}

	w2d := CreateWalkerFix(UnitVectors2D)
	w3d := CreateWalkerFix(UnitVectors3D)
	w4d := CreateWalkerFix(UnitVectors4D)
	t.Run(fmt.Sprintf("Fix 2D"), func(t *testing.T) {
		for _, n := range ns {
			res := w2d(n)
			if len(res) != (n+1) || IsPointIntersectingFix(res) {
				t.Fatalf("n=%v", n)
			}
		}
	})
	t.Run(fmt.Sprintf("Fix 3D"), func(t *testing.T) {
		for _, n := range ns {
			res := w3d(n)
			if len(res) != (n+1) || IsPointIntersectingFix(res) {
				t.Fatalf("n=%v", n)
			}
		}
	})
	t.Run(fmt.Sprintf("Fix 4D"), func(t *testing.T) {
		for _, n := range ns {
			res := w4d(n)
			if len(res) != (n+1) || IsPointIntersectingFix(res) {
				t.Fatalf("n=%v", n)
			}
		}
	})
}

func TestCreateWalker(t *testing.T) {
	var dims = []int{4, 10, 100}
	var ns = []int{0, 1, 2, 3, 5, 10, 20, 30, 40, 50, 100, 300}

	for _, dim := range dims {
		walker := CreateWalker(dim)
		t.Run(fmt.Sprintf("SAW %vD", dim), func(t *testing.T) {
			for _, n := range ns {
				res := walker(n)
				if len(res) != (n+1) || walk.IsPointIntersecting(res) {
					t.Fatalf("n=%v", n)
				}
			}
		})
	}
}

func BenchmarkCreateWalkerFix(b *testing.B) {
	w2d := CreateWalkerFix(UnitVectors2D)
	w3d := CreateWalkerFix(UnitVectors3D)
	w4d := CreateWalkerFix(UnitVectors4D)
	n := 100
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
func BenchmarkCreateWalker(b *testing.B) {
	var dims = []int{2, 3, 4, 10, 100}
	n := 100
	for _, dim := range dims {
		walker := CreateWalker(dim)
		b.Run(fmt.Sprintf("Generalized algorithm %vD", dim), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				walker(n)
			}
		})
	}
}
