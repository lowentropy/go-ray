package main

import (
	"math"
)

type body struct {
	shape    shape
	material material
}

type scene struct {
	camera *camera
	bodies []*body
}

type result struct {
	color color
	x, y  int
}

func NewScene() *scene {
	scene := new(scene)
	scene.bodies = make([]*body, 0)
	return scene
}

func (scene *scene) Add(body *body) {
	scene.bodies = append(scene.bodies, body)
}

func (scene *scene) Intersect(ray ray) (closest hit, found *body, any bool) {
	min := math.Inf(1)
	for _, body := range scene.bodies {
		hit, ok := body.shape.intersect(ray)
		if ok && hit.dist < min {
			min = hit.dist
			closest = hit
			found = body
			any = true
		}
	}
	return
}

func (scene *scene) Trace(x, y int) color {
	jitter := 1.0
	ray := scene.camera.Shoot(x, y, jitter)
	return trace(scene, ray, 0)
}
