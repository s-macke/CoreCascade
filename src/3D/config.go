package main

import (
	"flag"
)

type flagConfig struct {
	Width          int
	Height         int
	OutputFilename string
	Scene          string
	Time           float32
	Method         string
}

func parseConfig() (config flagConfig) {
	flag.IntVar(&config.Width, "width", 800, "Width of the output image")
	flag.IntVar(&config.Height, "height", 800, "Height of the output image")
	flag.StringVar(&config.OutputFilename, "output", "output", "Output filename for the rendered image")
	flag.StringVar(&config.Scene, "scene", "height", "Scene to render (e.g., height, fluid_height)")
	time := 0.
	flag.Float64Var(&time, "time", 0.0, "Time of the scene")
	flag.StringVar(&config.Method, "method", "radiance_cascade", "Rendering method to use :\n"+
		"  path_tracing\n"+
		"  radiance_cascade\n")
	flag.Parse()
	config.Time = float32(time)
	return
}
