package main

import (
	"math/rand"
)

func CollideOrAdd(visited *map[vec2]interface{}, pos vec2) bool {
	_, collision := (*visited)[pos] // Check if steps results in a collision
	if collision {
		return true
	}
	(*visited)[pos] = nil // If there was no collision, add the position to the visited set and continue.
	return false
}
func SimpleSAW(n int) (*[]vec2, int) {
	res := make([]vec2, n+1)
	r := rand.New(rand.NewSource(rand.Int63()))
	attempts := 0
	for true {
		attempts += 1
		visited := map[vec2]interface{}{0: nil}
		res[0] = 0

		direction := r.Intn(4)
		fail := false
		for i := 0; i < n; i++ {
			direction = (direction + r.Intn(3) - 1 + 4) % 4 // Add random in interval [-1, 1] and modulo
			res[i+1] = res[i] + unitVectors[direction]      // Take a step
			fail = CollideOrAdd(&visited, res[i+1])
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

// STRSAW stands for Single Threaded Self Avoiding Walk 1
func STRSAW(n, blockSize int) *[]vec2 {
	res := make([]vec2, n+1)
	r := rand.New(rand.NewSource(rand.Int63()))

	strsawHelper(1, n+1, blockSize, &res, r)
	return &res
}

func strsawHelper(a, b, blockSize int, pos *[]vec2, r *rand.Rand) {
	if b-a <= blockSize {
		for true {
			visited := map[vec2]interface{}{0: nil}
			direction := r.Intn(4)
			fail := false
			currPos := vec2(0)
			for i := a; i < b; i++ {
				direction = (direction + r.Intn(3) - 1 + 4) % 4 // Add random in interval [-1, 1] and modulo
				currPos += unitVectors[direction]
				(*pos)[i] = currPos // Take a step
				fail = CollideOrAdd(&visited, (*pos)[i])
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
		visited := map[vec2]interface{}{0: nil}

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

			fail = CollideOrAdd(&visited, (*pos)[i])
			if fail {
				break
			}
		}
		if !fail {
			return
		}
	}
}
