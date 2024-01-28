package grid

import (
	"math/rand"
	. "randomWalk/walk"
)

func CreateNaiveWalkerFix[T Vec2 | Vec3 | Vec4](unitVectors []T) Walker[T] {
	walker := CreateSIWalkerFix(unitVectors)
	return func(n int) []T {
		for true {
			res := walker(n)
			if !IsPointIntersectingFix(res) {
				return res
			}
		}
		panic("Escaped infinite loop")
	}
}
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

func CreateNaiveWalker(dim int) Walker[VectorInt] {
	walker := CreateSIWalker(dim)
	return func(n int) []VectorInt {
		for true {
			res := walker(n)
			if !IsPointIntersecting(res) {
				return res
			}
		}
		panic("Escaped infinite loop")
	}
}

func CreateWalker(dim int) Walker[VectorInt] {
	randDirFun := CreateRandomDirectionFunctionDim(dim)
	return func(n int) []VectorInt {
		res := make([]VectorInt, n+1)
		r := rand.New(rand.NewSource(rand.Int63()))
		zero := make(VectorInt, dim)
		for true {
			visited := map[string]interface{}{toKey(zero): nil}
			res[0] = zero
			fail := false
			dirVec := randDirFun(r)
			for i := 1; i < n+1; i++ {
				dirVec = NonreturningRandomDirection(dirVec, r)
				res[i] = add(res[i-1], dirVec)
				if fail = collideOrAddToMap(visited, toKey(res[i])); fail {
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
