package main

import (
	"math"
	"math/rand"
)

type glass struct {
	color       color
	index       float64
	reflectance float64
}

func (m *glass) bounce(incoming, normal vec3) (out vec3, c color, e color) {
	c = m.color
	e = black
	c1 := -normal.dot(incoming)
	var n1, n2 float64
	if c1 >= 0 {
		n1, n2 = 1.0, m.index
	} else {
		n1, n2 = m.index, 1.0
	}
	r := n1 / n2
	s := 1 - r*r*(1-c1*c1)
	if s < 0 {
		return incoming, black, black
	}
	c2 := math.Sqrt(s)
	if c1 < 0 {
		c2 = -c2
	}
	rp := (n2*c1 - n1*c2) / (n2*c1 + n1*c2)
	rs := (n1*c1 - n2*c2) / (n1*c1 + n2*c2)
	reflectance := (rs*rs + rp*rp) / 2
	if rand.Float64() < reflectance {
		out = incoming.reflect(normal)
	} else {
		out = incoming.scale(r).add(normal.scale(r*c1 - c2)).norm()
	}
	return
}
