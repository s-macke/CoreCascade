package main

import (
	"CoreCascade/primitives"
	"CoreCascade/scene"
	"fmt"
	"sync"
)

// Implement https://www.shadertoy.com/view/3tsXzB

func RenderPathTracingIteration(scene *scene.Scene, s *primitives.SampledImage, samples int) {
	var wg sync.WaitGroup

	for y := 0; y < s.Height; y++ {
		for x := 0; x < s.Width; x++ {
			wg.Add(1)
			go func(x, y int) {
				defer wg.Done()
				// Convert pixel coordinates to scene coordinates
				uv := primitives.Vec2{X: (float64(x)/float64(s.Width))*2 - 1, Y: (float64(y)/float64(s.Height))*2 - 1}
				col := RenderPixel(scene, uv, samples)
				s.AddColorSamples(x, y, col, samples)
			}(x, y)
		}
	}
	wg.Wait()
}

func RenderPathTracingParallel(scene *scene.Scene, s *primitives.SampledImage, maxIterations int) *primitives.SampledImage {
	const SAMPLES = 200
	for i := 1; i <= maxIterations; i++ {
		fmt.Printf("Iteration %d / %d\n", i, maxIterations)
		RenderPathTracingIteration(scene, s, SAMPLES)
		s.Store("iter")
	}
	return s
}
