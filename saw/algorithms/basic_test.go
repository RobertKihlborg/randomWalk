package algorithms

import (
	. "randomWalk/saw"
	"testing"
)

func TestSimpleSAW(t *testing.T) {
	ns := []int{3, 10, 20, 50}
	retries := 20
	var fun = func(n int) *[]Vec2 {
		pos, _ := SimpleSAW(n)
		return pos
	}
	TestSAW(t, fun, ns, retries)

}

func TestSTRSAW(t *testing.T) {
	ns := []int{3, 10, 20, 200}
	retries := 20
	var fun = func(n int) *[]Vec2 {
		pos := STRSAW(n, 30)
		return pos
	}
	TestSAW(t, fun, ns, retries)
}
