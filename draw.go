package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var width = flag.Int("w", 300, "image width")
var height = flag.Int("h", 300, "image height")
var iter = flag.Int("i", 20, "rendering iterations")
var cpus = flag.Int("c", 0, "CPUs to use")
var jit = flag.Float64("j", 1.0, "Pixel jitter")

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

	scene := cornell(*width, *height)
	rend := rendering{*cpus, *iter, *jit, *width, *height, scene}
	rend.Render("/Users/lowentropy/desktop/shiny.png")
}
