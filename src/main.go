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
	case "title":
		return scene.NewSceneTitle(time)
	}
	panic("Unknown scene")
}

func main() {
	config := parseConfig()
	fmt.Printf("Scene: %s\n", config.Scene)
	fmt.Printf("Size: %dx%d\n", config.Width, config.Height)
	fmt.Printf("Method: %s\n", config.Method)
	sc := NewScene(config.Scene, config.Time)

	var image *primitives.SampledImage
	if config.InputFilename != "" {
		fmt.Println("Reading image from file", config.InputFilename)
		image = primitives.NewSampledImageFromFile(config.InputFilename)
	} else {
		image = primitives.NewSampledImage(config.Height, config.Width)
	}

	switch config.Method {
	case "path_tracing":
		path_tracing.RenderPathTracing(sc, image)
	case "path_tracing_parallel":
		path_tracing.RenderPathTracingParallel(sc, image, 10)
	case "vanilla_radiance_cascade":
		radiance_cascade.NewRadianceCascade(sc, image, false).Render()
	case "bilinear_fix_radiance_cascade":
		radiance_cascade.NewRadianceCascade(sc, image, true).Render()
	case "error":
		//truth := NewSampledImageFromFile("ring_shadow.raw")
		//truth.Error(img)
		//truth.StorePNG("diff.png")
	case "plot":
		PlotCascade()
		PlotCascadeBilinearFix()
		PlotProbeCenter()
		PlotProbeCascadesNonSpatial()
		PlotCascadeBilinearFixSimple()
		return
		/*
			PlotSignedDistance()
			PlotCascade2()
			PlotCascade3()
			PlotEnergy(img)
		*/
	default:
		panic("Unknown method")
	}

	//img := primitives.NewSampledImageFromJpeg("assets/pexels-fwstudio-33348-129731.jpg")
	//img := primitives.NewSampledImageFromJpeg("assets/Texture_P7150102.JPG")
	//image.Blend(img)

	image.Store(config.OutputFilename)

}
