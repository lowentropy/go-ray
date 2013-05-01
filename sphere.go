package main

import "math"

type sphere struct {
	center vec3
	radius float64
}

func (s sphere) intersect(ray ray) (hit, bool) {
	d := ray.origin.sub(s.center)
	dist2 := d.dot(d)
	rad2 := s.radius * s.radius
	inside := dist2 < rad2
	tca := -d.dot(ray.normal)
	if !inside && tca < 0 {
		return hit{}, false
	}
	t2hc := rad2 - dist2 + (tca * tca)
	if !inside && t2hc <= 0 {
		return hit{}, false
	}
	thc := math.Sqrt(t2hc)
	t1 := tca - thc
	t2 := tca + thc
	var t float64
	if t1 < 0 {
		t = t2
	} else {
		t = t1
	}
	return newHit(s, ray, t), true
}

func (s sphere) normal(pt vec3) vec3 {
	return pt.sub(s.center).norm()
}
