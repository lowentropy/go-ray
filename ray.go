package main

type ray struct {
	origin, normal vec3
}

type hit struct {
	ray        *ray
	dist       float64
	pt, normal vec3
}

func newHit(s shape, ray ray, t float64) hit {
	pt := ray.normal.scale(t).add(ray.origin)
	normal := s.normal(pt)
	return hit{&ray, t, pt, normal}
}
