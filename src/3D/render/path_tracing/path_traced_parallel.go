package path_tracing

import (
	"CoreCascade3D/scene"
	"color"
	math "github.com/chewxy/math32"
	"linear_image"
	"sync"
	"vector"
)

type PathTracingParallel struct {
	Width  int
	Height int
	Depth  int
	Scene  scene.Scene
	image  *linear_image.SampledImage
	layers []*linear_image.SampledImage
}

func NewPathTracing(s scene.Scene, image *linear_image.SampledImage) *PathTracingParallel {
	depth := 40
	var layers []*linear_image.SampledImage
	for layer := 0; layer < depth; layer++ {
		layers = append(layers, linear_image.NewSampledImage(image.Width, image.Height))
	}
	return &PathTracingParallel{
		Width:  image.Width,
		Height: image.Height,
		Depth:  depth,
		Scene:  s,
		image:  image,
		layers: layers,
	}
}

// IndexToSceneUVW from (-1, -1, 0) to (1, 1, 0.1)
func (pt *PathTracingParallel) IndexToSceneUVW(x, y, z int) vector.Vec3 {
	return vector.Vec3{
		X: (float32(x)/float32(pt.Width))*2. - 1.,
		Y: (float32(y)/float32(pt.Height))*2. - 1.,
		Z: float32(z) / float32(pt.Depth) * 0.1,
	}
}

func (pt *PathTracingParallel) RenderPixel(uvw vector.Vec3, samples int) color.Color {
	col := color.Black
	for i := 0; i <= samples; i++ {
		//dir := primitives.NewRandomUnitVec3()
		dir := vector.Vec3{X: 0. - uvw.X, Y: (-1.) - uvw.Y, Z: 0.09 - uvw.Z}
		length := dir.Normalize()
		ray := vector.Ray3D{P: uvw, Dir: dir}
		_, c := pt.Scene.Trace(ray, 4.0)
		c.Mul(1. / (length * length))
		col.Add(c)
	}
	return col
}

func (pt *PathTracingParallel) RenderPathTracingIteration(samples int) {
	var wg sync.WaitGroup
	for z := 0; z < pt.Depth; z++ {
		for y := 0; y < pt.Height; y++ {
			for x := 0; x < pt.Width; x++ {
				wg.Add(1)
				go func(x, y, z int) {
					defer wg.Done()
					// Convert pixel coordinates to scene coordinates
					uvw := pt.IndexToSceneUVW(x, y, z)
					col := pt.RenderPixel(uvw, samples)
					pt.layers[z].AddColorSamples(x, y, col, samples)
				}(x, y, z)
			}
		}
	}
	wg.Wait()
}

func (pt *PathTracingParallel) MergeLayers() {
	pt.image.Clear()
	dz := 0.1 / float32(pt.Depth)
	for y := 0; y < pt.Height; y++ {
		for x := 0; x < pt.Width; x++ {
			visibility := float32(1.0)
			color := color.Black
			for z := pt.Depth - 1; z >= 0; z-- {
				uvw := pt.IndexToSceneUVW(x, y, z)
				mat := pt.Scene.GetMaterial(uvw)
				fluence := pt.layers[z].GetColor(x, y)
				color.R += fluence.R * mat.Diffuse.R * visibility
				color.G += fluence.G * mat.Diffuse.G * visibility
				color.B += fluence.B * mat.Diffuse.B * visibility
				visibility *= math.Exp(-mat.Absorption * dz)
			}
			pt.image.SetColor(x, y, color)
		}
	}
}

func (pt *PathTracingParallel) Render(maxIterations int) {
	const SAMPLES = 1
	//for i := 1; i <= maxIterations; i++ {
	//fmt.Printf("Iteration %d / %d\n", i, maxIterations)
	pt.RenderPathTracingIteration(SAMPLES)
	pt.MergeLayers()
	//pt.image.Store("iter")
	//pt.layers[0].Store("iter")
	//}
}
