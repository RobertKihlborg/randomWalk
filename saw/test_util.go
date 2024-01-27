package saw

import "testing"

func uniquePosCount(positions *[]Vec2) int {
	count := 0
	visited := make(map[Vec2]interface{})
	for _, p := range *positions {
		_, alreadyVisited := visited[p]
		visited[p] = nil
		if !alreadyVisited {
			count += 1
		}
	}
	return count
}

func TestSAW(t *testing.T, fun Walker, ns []int, retries int) {
	for _, n := range ns {
		for i := 0; i < retries; i++ {
			pos := fun(n)
			if uniquePosCount(pos) != n+1 {
				t.Fatalf("Wrong amount of steps in walk")

			}
		}
	}
}
