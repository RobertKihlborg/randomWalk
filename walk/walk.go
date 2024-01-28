package walk

import (
	"math/rand"
)

type Walker[T any] func(n int) []T

type PoolableWalker[T any] func(n int, r *rand.Rand, interrupt *bool) []T

type CollisionChecker[T any] func(walk []T) bool

type RandomFunction[T any] func(*rand.Rand) T

type NonreturningRandomFunction[T any] func(prev T, r *rand.Rand) T
