package main

import (
	"math"
	"math/rand"
)

type vec3 struct {
	x, y, z float64
}

const (
	FTOL = 0.00000001
)

var V0 = vec3{0, 0, 0}
var V1 = vec3{1, 1, 1}
var X = vec3{1, 0, 0}
var Y = vec3{0, 1, 0}
var Z = vec3{0, 0, 1}

func fabs(f float64) float64 {
	if f < 0 {
		return -f
	}
	return f
}

func fzero(f float64) bool {
	return fabs(f) < FTOL
}

func fneg(f float64) bool {
	return !(f > 0 || fzero(f))
}

func fpos(f float64) bool {
	return !(f < 0 || fzero(f))
}

func (a vec3) add(b vec3) vec3 {
	return vec3{a.x + b.x, a.y + b.y, a.z + b.z}
}

func (a vec3) sub(b vec3) vec3 {
	return vec3{a.x - b.x, a.y - b.y, a.z - b.z}
}

func (a *vec3) inc(b vec3) {
	a.x += b.x
	a.y += b.y
	a.z += b.z
}

func (a *vec3) dec(b vec3) {
	a.x -= b.x
	a.y -= b.y
	a.z -= b.z
}

func (a vec3) dot(b vec3) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func (a vec3) cross(b vec3) vec3 {
	return vec3{
		a.y*b.z - a.z*b.y,
		a.z*b.x - a.x*b.z,
		a.x*b.y - a.y*b.x,
	}
}

func (a vec3) scale(s float64) vec3 {
	return vec3{a.x * s, a.y * s, a.z * s}
}

func (a vec3) mag() float64 {
	return math.Sqrt(a.dot(a))
}

func (a vec3) norm() vec3 {
	m := a.mag()
	if m < FTOL {
		return V0
	}
	return a.scale(1.0 / m)
}

func (a vec3) reflect(n vec3) vec3 {
	return a.sub(n.scale(a.dot(n) * 2.0))
}

func (a vec3) eq(b vec3) bool {
	return (math.Abs(a.x-b.x) < FTOL) &&
		(math.Abs(a.y-b.y) < FTOL) &&
		(math.Abs(a.z-b.z) < FTOL)
}

func randVec3() vec3 {
	return vec3{
		rand.Float64(),
		rand.Float64(),
		rand.Float64(),
	}
}

func uniHemiSample(normal vec3) (v vec3) {
	for {
		v = randVec3().scale(2).sub(V1)
		if v.dot(v) > 1 {
			break
		}
	}

	v = v.norm()
	if v.dot(normal) < 0 {
		v = v.scale(-1)
	}

	return
}

func cosHemiSample() vec3 {
	u1, u2 := rand.Float64(), rand.Float64()
	r := math.Sqrt(u1)
	theta := 2.0 * math.Pi * u2
	x := r * math.Cos(theta)
	y := r * math.Sin(theta)
	z := math.Sqrt(1.0 - u1)
	return vec3{x, y, z}
}
