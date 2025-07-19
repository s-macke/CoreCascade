package main

import (
	"fmt"
	"image"
	"sync"
)

// Implement https://www.shadertoy.com/view/3tsXzB

type SampledPixel struct {
	Color Color
	Count int
}

type SampledImage struct {
	pixels        [][]SampledPixel
	width, height int
}

func NewSampledImage(width, height int) *SampledImage {
	s := &SampledImage{
		width:  width,
		height: height,
	}
	s.pixels = make([][]SampledPixel, height)
	for i := range s.pixels {
		s.pixels[i] = make([]SampledPixel, width)
		for j := range s.pixels[i] {
			s.pixels[i][j] = SampledPixel{Color: Color{R: 0, G: 0, B: 0}, Count: 0}
		}
	}
	return s
}

func (s *SampledImage) ToImage() *image.RGBA {
	energy := 0.
	img := image.NewRGBA(image.Rect(0, 0, s.width, s.height))
	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			p := s.pixels[y][x]
			c := p.Color
			c.Div(float64(p.Count))
			energy += c.Intensity()
			img.Set(x, y, c.ToSRGBA())
		}
	}
	fmt.Println("Energy", energy)
	return img
}

func RenderPathTracingIteration(scene *Scene, s *SampledImage, samples int) {
	var wg sync.WaitGroup

	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			// Convert pixel coordinates to normalized device coordinates
			wg.Add(1)
			go func(x, y int) {
				defer wg.Done()
				uv := Vec2{X: (float64(x)/float64(s.width))*2 - 1, Y: (float64(y)/float64(s.height))*2 - 1}
				col := RenderPixel(scene, uv, samples)
				s.pixels[y][x].Color.Add(col)
				s.pixels[y][x].Count += samples
			}(x, y)
		}
	}
	wg.Wait()
}

func RenderPathTracingParallel(scene *Scene) *image.RGBA {
	const WIDTH, HEIGHT = 800, 800
	const SAMPLES = 50
	const ITER = 1000
	s := NewSampledImage(WIDTH, HEIGHT)
	for i := 1; i <= ITER; i++ {
		fmt.Printf("Iteration %d / %d\n", i, ITER)
		RenderPathTracingIteration(scene, s, SAMPLES)
		StoreImage(s.ToImage(), "iter.png")
	}
	return s.ToImage()
}
