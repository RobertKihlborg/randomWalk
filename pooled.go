package main

func PooledMaplessSA(n, blockSize, poolSize int) *[]vec2 {
	resultChan := make(chan *[]vec2)
	fun := func(c chan *[]vec2) {
		c <- MaplessSA(n, blockSize)
	}

	for i := 0; i < poolSize; i++ {
		go fun(resultChan)
	}

	return <-resultChan
}
