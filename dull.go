package main

type dull struct {
	color color
}

func (m dull) bounce(incoming, normal vec3) (vec3, color, color) {
	v := uniHemiSample(normal)
	return v, m.color, black
}
