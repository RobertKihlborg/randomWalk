package algorithms

import (
	. "randomWalk/saw"
	"testing"
)

func TestParallel(t *testing.T) {
	ns := []int{3, 10, 20, 200}
	retries := 20
	var fun = func(n int) *[]Vec2 {
		pos := Parallel(n, 30)
		return pos
	}
	TestSAW(t, fun, ns, retries)
}
