package main

import (
	"math"
	"math/rand"
)

type camera struct {
	pos, target, up, normal vec3
	aspect, fx, fy, fd      float64
	w, h                    int
}

func NewCamera(pos, target vec3, w, h int, fd float64) *camera {
	c := &camera{pos: pos, target: target, w: w, h: h, fd: fd}
	retarget(c)
	return c
}

func retarget(c *camera) {
	c.aspect = float64(c.w) / float64(c.h)
	c.fy = 2 * math.Atan(c.fd)
	c.fx = 2 * math.Atan(c.fd*c.aspect)
	c.up = Y
	c.normal = c.target.sub(c.pos).norm()
}

func (c *camera) Shoot(x, y int, jitter float64) ray {
	xj := (rand.Float64() - 0.5) * jitter
	yj := (rand.Float64() - 0.5) * jitter
	xp := (float64(x) + xj) / float64(c.w-1)
	yp := (float64(y) + yj) / float64(c.h-1)
	sx := math.Tan(c.fx / 2)
	sy := math.Tan(c.fy / 2)
	r := c.normal.cross(c.up).norm()
	u := r.cross(c.normal)
	dx := (xp - 0.5) * sx
	dy := (0.5 - yp) * sy
	v := c.normal.add(r.scale(dx)).add(u.scale(dy))
	return ray{c.pos, v.norm()}
}
