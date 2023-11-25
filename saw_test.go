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
				t.Log(pos)
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

func TestMaplessSA(t *testing.T) {
	ns := []int{3, 10, 20, 200}
	retries := 20
	var fun = func(n int) *[]vec2 {
		pos := MaplessSA(n, 30)
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
		STRSAW(300, 30)
	}
}

func BenchmarkMaplessSA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaplessSA(300, 30)
	}
}

func BenchmarkMaplessSABig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaplessSA(2000, 30)
	}
}
