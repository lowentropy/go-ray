package main

import _ "fmt"

func trace(scene *scene, ray ray, n int) color {
	if n > 4 {
		return black
	}

	hit, body, ok := scene.Intersect(ray)
	if !ok {
		//		fmt.Print(".")
		return black
	}

	//	fmt.Print("+")

	out, color, emit := body.material.bounce(ray.normal, hit.normal)
	var pt vec3

	if out.dot(ray.normal) > 0 {
		pt = ray.origin.add(ray.normal.scale(hit.dist * 1.0000001))
	} else {
		pt = ray.origin.add(ray.normal.scale(hit.dist * 0.9999999))
	}

	ray.origin = pt
	ray.normal = out
	return trace(scene, ray, n+1).mul(color).add(emit)
}
