package main

type material interface {
	bounce(incoming, normal vec3) (vec3, color, color)
}
