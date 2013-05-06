package main

import "math"

type sphere struct {
	center vec3
	radius float64
}

func (s *sphere) intersect(ray ray) (hit, bool) {
	o := ray.origin.sub(s.center)
	ll := o.dot(o)
	rr := s.radius * s.radius
	i := ll < rr
	tca := -o.dot(ray.normal)
	if !i && fneg(tca) {
		return hit{}, false
	}
	t2hc := rr - ll + tca*tca
	if !i && !fpos(t2hc) {
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
