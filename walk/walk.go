package walk

import (
	"math/rand"
)

type Walker[T any] func(n int) []Vector[T]

type PoolableWalker[T any] func(n int, r *rand.Rand, interrupt *bool) []Vector[T]

type CollisionChecker[T any] func(walk []Vector[T]) bool

type LatestCollisionChecker[T any] func(point Vector[T], walk []Vector[T]) bool

type RandomFunction[T any] func(*rand.Rand) Vector[T]

type NonreturningRandomFunction[T any] func(prev Vector[T], r *rand.Rand) Vector[T]

type Vector[T any] interface {
	Get(i int) T
	Size() int
	Add(Vector[T]) Vector[T]
}

type GridWalk[T int] []Vector[T]

type SpaceWalk []Vector[float64]
