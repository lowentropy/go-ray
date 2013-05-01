package main

type color struct {
	r, g, b float64
}

const (
	SUN_INTENSITY = 1.0
)

var red = color{1, 0, 0}
var green = color{0, 1, 0}
var blue = color{0, 0, 1}
var black = color{0, 0, 0}
var white = color{1, 1, 1}
var sunlight = color{1.6 * SUN_INTENSITY, 1.47 * SUN_INTENSITY, 1.29 * SUN_INTENSITY}

func (a color) mul(b color) color {
	return color{a.r * b.r, a.g * b.g, a.b * b.b}
}

func (a color) add(b color) color {
	return color{a.r + b.r, a.g + b.g, a.b + b.b}
}
