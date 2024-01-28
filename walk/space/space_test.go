package space

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func testVectorRandomizer(t *testing.T, fun func(int, *rand.Rand) VectorFloat, maxDim int, reps int) {
	r := rand.New(rand.NewSource(rand.Int63()))
	for dim := 1; dim < maxDim; dim++ {
		res := make([]VectorFloat, reps)
		var maxAvgDev float64 = 0
		for i := 0; i < reps; i++ {
			res[i] = fun(dim, r)
		}
		avg := make([]float64, dim)
		for i := 0; i < dim; i++ {
			for _, re := range res {
				avg[i] += re[i]
			}
			avg[i] /= float64(reps)
			maxAvgDev = max(maxAvgDev, math.Abs(avg[i]))
		}

		if maxAvgDev > 0.1 {
			t.Fatalf("Failed at %v dimensions with avg %v", dim, avg)
		}
		fmt.Printf("dim %v maxAvgDev %v \n", dim, maxAvgDev)
	}
}

func TestRandomDirectionAngles(t *testing.T) {
	testVectorRandomizer(t, randomDirectionAngles, 10, 10000)
}

func TestRandomDirectionRS(t *testing.T) {
	testVectorRandomizer(t, randomDirectionRS, 10, 10000)
}

func BenchmarkRandomizers(b *testing.B) {
	r := rand.New(rand.NewSource(rand.Int63()))
	for _, dim := range []int{2, 6, 7, 10, 15} {
		b.Run(fmt.Sprintf("Angles %vD", dim), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				randomDirectionAngles(dim, r)
			}
		})
		b.Run(fmt.Sprintf("RS %vD", dim), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				randomDirectionRS(dim, r)
			}
		})
	}
}
