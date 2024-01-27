package algorithms

import (
	"math/rand"
	. "randomWalk/saw"
)

type searcher interface {
	search(signal chan bool)
}

type searchNode struct {
	a, b, mid    int
	list, sorted *[]Vec2
	r            *rand.Rand
	left, right  searcher
	chl, chr     chan bool
}

func (s searchNode) search(signal chan bool) {
	go s.left.search(s.chl)
	go s.right.search(s.chr)
	for true {
		<-signal
		for true {
			fail := false
			s.chl <- true
			s.chr <- true
			<-s.chl
			<-s.chr

			// Note which nodes were visited in the first part (return this directly maybe?)
			lastPosOfFirst := (*s.list)[s.mid-1]

			// Check which nodes were visited in the second part if we put it at the end of the first, and check for collisions
			for i := s.mid; i < s.b; i++ {
				(*s.list)[i] += lastPosOfFirst
				fail = CollideOrAddSort(s.a, i, (*s.sorted)[i]+lastPosOfFirst, s.sorted)
				if fail {
					break
				}
			}
			if !fail {
				signal <- true
				break
			}
		}
	}

}

type searchLeaf struct {
	a, b         int
	list, sorted *[]Vec2
	r            *rand.Rand
}

func (s searchLeaf) search(signal chan bool) {
	for true {
		<-signal
		for true {
			direction := s.r.Intn(4)
			fail := false
			currPos := Vec2(0)
			for i := s.a; i < s.b; i++ {
				direction = (direction + s.r.Intn(3) - 1 + 4) % 4 // Add random in interval [-1, 1] and modulo
				currPos += UnitVectors[direction]
				(*s.list)[i] = currPos // Take a step
				fail = CollideOrAddSort(s.a, i, currPos, s.sorted)
				if fail {
					break
				}
			}
			if !fail {
				signal <- true
				break
			}
		}
	}

}

func newSearcher(a, b, blockSize int, list, sorted *[]Vec2, r *rand.Rand) searcher {
	if b-a <= blockSize {
		return searcher(&searchLeaf{a, b, list, sorted, r})
	}

	mid := (a + b) / 2
	leftChild := newSearcher(a, mid, blockSize, list, sorted, rand.New(rand.NewSource(r.Int63())))
	rightChild := newSearcher(mid, b, blockSize, list, sorted, rand.New(rand.NewSource(r.Int63())))

	return searcher(&searchNode{a, b, (a + b) / 2, list, sorted, r, leftChild, rightChild, make(chan bool), make(chan bool)})
}

func Parallel(n, blockSize int) *[]Vec2 {
	res := make([]Vec2, n+1)
	sorted := make([]Vec2, n+1)
	r := rand.New(rand.NewSource(rand.Int63()))
	s := newSearcher(1, n+1, blockSize, &res, &sorted, r)
	c := make(chan bool)
	go s.search(c)
	c <- true
	<-c
	return &res
}
