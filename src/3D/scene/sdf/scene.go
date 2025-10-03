package sdf

import (
	"CoreCascade3D/primitives"
	"color"
	math "github.com/chewxy/math32"
	"vector"
)

type SdObject interface {
	Distance(p vector.Vec3) float32
	GetMaterial() *primitives.Material
}

type Scene struct {
	Objects []SdObject
}

func (s *Scene) GetMaterial(p vector.Vec3) primitives.Material {
	_, m := s.SignedDistance(p)
	return m
}

func (s *Scene) SignedDistance(p vector.Vec3) (float32, primitives.Material) {
	// Calculate the total signed distance to the objects
	d := float32(math.MaxFloat32) // Initialize with a large distance
	for _, obj := range s.Objects {
		distance := obj.Distance(p)
		if distance < 0 {
			m := obj.GetMaterial()
			return distance, *m // in 3D, we do not merge materials, but return the first one we are inside of
		}
		if math.Abs(distance) < d {
			d = distance
		}
	}
	return d, primitives.VoidMaterial
}

func (s *Scene) Trace(r vector.Ray3D, tmax float32) (visibility float32, c color.Color) {
	t := float32(0.0)
	vis := float32(1.0)
	const eps = 1e-4
	const minStep = 0.01
	stop := false

	for j := 0; j < 128; j++ {
		p := r.Trace(t)

		// Scene bounds
		// TODO. The bounds must be larger for radiance cascades or aabb test
		if p.X < -3.0 || p.X > 3.0 || p.Y < -3.0 || p.Y > 3.0 || p.Z < -3. || p.Z > 3.0 {
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
			e := m.Emissive
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

		if stop {
			return vis, c
		}
		t += step
	}
	return vis, c
}
