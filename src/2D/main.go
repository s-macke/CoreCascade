package main

import (
	"CoreCascade2D/render"
	"CoreCascade2D/render/light_propagation_volumes"
	"CoreCascade2D/scene"
	"CoreCascade2D/scene/scenes"
	"fmt"
	"linear_image"
	"strings"
)

func NewScene(sceneAsString string, time float64) scene.Scene {
	sceneAsString = strings.ToLower(sceneAsString)
	var s scene.Scene = nil
	switch sceneAsString {
	case "center":
		s = scenes.NewSceneCenter()
	case "directional":
		s = scenes.NewSceneDirectional(time)
	case "shadows":
		s = scenes.NewSceneShadows(time)
	case "pinhole":
		s = scenes.NewScenePinhole()
	case "penumbra":
		s = scenes.NewScenePenumbra()
	case "beam":
		s = scenes.NewSceneBeam()
	case "title":
		s = scenes.NewSceneTitle(time)
	case "fluid":
		s = scenes.NewSceneFluid(time, 1)
	case "absorption":
		s = scenes.NewSceneAbsorption(time)
	default:
		panic("Unknown scene: " + sceneAsString)
	}
	return s
}

func Special(config flagConfig) {
	image := linear_image.NewSampledImage(config.Height, config.Width)
	part := linear_image.NewSampledImage(config.Height, config.Width)
	for i := 1; i <= 6; i++ {
		//i := 1
		s := scenes.NewSceneFluid(config.Time, i)
		render.MultiPassRenderer(s, part, config.Method, config.Multipass)
		part.Store(config.OutputFilename + fmt.Sprintf("_part%02d", i))
		image.Add(part)
		//image.Store(config.OutputFilename + fmt.Sprintf("_part%02d", i))
	}
	image.Store(config.OutputFilename)
}

func main() {
	config := parseConfig()
	fmt.Printf("Scene: %s\n", config.Scene)
	fmt.Printf("Size: %dx%d\n", config.Width, config.Height)
	fmt.Printf("Method: %s\n", config.Method)
	//Special(config)
	//return
	sc := NewScene(config.Scene, config.Time)

	var img *linear_image.SampledImage
	if config.InputFilename != "" {
		fmt.Println("Reading image from file", config.InputFilename)
		img = linear_image.NewSampledImageFromFile(config.InputFilename)
	} else {
		img = linear_image.NewSampledImage(config.Height, config.Width)
	}

	switch config.Method {
	case "path_tracing", "path_tracing_parallel", "vanilla_radiance_cascade", "bilinear_fix_radiance_cascade":
		render.MultiPassRenderer(sc, img, config.Method, config.Multipass)
	case "light_propagation_volumes":
		light_propagation_volumes.LightPropagationVolume(sc, img)
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
	/*
		img := primitives.NewSampledImageFromJpeg("assets/pexels-fwstudio-33348-129731.jpg")
		//img := primitives.NewSampledImageFromJpeg("assets/Texture_P7150102.JPG")
		image.Blend(img)
	*/
	img.Store(config.OutputFilename)
}
