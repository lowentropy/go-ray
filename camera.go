package main

import (
	"math"
	"math/rand"
)

type camera struct {
	pos, target, up, normal vec3
	aspect, fx, fy          float64
	w, h                    int
}

func NewCamera(pos, target vec3, w, h int) *camera {
	c := &camera{pos: pos, target: target, w: w, h: h}
	retarget(c)
	return c
}

func retarget(c *camera) {
	f := 1.5
	c.aspect = float64(c.w) / float64(c.h)
	c.fy = 2 * math.Atan(f)
	c.fx = 2 * math.Atan(f*c.aspect)
	c.up = Y
	c.normal = c.target.sub(c.pos).norm()
}

func (c *camera) Shoot(x, y int) ray {
	xp := (float64(x) + rand.Float64() - 0.5) / float64(c.w-1)
	yp := (float64(y) + rand.Float64() - 0.5) / float64(c.h-1)
	sx := math.Tan(c.fx / 2)
	sy := math.Tan(c.fy / 2)
	r := c.normal.cross(c.up).norm()
	u := r.cross(c.normal)
	dx := (xp - 0.5) * sx
	dy := (0.5 - yp) * sy
	v := c.normal.add(r.scale(dx)).add(u.scale(dy))
	return ray{c.pos, v.norm()}
}
