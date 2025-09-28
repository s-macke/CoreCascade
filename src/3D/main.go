package main

import (
	"CoreCascade3D/render/path_tracing"
	"CoreCascade3D/scene"
	"CoreCascade3D/scene/scenes"
	"fmt"
	"linear_image"
	"strings"
)

func NewScene(sceneAsString string, time float64) scene.Scene {
	sceneAsString = strings.ToLower(sceneAsString)
	var s scene.Scene = nil
	switch sceneAsString {
	case "height":
		s = scenes.NewSceneHeight(time)
	case "fluid_height":
		s = scenes.NewSceneFluidHeight(time)
	default:
		panic("Unknown scene: " + sceneAsString)
	}
	return s
}

func main() {
	config := parseConfig()
	fmt.Printf("Scene: %s\n", config.Scene)
	fmt.Printf("Size: %dx%d\n", config.Width, config.Height)
	fmt.Printf("Method: %s\n", config.Method)
	sc := NewScene(config.Scene, config.Time)

	image := linear_image.NewSampledImage(config.Height, config.Width)

	switch config.Method {
	case "path_tracing":
		path_tracing.NewPathTracing3D(sc, image).Render(1000)
	case "error":
	default:
		panic("Unknown method")
	}
	image.Store(config.OutputFilename)
}
