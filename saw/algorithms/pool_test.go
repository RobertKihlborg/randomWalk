package algorithms

import (
	. "randomWalk/saw"
	"testing"
)

func TestZombiePool(t *testing.T) {
	ns := []int{3, 10, 20, 200}
	retries := 20
	var fun = func(n int) *[]Vec2 {
		pos := ZombiePool(n, ZombieWalkers["mapless"], 10)
		return pos
	}
	TestSAW(t, fun, ns, retries)
}

func TestPool(t *testing.T) {
	ns := []int{3, 10, 20, 200}
	retries := 20
	var fun = func(n int) *[]Vec2 {
		pos := Pool(n, ReactiveWalkers["mapless"], 10)
		return pos
	}
	TestSAW(t, fun, ns, retries)
}
