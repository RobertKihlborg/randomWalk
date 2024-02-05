package grid_optimized

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkRandomIntExcepts(b *testing.B) {
	except := 0
	r := rand.New(rand.NewSource(rand.Int63()))
	for _, n := range []int{4, 6, 8, 10} {
		b.Run(fmt.Sprintf("n=%v, Alt 1", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				randomIntExcept1(n, except, r)
			}
		})
		b.Run(fmt.Sprintf("n=%v, Alt 2", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				randomIntExcept2(n, except, r)
			}
		})

		b.Run(fmt.Sprintf("n=%v, Best?", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				RandomIntExcept(n, except, r)
			}
		})
	}
}
