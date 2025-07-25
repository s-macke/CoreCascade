package main

import (
	"flag"
)

type flagConfig struct {
	Width          int
	Height         int
	OutputFilename string
	InputFilename  string
	Scene          string
	Method         string
}

func parseConfig() (config flagConfig) {
	flag.IntVar(&config.Width, "width", 800, "Width of the output image")
	flag.IntVar(&config.Height, "height", 800, "Height of the output image")
	flag.StringVar(&config.OutputFilename, "output", "output", "Output filename for the rendered image")
	flag.StringVar(&config.InputFilename, "input", "", "Input raw file")
	flag.StringVar(&config.Scene, "scene", "shadows", "Scene to render (e.g., center, pinhole, penumbra, shadows, beam)")
	flag.StringVar(&config.Method, "method", "vanilla_radiance_cascade", "Rendering method to use (e.g., path_tracing, path_tracing_parallel, vanilla_radiance_cascade)")
	flag.Parse()
	return
}
