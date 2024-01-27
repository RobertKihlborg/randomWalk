package algorithms

import (
	"fmt"
	. "randomWalk/saw"
	"testing"
)

func BenchmarkWalkers(b *testing.B) {
	n := 100
	for k, walker := range ZombieWalkers {
		b.Run(fmt.Sprintf("%v", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				walker(n)
			}
		})
	}
}

func BenchmarkWalkersBig(b *testing.B) {
	n := 1000
	for k, walker := range map[string]Walker{"mapless": ZombieWalkers["mapless"], "parallel": ZombieWalkers["parallel"]} {
		b.Run(fmt.Sprintf("%v", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				walker(n)
			}
		})
	}
}

func BenchmarkPools(b *testing.B) {
	n := 2000
	procs := 30

	b.Run("Reactive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Pool(n, ReactiveWalkers["mapless"], procs)
		}
	})

	b.Run("Zombie", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ZombiePool(n, ZombieWalkers["mapless"], procs)
		}
	})
}

func BenchmarkPooledMaplessSAOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZombiePool(100, ZombieWalkers["maplessOld"], 10)
	}
}

func BenchmarkPooledMaplessSAOldBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZombiePool(2000, ZombieWalkers["maplessOld"], 10)
	}
}

func BenchmarkPooledMaplessSA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZombiePool(100, ZombieWalkers["mapless"], 10)
	}
}

func BenchmarkPooledMaplessSABig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ZombiePool(2000, ZombieWalkers["mapless"], 10)
	}
}
