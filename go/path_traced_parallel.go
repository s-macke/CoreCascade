package main

import (
	"fmt"
	"sync"
)

// Implement https://www.shadertoy.com/view/3tsXzB

func RenderPathTracingIteration(scene *Scene, s *SampledImage, samples int) {
	var wg sync.WaitGroup

	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			wg.Add(1)
			go func(x, y int) {
				defer wg.Done()
				// Convert pixel coordinates to normalized device coordinates
				uv := Vec2{X: (float64(x)/float64(s.width))*2 - 1, Y: (float64(y)/float64(s.height))*2 - 1}
				col := RenderPixel(scene, uv, samples)
				s.AddColorSamples(x, y, col, samples)
			}(x, y)
		}
	}
	wg.Wait()
}

func RenderPathTracingParallel(scene *Scene) *SampledImage {
	const WIDTH, HEIGHT = 800, 800
	const SAMPLES = 200
	const ITER = 10000
	s := NewSampledImage(WIDTH, HEIGHT)
	for i := 1; i <= ITER; i++ {
		fmt.Printf("Iteration %d / %d\n", i, ITER)
		RenderPathTracingIteration(scene, s, SAMPLES)
		s.Store("iter")
	}
	return s
}
