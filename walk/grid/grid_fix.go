package grid

import (
	"fmt"
)

const intSize = 1 << 16

type Vec2 int32
type Vec3 int64
type Vec4 int64

func (v Vec2) xy() (int16, int16) {
	x := int16(v % intSize)
	y := int16((int32(v) - int32(x)) / intSize)
	return x, y
}
func (v Vec2) String() string {
	x, y := v.xy()
	return fmt.Sprintf("(%v, %v)", x, y)
}

func (v Vec3) xyz() (int16, int16, int16) {

	x := int16(v % intSize)
	y := int16((int64(v) - int64(x)) / intSize % intSize)
	z := int16((int64(v) - int64(y) - int64(x)) / (intSize * intSize))
	return x, y, z
}
func (v Vec3) String() string {
	x, y, z := v.xyz()
	return fmt.Sprintf("(%v, %v, %v)", x, y, z)
}

func (v Vec4) xyzw() (int16, int16, int16, int16) {

	x := int16(v % intSize)
	y := int16((int64(v) - int64(x)) / intSize % intSize)
	z := int16((int64(v) - int64(y) - int64(x)) / (intSize * intSize))
	w := int16((int64(v) - -int64(z) - int64(y) - int64(x)) / (intSize * intSize * intSize))

	return x, y, z, w
}
func (v Vec4) String() string {
	x, y, z, w := v.xyzw()
	return fmt.Sprintf("(%v, %v, %v, %v)", x, y, z, w)
}

const right = 1
const left = -right
const up = 1 << 16
const down = -up
const front = 1 << 32
const back = -front
const inwards = 1 << 48
const outwards = -inwards

var UnitVectors2D = []Vec2{right, up, left, down}
var UnitVectors3D = []Vec3{right, up, left, down, front, back}
var UnitVectors4D = []Vec4{right, up, left, down, front, back, inwards, outwards}

func IsPointIntersectingFix[T Vec2 | Vec3 | Vec4](walk []T) bool {
	visited := make(map[T]interface{})
	for _, p := range walk {
		_, alreadyVisited := visited[p]
		visited[p] = nil
		if alreadyVisited {
			return true
		}
	}
	return false
}
