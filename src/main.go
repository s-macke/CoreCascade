package main

import (
	"CoreCascade/primitives"
	"CoreCascade/render/path_tracing"
	"CoreCascade/render/radiance_cascade"
	"CoreCascade/scene"
	"fmt"
	"strings"
)

func NewScene(sceneAsString string, time float64) *scene.Scene {
	sceneAsString = strings.ToLower(sceneAsString)
	switch sceneAsString {
	case "center":
		return scene.NewSceneCenter()
	case "shadows":
		return scene.NewSceneShadows(time)
	case "pinhole":
		return scene.NewScenePinhole()
	case "penumbra":
		return scene.NewScenePenumbra()
	case "beam":
		return scene.NewSceneBeam()
	}
	panic("Unknown scene")
}

func main() {
	config := parseConfig()
	fmt.Printf("Scene: %s\n", config.Scene)
	fmt.Printf("Size: %dx%d\n", config.Width, config.Height)
	sc := NewScene(config.Scene, config.Time)

	var image *primitives.SampledImage
	if config.InputFilename != "" {
		image = primitives.NewSampledImageFromFile(config.InputFilename)
	} else {
		image = primitives.NewSampledImage(config.Height, config.Width)
	}

	switch config.Method {
	case "path_tracing":
		path_tracing.RenderPathTracing(sc, image)
		image.Store(config.OutputFilename)
	case "path_tracing_parallel":
		path_tracing.RenderPathTracingParallel(sc, image, 2)
		image.Store(config.OutputFilename)
	case "vanilla_radiance_cascade":
		radiance_cascade.NewRadianceCascade(sc, image, false).Render()
		image.Store(config.OutputFilename)
	case "bilinear_fix_radiance_cascade":
		radiance_cascade.NewRadianceCascade(sc, image, true).Render()
		image.Store(config.OutputFilename)
	case "error":
		//truth := NewSampledImageFromFile("ring_shadow.raw")
		//truth.Error(img)
		//truth.StorePNG("diff.png")
	case "plot":
		PlotCascade()
		PlotProbeCenter()
		PlotProbeCascadesNonSpatial()
		/*
			PlotSignedDistance()
			PlotCascade2()
			PlotCascade3()
			PlotCascade4()
			PlotCascade5()
			PlotEnergy(img)
		*/

	default:
		panic("Unknown method")
	}
}
