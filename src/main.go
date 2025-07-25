package main

import (
	"CoreCascade/primitives"
	"CoreCascade/scene"
	"fmt"
)

func NewScene(sceneAsString string) *scene.Scene {
	switch sceneAsString {
	case "center":
		return scene.NewSceneCenter()
	case "shadows":
		return scene.NewSceneShadows()
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
	fmt.Printf("Size: %dx%d\n", config.Width, config.Height)
	fmt.Printf("Scene: %s\n", config.Scene)
	sc := NewScene(config.Scene)

	var image *primitives.SampledImage
	if config.InputFilename != "" {
		image = primitives.NewSampledImageFromFile(config.InputFilename)
	} else {
		image = primitives.NewSampledImage(config.Height, config.Width)
	}

	switch config.Method {
	case "path_tracing":
		RenderPathTracing(sc, image)
	case "path_tracing_parallel":
		RenderPathTracingParallel(sc, image, 100)
	case "vanilla_radiance_cascade":
		NewRadianceCascadeVanilla(sc, image).Render()
	default:
		panic("Unknown method")
	}

	image.Store(config.OutputFilename)

	//truth := NewSampledImageFromFile("ring_shadow.raw")
	//truth.Error(img)
	//truth.StoreImage("diff.png")

	/*
		PlotSignedDistance()
		PlotCascade2()
		PlotCascade3()
		PlotCascade4()
		PlotCascade5()
	*/
	//PlotCascade()
	//PlotEnergy(img)

}
