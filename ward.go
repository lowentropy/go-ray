package main

import (
	"math"
)

type ward struct {
	diffuse, specular, emit color
	roughness               float64
}

func (m *ward) bounce(incoming, normal vec3) (vec3, color, color) {
	out := incoming.reflect(normal)
	half := out.sub(incoming).norm()
	theta := math.Acos(half.dot(normal))
	tan := math.Tan(theta)
	a2 := m.roughness * m.roughness
	exp := math.Exp(-tan * tan / a2)
	rad := -normal.dot(incoming) * normal.dot(out)
	rhs := exp / (4 * a2 * math.Sqrt(rad))
	spec := m.specular.scale(rhs)
	color := m.diffuse.add(spec)
	return out, color, m.emit
}
