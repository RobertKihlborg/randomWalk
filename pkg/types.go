package pkg

import (
	"fmt"
	"math/rand"
)

type Walk interface {
	fmt.Stringer
	Distance() float64
}

type Walker func(n int) Walk

type WalkerBuilder func(args []string) (walker Walker, name string, usedArgs int, err error)

type PoolableWalker func(n int, r *rand.Rand, interrupt *bool) Walk
