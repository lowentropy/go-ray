package main

import (
	"flag"
	"fmt"
	"image"
	c "image/color"
	"image/png"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

type shiny struct {
	window        [][]c.Color
	width, height int
}

func (s shiny) ColorModel() c.Model {
	return c.RGBAModel
}

func (s shiny) Bounds() image.Rectangle {
	return image.Rect(0, 0, s.width-1, s.height-1)
}

func (s shiny) At(x, y int) c.Color {
	return s.window[y][x]
}

func cornell(scene *scene, w, h int) {
	scene.camera = NewCamera(vec3{0, 0, 0.5}, vec3{0, 0, 0}, w, h)
	// glass sphere
	scene.Add(&body{sphere{vec3{1, 0, -2}, 0.5}, glass{white, 1.5, 0.1}})
	// chrome sphere
	scene.Add(&body{sphere{vec3{-1.1, 0, -2.8}, 0.5}, chrome{color{0.8, 0.8, 0.8}}})
	// floor
	scene.Add(&body{sphere{vec3{0, -10e6, -3.5}, 10e6 - 0.5}, dull{color{0.9, 0.9, 0.9}}})
	// back
	scene.Add(&body{sphere{vec3{0, 0, -10e6}, 10e6 - 4.5}, dull{color{0.9, 0.9, 0.9}}})
	// left
	scene.Add(&body{sphere{vec3{-10e6, 0, -3.5}, 10e6 - 1.9}, dull{color{0.9, 0.5, 0.5}}})
	// right
	scene.Add(&body{sphere{vec3{10e6, 0, -3.5}, 10e6 - 1.9}, dull{color{0.5, 0.5, 0.9}}})
	// top light
	scene.Add(&body{sphere{vec3{0, 10e6, 0}, 10e6 - 2.5}, glow{sunlight}})
	// front
	scene.Add(&body{sphere{vec3{0, 0, 10e6}, 10e6 - 2.5}, dull{color{0.9, 0.9, 0.9}}})
}

func glowSphere(scene *scene, w, h int) {
	scene.camera = NewCamera(vec3{0, 0, 5}, V0, w, h)
	// scene.Add(&body{sphere{V0, 2}, glow{sunlight}})
	scene.Add(&body{sphere{vec3{0, 10e6 - 2.5, 0}, 10e6 - 2.5}, glow{sunlight}})
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	w, h := 200, 200

	// cpus := runtime.NumCPU()
	cpus := 2
	runtime.GOMAXPROCS(cpus)
	fmt.Println("Using CPUS:", cpus)

	scene := NewScene()

	// glowSphere(scene, w, h)
	cornell(scene, w, h)

	window := scene.Render()
	img := shiny{window, w, h}

	f, err := os.Create("/Users/lowentropy/desktop/shiny.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	png.Encode(f, img)
}
