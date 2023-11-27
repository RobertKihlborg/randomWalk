package main

func PooledSA(n int, fun Walker, poolSize int) *[]vec2 {
	resultChan := make(chan *[]vec2)
	droplet := func(c chan *[]vec2) {
		c <- fun(n)
	}
	for i := 0; i < poolSize; i++ {
		go droplet(resultChan)
	}

	return <-resultChan
}
