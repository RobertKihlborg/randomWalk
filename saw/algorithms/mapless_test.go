package algorithms

import (
	. "randomWalk/saw"
	"testing"
)

func TestMaplessOld(t *testing.T) {
	ns := []int{3, 10, 20, 200}
	retries := 20
	var fun = func(n int) *[]Vec2 {
		pos := MaplessOld(n, 30)
		return pos
	}
	TestSAW(t, fun, ns, retries)
}

func TestZombieMapless(t *testing.T) {
	ns := []int{3, 10, 20, 200}
	retries := 20
	var fun = func(n int) *[]Vec2 {
		pos := ZombieMapless(n, 30)
		return pos
	}
	TestSAW(t, fun, ns, retries)
}

func TestMapless(t *testing.T) {
	ns := []int{3, 10, 20, 200}
	retries := 20
	var fun = func(n int) *[]Vec2 {
		pos := ZombieMapless(n, 30)
		return pos
	}
	TestSAW(t, fun, ns, retries)
}
