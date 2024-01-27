package algorithms

import (
	"math/rand"
	. "randomWalk/saw"
)

func CollideOrAddToMap(visited *map[Vec2]interface{}, pos Vec2) bool {
	_, collision := (*visited)[pos] // Check if steps results in a collision
	if collision {
		return true
	}
	(*visited)[pos] = nil // If there was no collision, add the position to the visited set and continue.
	return false
}

func SimpleSAW(n int) (*[]Vec2, int) {
	res := make([]Vec2, n+1)
	r := rand.New(rand.NewSource(rand.Int63()))
	attempts := 0
	for true {
		attempts += 1
		visited := map[Vec2]interface{}{0: nil}
		res[0] = 0

		direction := r.Intn(4)
		fail := false
		for i := 0; i < n; i++ {
			direction = (direction + r.Intn(3) - 1 + 4) % 4 // Add random in interval [-1, 1] and modulo
			res[i+1] = res[i] + UnitVectors[direction]      // Take a step
			fail = CollideOrAddToMap(&visited, res[i+1])
			if fail {
				break
			}
		}
		if !fail {
			break
		}
	}
	return &res, attempts
}

// STRSAW stands for Single Threaded Self Avoiding Walk
func STRSAW(n, blockSize int) *[]Vec2 {
	res := make([]Vec2, n+1)
	r := rand.New(rand.NewSource(rand.Int63()))

	strsawHelper(1, n+1, blockSize, &res, r)
	return &res
}

func strsawHelper(a, b, blockSize int, pos *[]Vec2, r *rand.Rand) {
	if b-a <= blockSize {
		for true {
			visited := map[Vec2]interface{}{0: nil}
			direction := r.Intn(4)
			fail := false
			currPos := Vec2(0)
			for i := a; i < b; i++ {
				direction = (direction + r.Intn(3) - 1 + 4) % 4 // Add random in interval [-1, 1] and modulo
				currPos += UnitVectors[direction]
				(*pos)[i] = currPos // Take a step
				fail = CollideOrAddToMap(&visited, (*pos)[i])
				if fail {
					break
				}
			}
			if !fail {
				return
			}
		}
	}

	// Case for when to do recursive call
	mid := (a + b) / 2

	for true {
		fail := false
		visited := map[Vec2]interface{}{0: nil}

		strsawHelper(a, mid, blockSize, pos, r)
		strsawHelper(mid, b, blockSize, pos, r)

		// Note which nodes were visited in the first part (return this directly maybe?)
		for i := a; i < mid; i++ {
			visited[(*pos)[i]] = nil
		}
		lastPosOfFirst := (*pos)[mid-1]

		// Check which nodes were visited in the second part if we put it at the end of the first, and check for collisions
		for i := mid; i < b; i++ {
			(*pos)[i] += lastPosOfFirst

			fail = CollideOrAddToMap(&visited, (*pos)[i])
			if fail {
				break
			}
		}
		if !fail {
			return
		}
	}
}
