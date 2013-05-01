package main

type shape interface {
	intersect(ray ray) (hit, bool)
	normal(pt vec3) vec3
}
