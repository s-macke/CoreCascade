package scene

import (
	"CoreCascade/primitives"
)

type sdObject interface {
	Distance(p primitives.Vec2) float64
	GetMaterial() primitives.Material
}

type Scene struct {
	objects []sdObject
}

func (s *Scene) GetExtent() (primitives.Vec2, primitives.Vec2) {
	// Return the extent of the scene
	// This is used to define the bounds for rendering or ray tracing
	return primitives.Vec2{X: -1.0, Y: -1.0}, primitives.Vec2{X: 1.0, Y: 1.0}
}

func (s *Scene) SignedDistance(p primitives.Vec2) (float64, primitives.Material) {
	// Calculate the signed distance to the circle and box
	m := primitives.Material{
		Emissive:   primitives.Black,
		Absorption: 0,
	}
	d := 1e99 // Initialize with a large distance
	for _, obj := range s.objects {
		distance := obj.Distance(p)
		if distance < d {
			d = distance
			m = obj.GetMaterial()
		}
	}
	return d, m
}

func (s *Scene) Intersect(r primitives.Ray, tmax float64) (bool, primitives.Color) {
	black := primitives.Color{R: 0., G: 0., B: 0.}
	t := 0.
	for j := 0; j < 50; j++ {
		p := r.Trace(t)
		if p.X < -2.1 || p.X > 2.1 || p.Y < -2.1 || p.Y > 2.1 {
			return false, black // Out of bounds
		}
		d, m := s.SignedDistance(p)
		if d < 1e-3 {
			return true, m.Emissive
		}
		t += max(d, 0.01) // define some minimum step size, which is determined by the smallest object in the scene
		if t > tmax {
			return false, black
		}
	}
	return false, black
}
