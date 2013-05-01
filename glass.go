package main

import (
	"math/rand"
)

type glass struct {
	color       color
	index       float64
	reflectance float64
}

func (m glass) bounce(incoming, normal vec3) (out vec3, c color, e color) {
	var i1, i2 float64
	t1 := incoming.dot(normal)
	c = m.color
	e = black
	if t1 >= 0 {
		i1, i2 = m.index, 1.0
	} else {
		i1, i2 = 1.0, m.index
	}
	r := i1 / i2
	t2 := 1 - r*r*(1-t1*t1)
	rs := (i1*t1 - i2*t2) / (i1*t1 + i2*t2)
	rp := (i2*t1 - i1*t2) / (i2*t1 + i1*t2)
	reflectance := m.reflectance + rs*rs + rp*rp
	if rand.Float64() < reflectance {
		out = incoming.add(normal.scale(t1 * 2))
	} else {
		out = incoming.add(normal.scale(t1)).scale(r).add(normal.scale(-t2))
	}
	return
}
