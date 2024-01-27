package space

import (
	"fmt"
	"testing"
)

var ns = []int{0, 1, 2, 3, 5, 10, 100}

var dims = []int{1, 2, 3, 10}

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
