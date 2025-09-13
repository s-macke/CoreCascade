package render

import (
	"CoreCascade/primitives"
	"CoreCascade/render/path_tracing"
	"CoreCascade/render/radiance_cascade"
	"CoreCascade/scene"
	"CoreCascade/scene/grid"
	"CoreCascade/scene/sdf"
	"fmt"
)

func ToSecondaryLight(sc scene.Scene, image *primitives.SampledImage) *grid.Scene {

	var secondaryLight *grid.Scene
	switch m := sc.(type) {
	case *grid.Scene:
		secondaryLight = m
	case *sdf.Scene:
		secondaryLight = grid.NewSceneFromSDF(image.Width, image.Height, m)
	default:
		panic("Unsupported scene type")
	}

	for x := 0; x < secondaryLight.Width; x++ {
		for y := 0; y < secondaryLight.Height; y++ {
			mat := &secondaryLight.M[y][x].Material
			fluence := image.GetColor(x, y)
			mat.Emissive.R = fluence.R * mat.Diffuse.R
			mat.Emissive.G = fluence.G * mat.Diffuse.G
			mat.Emissive.B = fluence.B * mat.Diffuse.B
		}
	}
	return secondaryLight
}

func MultiPassRenderer(sc scene.Scene, image *primitives.SampledImage, method string) {
	Pass(sc, image, method)

	secondaryLight := ToSecondaryLight(sc, image)
	if secondaryLight.IsBlack() { // No secondary light. E.g. no diffuse light sources
		fmt.Println("No secondary light. Skipping second pass.")
		return
	}
	secondaryImage := primitives.NewSampledImage(image.Width, image.Height)
	Pass(secondaryLight, secondaryImage, method)
	image.Add(secondaryImage)
}

func Pass(sc scene.Scene, image *primitives.SampledImage, method string) {

	switch method {
	case "path_tracing":
		path_tracing.RenderPathTracing(sc, image)
	case "path_tracing_parallel":
		path_tracing.RenderPathTracingParallel(sc, image, 10)
	case "vanilla_radiance_cascade":
		radiance_cascade.NewRadianceCascade(sc, image, false).Render()
	case "bilinear_fix_radiance_cascade":
		radiance_cascade.NewRadianceCascade(sc, image, true).Render()
	}
}
