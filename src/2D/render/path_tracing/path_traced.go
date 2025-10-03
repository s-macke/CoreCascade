package path_tracing

import (
	"CoreCascade2D/scene"
	"color"
	"fmt"
	math "github.com/chewxy/math32"
	"linear_image"
	"math/rand/v2"
	"vector"
)

func RenderPixel(s scene.Scene, uv vector.Vec2, samples int) color.Color {
	col := color.Black
	for i := 0; i <= samples; i++ {
		angle := rand.Float32() * 2. * math.Pi
		dir := vector.NewVec2fromAngle(angle)
		ray := vector.Ray2D{P: uv, Dir: dir}
		_, c := s.Trace(ray, 4.0)
		//c.Mul(2. * math.Pi) // compensate for random angle sampling
		col.Add(c)
	}
	return col
}

func RenderPathTracing(scene scene.Scene, s *linear_image.SampledImage) {
	const SAMPLES = 100

	for y := 0; y < s.Height; y++ {
		fmt.Println(y, "/", s.Height)
		for x := 0; x < s.Width; x++ {
			// Convert pixel coordinates to scene coordinates
			uv := vector.Vec2{X: (float32(x)/float32(s.Width))*2 - 1, Y: (float32(y)/float32(s.Height))*2 - 1}
			col := RenderPixel(scene, uv, SAMPLES)
			s.AddColorSamples(x, y, col, SAMPLES)
		}
	}
	fmt.Println("Energy", s.Energy())
}
