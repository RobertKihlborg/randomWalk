package grid

import (
	"fmt"
	"testing"
)

var ns = []int{0, 1, 2, 3, 5, 10, 100, 1000, 1_000_000}

func TestSIWalk2D(t *testing.T) {
	for _, n := range ns {
		SIWalk2D(n)
	}
}

func TestSIWalk3D(t *testing.T) {
	for _, n := range ns {
		SIWalk3D(n)
	}
}

func TestSIWalk4D(t *testing.T) {
	for _, n := range ns {
		SIWalk4D(n)
	}
}

var dims = []int{1, 2, 3, 10, 100}

func TestSIWalkGeneral(t *testing.T) {
	for _, dim := range dims {
		walker := CreateSIWalker(dim)
		t.Run(fmt.Sprintf("Generalized algorithm %v dimensions", dim), func(t *testing.T) {
			for _, n := range ns {
				walker(n)
			}
		})
	}
}
