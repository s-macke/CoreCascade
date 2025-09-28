package render

import (
	"CoreCascade2D/render/path_tracing"
	"CoreCascade2D/render/radiance_cascade"
	"CoreCascade2D/scene"
	"CoreCascade2D/scene/grid"
	"CoreCascade2D/scene/sdf"
	"fmt"
	"linear_image"
	"vector"
)

func ToSecondaryLight(sc scene.Scene, image *linear_image.SampledImage) *grid.Scene {
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

func AddDiffuse(sc scene.Scene, image *linear_image.SampledImage) {
	for x := 0; x < image.Width; x++ {
		for y := 0; y < image.Height; y++ {
			uv := vector.Vec2{X: (float64(x)/float64(image.Width))*2 - 1, Y: (float64(y)/float64(image.Height))*2 - 1}
			mat := sc.GetMaterial(uv)
			fluence := image.GetColor(x, y)
			fluence.R += fluence.R * mat.Diffuse.R * 0.99
			fluence.G += fluence.G * mat.Diffuse.G * 0.99
			fluence.B += fluence.B * mat.Diffuse.B * 0.99
			image.SetColor(x, y, fluence)
		}
	}
}

func MultiPassRenderer(sc scene.Scene, image *linear_image.SampledImage, method string, multipass int) {
	Pass(sc, image, method)

	switch multipass {
	case 0:
		AddDiffuse(sc, image)
	case 1:
		secondaryLight := ToSecondaryLight(sc, image)
		if secondaryLight.IsBlack() { // No secondary light. E.g. no diffuse light sources
			fmt.Println("No secondary light. Skipping second pass.")
			return
		}
		secondaryImage := linear_image.NewSampledImage(image.Width, image.Height)
		Pass(secondaryLight, secondaryImage, method)
		image.Add(secondaryImage)
	default:
		panic("Unsupported multipass method")
	}
}

func Pass(sc scene.Scene, image *linear_image.SampledImage, method string) {

	switch method {
	case "path_tracing":
		path_tracing.RenderPathTracing(sc, image)
	case "path_tracing_parallel":
		path_tracing.RenderPathTracingParallel(sc, image, 10)
	case "vanilla_radiance_cascade":
		radiance_cascade.NewRadianceCascade(sc, image, false).Render()
	case "bilinear_fix_radiance_cascade":
		radiance_cascade.NewRadianceCascade(sc, image, true).Render()
	default:
		panic("Unsupported method")
	}
}
