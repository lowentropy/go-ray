package main

type chrome struct {
	color color
}

func (m *chrome) bounce(incoming, normal vec3) (vec3, color, color) {
	return incoming.reflect(normal), m.color, black
}
