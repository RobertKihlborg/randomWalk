package main

import (
	"testing"
)

type sawfunc func(n int) *[]vec2

func UniquePosCount(positions *[]vec2) int {
	count := 0
	visited := make(map[vec2]interface{})
	for _, p := range *positions {
		_, alreadyVisited := visited[p]
		visited[p] = nil
		if !alreadyVisited {
			count += 1
		}
	}
	return count
}

func testSAW(t *testing.T, fun sawfunc, ns []int, retries int) {
	for _, n := range ns {
		for i := 0; i < retries; i++ {
			pos := fun(n)
			if UniquePosCount(pos) != n+1 {
				t.Fatalf("Wrong amount of steps in walk")

			}
		}
	}
}

func TestSimpleSAW(t *testing.T) {
	ns := []int{3, 10, 20, 50}
	retries := 20
	var fun = func(n int) *[]vec2 {
		pos, _ := SimpleSAW(n)
		return pos
	}
	testSAW(t, fun, ns, retries)

}

func TestSTRSAW(t *testing.T) {
	ns := []int{3, 10, 20, 200}
	retries := 20
	var fun = func(n int) *[]vec2 {
		pos := STRSAW(n, 30)
		return pos
	}
	testSAW(t, fun, ns, retries)
}

func TestMaplessSAOld(t *testing.T) {
	ns := []int{3, 10, 20, 200}
	retries := 20
	var fun = func(n int) *[]vec2 {
		pos := MaplessSAOld(n, 30)
		return pos
	}
	testSAW(t, fun, ns, retries)
}

func TestMaplessSA(t *testing.T) {
	ns := []int{3, 10, 20, 200}
	retries := 20
	var fun = func(n int) *[]vec2 {
		pos := MaplessSA(n, 30)
		return pos
	}
	testSAW(t, fun, ns, retries)
}

func TestParallelSA(t *testing.T) {
	ns := []int{3, 10, 20, 200}
	retries := 20
	var fun = func(n int) *[]vec2 {
		pos := ParallelSA(n, 30)
		return pos
	}
	testSAW(t, fun, ns, retries)
}

func BenchmarkSimpleSAW(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SimpleSAW(100)
	}
}

func BenchmarkSTRSAW(b *testing.B) {
	for i := 0; i < b.N; i++ {
		STRSAW(100, 30)
	}
}

func BenchmarkMaplessSAOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaplessSAOld(100, 30)
	}
}

func BenchmarkMaplessSAOldBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaplessSAOld(1000, 30)
	}
}

func BenchmarkMaplessSA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaplessSA(100, 30)
	}
}

func BenchmarkMaplessSABig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaplessSA(2000, 30)
	}
}

func BenchmarkParallelSA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParallelSA(100, 30)
	}
}

func BenchmarkParallelSABig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParallelSA(1000, 30)
	}
}

func BenchmarkPooledMaplessSAOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PooledSA(100, walkers["maplessOld"], 10)
	}
}

func BenchmarkPooledMaplessSAOldBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PooledSA(2000, walkers["maplessOld"], 10)
	}
}

func BenchmarkPooledMaplessSA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PooledSA(100, walkers["mapless"], 10)
	}
}

func BenchmarkPooledMaplessSABig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PooledSA(2000, walkers["mapless"], 10)
	}
}
