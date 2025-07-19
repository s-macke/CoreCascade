package main

import (
	"fmt"
	"image"
	"math"
	"math/rand/v2"
)

// Implement https://www.shadertoy.com/view/3tsXzB

func RenderPixel(s *Scene, uv Vec2, samples int) Color {
	col := Color{R: 0, G: 0, B: 0}
	for i := 0; i <= samples; i++ {
		angle := rand.Float64() * 2. * math.Pi
		dir := Vec2{
			X: math.Cos(angle),
			Y: math.Sin(angle),
		}
		ray := Ray{uv, dir}
		_, c := s.Intersect(ray, 4.0)
		col.Add(c)
	}
	return col
}

func RenderPathTracing(scene *Scene) *image.RGBA {
	const WIDTH, HEIGHT = 800, 800
	const SAMPLES = 100
	img := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
	energy := 0.

	for y := 0; y < HEIGHT; y++ {
		fmt.Println(y, "/", HEIGHT)
		for x := 0; x < WIDTH; x++ {
			// Convert pixel coordinates to normalized device coordinates
			uv := Vec2{X: (float64(x)/WIDTH)*2 - 1, Y: (float64(y)/HEIGHT)*2 - 1}
			col := RenderPixel(scene, uv, SAMPLES)
			col.Div(float64(SAMPLES))
			energy += col.Intensity()
			img.Set(x, y, col.ToSRGBA())
		}
	}
	fmt.Println("Energy", energy)
	return img
}
