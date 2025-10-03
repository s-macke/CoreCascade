package grid

import (
	"CoreCascade2D/primitives"
	"CoreCascade2D/scene/sdf"
	"color"
	math "github.com/chewxy/math32"
	"vector"
)

// TODO: there might be a problem with the stepsize close to tmax

type Voxel struct {
	Material primitives.Material
	distance float32
}

type Scene struct {
	Width, Height int
	M             [][]Voxel
}

func NewScene(width, height int) *Scene {
	s := &Scene{
		Width:  width,
		Height: height,
	}
	s.M = make([][]Voxel, height)
	for i := range s.M {
		s.M[i] = make([]Voxel, width)
		for j := range s.M[i] {
			s.M[i][j].Material = primitives.VoidMaterial
			s.M[i][j].distance = 0.0
		}
	}
	return s
}

func NewSceneFromSDF(width, height int, sdf *sdf.Scene) *Scene {
	s := NewScene(width, height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			uv := vector.Vec2{X: (float32(x)/float32(width))*2 - 1, Y: (float32(y)/float32(height))*2 - 1}
			d, m := sdf.SignedDistance(uv)
			s.M[y][x].Material = m
			s.M[y][x].distance = d
		}
	}
	return s
}

func (s *Scene) IsBlack() bool {
	for y := 0; y < s.Height; y++ {
		for x := 0; x < s.Width; x++ {
			if s.M[y][x].Material.Emissive.Intensity() > 1e-2 {
				return false
			}
		}
	}
	return true
}

func (s *Scene) GetMaterial(p vector.Vec2) primitives.Material {
	x := (p.X + 1.) / 2.0 * float32(s.Width)
	y := (p.Y + 1.) / 2.0 * float32(s.Height)
	if int(x) < 0 || int(y) < 0 || int(x) >= s.Width || int(y) >= s.Height {
		return primitives.VoidMaterial
	}
	return s.M[int(y)][int(x)].Material // bilinear maybe?
}

func (s *Scene) Trace(r vector.Ray2D, tmax float32) (vis float32, c color.Color) {
	vis = 1.0
	const eps = 1e-4

	// hit the scene?
	hit := IntersectRayAABB(r, vector.Vec2{X: -1.0, Y: -1.0}, vector.Vec2{X: 1.0, Y: 1.0})
	if !hit.Hit {
		return vis, c
	}

	tmin := max(0.0, hit.TEnter)
	tmax = min(tmax, hit.TExit)

	dt := 2. / float32(s.Width)
	for t := tmin; t < tmax; t += dt {
		p := r.Trace(t)

		x := (p.X + 1.) / 2.0 * float32(s.Width)
		y := (p.Y + 1.) / 2.0 * float32(s.Height)
		if int(x) < 0 || int(y) < 0 || int(x) >= s.Width || int(y) >= s.Height {
			continue
			// return vis, c
			// out of bounds
		}
		v := s.M[int(y)][int(x)]
		sa := math.Max(0.0, v.Material.Absorption)
		// Inside medium? integrate absorption + volumetric emission over the step
		// If Emissive here is per-length volume emission, integrate closed-form
		if sa > eps {
			loss := math.Exp(-sa * dt)
			k := 1.0 - loss
			c.R += vis * (v.Material.Emissive.R / sa) * k
			c.G += vis * (v.Material.Emissive.G / sa) * k
			c.B += vis * (v.Material.Emissive.B / sa) * k
			vis *= loss
		} else {
			// No absorption: pure additive over distance
			c.R += vis * v.Material.Emissive.R * dt
			c.G += vis * v.Material.Emissive.G * dt
			c.B += vis * v.Material.Emissive.B * dt
		}
		// Early terminate if basically fully absorbed
		if vis < eps {
			return vis, c // basically fully absorbed
		}
	}
	return vis, c
}
