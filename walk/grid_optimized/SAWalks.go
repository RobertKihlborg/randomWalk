package grid_optimized

import (
	"math/rand"
	. "randomWalk/walk"
)

func collideOrAddToMap[T comparable](visited map[T]interface{}, pos T) bool {
	_, collision := visited[pos] // Check if steps results in a collision
	if collision {
		return true
	}
	visited[pos] = nil // If there was no collision, add the position to the visited set and continue.
	return false
}

func CreateWalkerFix[T Vec2 | Vec3 | Vec4](unitVectors []T) Walker[T] {
	unitVecCount := len(unitVectors)
	return func(n int) []T {
		res := make([]T, n+1)
		r := rand.New(rand.NewSource(rand.Int63()))

		for true {
			visited := map[T]interface{}{0: nil}
			res[0] = 0
			dir := r.Intn(unitVecCount)
			fail := false

			for i := 1; i < n+1; i++ {
				dir = RandomIntExcept(unitVecCount, (dir+unitVecCount/2)%unitVecCount, r)
				res[i] = res[i-1] + unitVectors[dir]
				if fail = collideOrAddToMap(visited, res[i]); fail {
					break
				}
			}
			if !fail {
				return res
			}
		}
		panic("Escaped infinite loop")
	}
}
