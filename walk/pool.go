package walk

import "math/rand"

func pool[T any](n int, fun PoolableWalker[T], poolSize int) []T {
	resultChan := make(chan []T)
	stopWalking := false
	r := rand.New(rand.NewSource(rand.Int63()))
	droplet := func(c chan []T) {
		select {
		case c <- fun(n, r, &stopWalking):
		default:
		}
	}

	for i := 0; i < poolSize; i++ {
		go droplet(resultChan)
	}

	result := <-resultChan
	stopWalking = true
	return result

}

func CreatePool[T any](fun PoolableWalker[T], poolSize int) Walker[T] {
	return func(n int) []T {
		return pool[T](n, fun, poolSize)
	}
}
