package algorithms

import . "randomWalk/saw"

func ZombiePool(n int, fun Walker, poolSize int) *[]Vec2 {
	resultChan := make(chan *[]Vec2)
	droplet := func(c chan *[]Vec2) {
		c <- fun(n)
	}
	for i := 0; i < poolSize; i++ {
		go droplet(resultChan)
	}

	return <-resultChan
}

func Pool(n int, fun ReactiveWalker, poolSize int) *[]Vec2 {
	resultChan := make(chan *[]Vec2)
	stopWalking := false
	droplet := func(c chan *[]Vec2) {
		select {
		case c <- fun(n, &stopWalking):
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
