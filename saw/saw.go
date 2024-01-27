package saw

import (
	"fmt"
)

type Vec2 int32

func (v Vec2) xy() (int16, int16) {
	y := int16(v % (1 << 16))
	x := int16((int32(v) - int32(y)) / (1 << 16))
	return x, y
}
func (v Vec2) String() string {
	x, y := v.xy()
	return fmt.Sprintf("(%v, %v)", x, y)
}

const right = Vec2(1 << 16)
const left = -right
const up = Vec2(1)
const down = -up

var UnitVectors = []Vec2{right, up, left, down}

type Walker func(n int) *[]Vec2

type ReactiveWalker func(n int, stopWalking *bool) *[]Vec2
