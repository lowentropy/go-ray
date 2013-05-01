package main

import (
	_ "fmt"
	c "image/color"
	"math"
	"runtime"
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
	ray := scene.camera.Shoot(x, y)
	return trace(scene, ray, 0)
}

func (scene *scene) Render() [][]c.Color {
	w, h := scene.camera.w, scene.camera.h
	n := 10
	cpus := runtime.GOMAXPROCS(0)
	ch := make(chan bool, cpus)
	step := h/cpus + 1

	buffer := make([][]color, h)
	window := make([][]c.Color, h)

	for y := 0; y < h; y++ {
		window[y] = make([]c.Color, w)
		buffer[y] = make([]color, w)
	}

	for y0 := 0; y0 < h; y0 += step {
		go func(y0 int) {
			for y := y0; y < y0+step && y < h; y++ {
				for x := 0; x < w; x++ {
					for i := 0; i < n; i++ {
						light := scene.Trace(x, y)
						buffer[y][x] = buffer[y][x].add(light)
					}
				}
			}
			ch <- true
		}(y0)
	}

	for i := 0; i < cpus; i++ {
		<-ch
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			rc := buffer[y][x]
			r := uint8(math.Min(rc.r*255/float64(n), 255))
			g := uint8(math.Min(rc.g*255/float64(n), 255))
			b := uint8(math.Min(rc.b*255/float64(n), 255))
			window[y][x] = c.RGBA{r, g, b, 255}
		}
	}

	return window
}
