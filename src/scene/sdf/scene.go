package sdf

import (
	"CoreCascade/primitives"
	"math"
)

type SdObject interface {
	Distance(p primitives.Vec2) float64
	GetMaterial() primitives.Material
}

type Scene struct {
	Objects []SdObject
}

func (s *Scene) GetExtent() (primitives.Vec2, primitives.Vec2) {
	// Return the extent of the scene
	// This is used to define the bounds for rendering or ray tracing
	return primitives.Vec2{X: -1.0, Y: -1.0}, primitives.Vec2{X: 1.0, Y: 1.0}
}

func (s *Scene) SignedDistance(p primitives.Vec2) (float64, primitives.Material) {
	// Calculate the total signed distance to the objects
	m := primitives.VoidMaterial
	negative := false

	d := 1e99 // Initialize with a large distance
	for _, obj := range s.Objects {
		distance := obj.Distance(p)
		if distance < 0 {
			negative = true
			// add material properties in case they overlap
			mtemp := obj.GetMaterial()
			m.Merge(&mtemp)
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

func (s *Scene) Intersect(r primitives.Ray, tmax float64) (visibility float64, c primitives.Color) {
	t := 0.0
	vis := 1.0
	const eps = 1e-4
	const minStep = 0.01

	for j := 0; j < 128; j++ {
		p := r.Trace(t)

		// Scene bounds
		if p.X < -2.1 || p.X > 2.1 || p.Y < -2.1 || p.Y > 2.1 {
			return vis, c
		}

		d, m := s.SignedDistance(p)
		step := math.Max(math.Abs(d), minStep)

		// Inside medium? integrate absorption + volumetric emission over the step
		if d < 0 {
			sa := math.Max(0.0, m.Absorption)
			// If Emissive here is per-length volume emission, integrate closed-form
			if sa > eps {
				k := 1.0 - math.Exp(-sa*step)
				c.R += vis * (m.Emissive.R / sa) * k
				c.G += vis * (m.Emissive.G / sa) * k
				c.B += vis * (m.Emissive.B / sa) * k
				vis *= math.Exp(-sa * step)
			} else {
				// No absorption: pure additive over distance
				c.R += vis * m.Emissive.R * step
				c.G += vis * m.Emissive.G * step
				c.B += vis * m.Emissive.B * step
			}
			// Early terminate if basically fully absorbed
			if vis < eps {
				return vis, c // basically fully absorbed
			}
		}

		t += step
		if t > tmax {
			return vis, c
		}
	}

	return vis, c
}
