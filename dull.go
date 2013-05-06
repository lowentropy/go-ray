package main

type dull struct {
	color color
}

func (m *dull) bounce(incoming, normal vec3) (vec3, color, color) {
	v := uniHemiSample(normal)
	// color := m.color
	color := m.color.scale(v.dot(normal))
	return v, color, black
}
