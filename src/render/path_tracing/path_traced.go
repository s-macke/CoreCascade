package path_tracing

import (
	"CoreCascade/primitives"
	"CoreCascade/scene"
	"fmt"
	"math"
	"math/rand/v2"
)

// Implement https://www.shadertoy.com/view/3tsXzB

func RenderPixel(s *scene.Scene, uv primitives.Vec2, samples int) primitives.Color {
	col := primitives.Color{R: 0, G: 0, B: 0}
	for i := 0; i <= samples; i++ {
		angle := rand.Float64() * 2. * math.Pi
		dir := primitives.Vec2{
			X: math.Cos(angle),
			Y: math.Sin(angle),
		}
		ray := primitives.Ray{uv, dir}
		_, c := s.Intersect(ray, 4.0)
		col.Add(c)
	}
	return col
}

func RenderPathTracing(scene *scene.Scene, s *primitives.SampledImage) {
	const SAMPLES = 100

	for y := 0; y < s.Height; y++ {
		fmt.Println(y, "/", s.Height)
		for x := 0; x < s.Width; x++ {
			// Convert pixel coordinates to scene coordinates
			uv := primitives.Vec2{X: (float64(x)/float64(s.Width))*2 - 1, Y: (float64(y)/float64(s.Height))*2 - 1}
			col := RenderPixel(scene, uv, SAMPLES)
			s.AddColorSamples(x, y, col, SAMPLES)
		}
	}
	fmt.Println("Energy", s.Energy())
}
