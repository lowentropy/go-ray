package main

import (
	"fmt"
	"image"
	c "image/color"
	"image/png"
	"math"
	"os"
	"runtime"
)

type rendering struct {
	cpus, iterations int
	jitter           float64
	width, height    int
	scene            *scene
}

type shiny struct {
	window        [][]c.Color
	width, height int
}

func (s *shiny) ColorModel() c.Model {
	return c.RGBAModel
}

func (s *shiny) Bounds() image.Rectangle {
	return image.Rect(0, 0, s.width-1, s.height-1)
}

func (s *shiny) At(x, y int) c.Color {
	return s.window[y][x]
}

func (r *rendering) setCpus() int {
	var cpus int
	if r.cpus > 0 {
		cpus = r.cpus
	} else {
		cpus = runtime.NumCPU() + r.cpus
	}
	runtime.GOMAXPROCS(cpus)
	return cpus
}

func (r *rendering) Render(filename string) {
	scene := r.scene
	w, h := scene.camera.w, scene.camera.h
	n := r.iterations

	cpus := r.setCpus()
	ch := make(chan bool, cpus)
	step := h/cpus + 1

	fmt.Println("CPUS:", cpus)

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

	img := &shiny{window, w, h}

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	png.Encode(f, img)
}
