package main

type glow struct {
	color color
}

func (m *glow) bounce(incoming, normal vec3) (vec3, color, color) {
	return incoming, black, m.color
}
