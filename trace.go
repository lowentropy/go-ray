package main

import _ "fmt"

func trace(scene *scene, r ray, n int) color {
	if n > 6 {
		return black
	}

	hit, body, ok := scene.Intersect(r)
	if !ok {
		return black
	}

	out, color, emit := body.material.bounce(r.normal, hit.normal)

	newray := ray{hit.pt, out}
	liftRay(&newray)

	// color = color.scale(1 / hit.dist)

	return trace(scene, newray, n+1).mul(color).add(emit)
}
