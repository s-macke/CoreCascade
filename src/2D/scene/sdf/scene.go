package sdf

import (
	"CoreCascade2D/primitives"
	"color"
	math "github.com/chewxy/math32"
	"vector"
)

type SdObject interface {
	Distance(p vector.Vec2) float32
	GetMaterial() *primitives.Material
}

type Scene struct {
	Objects []SdObject
}

func (s *Scene) GetExtent() (vector.Vec2, vector.Vec2) {
	// Return the extent of the scene
	// This is used to define the bounds for rendering or ray tracing
	return vector.Vec2{X: -1.0, Y: -1.0}, vector.Vec2{X: 1.0, Y: 1.0}
}

func (s *Scene) GetMaterial(p vector.Vec2) primitives.Material {
	_, m := s.SignedDistance(p)
	return m
}

func (s *Scene) SignedDistance(p vector.Vec2) (float32, primitives.Material) {
	// Calculate the total signed distance to the objects
	m := primitives.VoidMaterial
	negative := false

	d := float32(math.MaxFloat32) // Initialize with a large distance
	for _, obj := range s.Objects {
		distance := obj.Distance(p)
		if distance < 0 {
			negative = true
			// add material properties in case they overlap
			mtemp := obj.GetMaterial()
			m.Merge(mtemp)
		}
		if math.Abs(distance) < d {
			d = distance
		}
	}

	if negative { // make absolutely sure, that the distance is negative to indicate that we are inside a medium. Otherwise, we cannot do one optimization later
		d = -math.Abs(d)
	}
	return d, m
}

func (s *Scene) Trace(r vector.Ray2D, tmax float32) (visibility float32, c color.Color) {
	t := float32(0.)
	vis := float32(1.0)
	c = color.Black
	const eps = 1e-4
	const minStep = 0.01
	stop := false

	for j := 0; j < 128; j++ {
		p := r.Trace(t)

		// Scene bounds
		if p.X < -2.1 || p.X > 2.1 || p.Y < -2.1 || p.Y > 2.1 {
			return vis, c
		}

		d, m := s.SignedDistance(p)
		step := math.Max(math.Abs(d), minStep)
		if t+step > tmax {
			step = tmax - t
			stop = true
		}

		// Inside medium? integrate absorption + volumetric emission over the step
		if d < 0 {
			sa := math.Max(0.0, m.Absorption)
			e := m.Emission(r.Dir)
			// If Emissive here is per-length volume emission, integrate closed-form
			if sa > eps {
				k := 1.0 - math.Exp(-sa*step)
				c.R += vis * (e.R / sa) * k
				c.G += vis * (e.G / sa) * k
				c.B += vis * (e.B / sa) * k
				vis *= math.Exp(-sa * step)
			} else {
				// No absorption: pure additive over distance
				c.R += vis * e.R * step
				c.G += vis * e.G * step
				c.B += vis * e.B * step
			}
			// Early terminate if basically fully absorbed
			if vis < eps {
				return vis, c // basically fully absorbed
			}
		}

		t += step
		if stop {
			return vis, c
		}
	}

	return vis, c
}
