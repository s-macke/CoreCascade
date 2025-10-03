package grid

import (
	"CoreCascade3D/primitives"
	"color"
	math "github.com/chewxy/math32"
	"vector"
)

type Voxel struct {
	Material primitives.Material
}

type Scene struct {
	Width, Height, Depth int
	M                    [][][]Voxel
}

func NewScene(width, height, depth int) *Scene {
	s := &Scene{
		Width:  width,
		Height: height,
		Depth:  depth,
	}
	s.M = make([][][]Voxel, height)
	for i := range s.M {
		s.M[i] = make([][]Voxel, width)
		for j := range s.M[i] {
			s.M[i][j] = make([]Voxel, depth)
			for k := range s.M[i][j] {
				s.M[i][j][k].Material = primitives.VoidMaterial
			}
		}
	}
	return s
}

func (s *Scene) XYZToIndex(p vector.Vec3) (x, y, z int, outside bool) {
	// from (-1, -1, 0) to (1, 1, 0.1)
	x = int((p.X + 1.) / 2. * float32(s.Width))
	y = int((p.Y + 1.) / 2. * float32(s.Height))
	z = int(p.Z * 10. * float32(s.Depth))
	outside = x < 0 || y < 0 || z < 0 || x >= s.Width || y >= s.Height || z >= s.Depth
	return x, y, z, outside
}

// IndexToSceneUVW from (-1, -1, 0) to (1, 1, 0.1)
func (s *Scene) IndexToXYZ(x, y, z int) vector.Vec3 {
	return vector.Vec3{
		X: (float32(x)/float32(s.Width))*2. - 1.,
		Y: (float32(y)/float32(s.Height))*2. - 1.,
		Z: float32(z) / float32(s.Depth) * 0.1,
	}
}

func (s *Scene) GetMaterial(p vector.Vec3) primitives.Material {
	x, y, z, outside := s.XYZToIndex(p)
	if outside {
		return primitives.VoidMaterial
	}
	m := &s.M[y][x][z].Material
	return *m // trilinear maybe?
}

func (s *Scene) Trace3D2(r vector.Ray3D, tmax float32) (vis float32, c color.Color) {
	vis = 1.0
	const eps = 1e-4
	emission := color.Color{R: 1.0, G: 1.0, B: 1.0} // fake emission light at end

	dx := 2. / float32(s.Width)
	dz := 0.1 / float32(s.Depth)
	stepsize := vector.Vec3{
		X: dx / math.Abs(r.Dir.X),
		Y: dx / math.Abs(r.Dir.Y),
		Z: dz / math.Abs(r.Dir.Z),
	}
	dt := min(stepsize.X, stepsize.Y, stepsize.Z) * 0.2

	for t := float32(0.0); t < tmax; t += dt {
		p := r.Trace(t)
		x, y, z, outside := s.XYZToIndex(p)
		if outside {
			// out of bounds
			//continue
			emission.Mul(vis)
			c.Add(emission)
			return vis, c
		}
		v := s.M[y][x][z]
		sa := v.Material.Absorption
		// Inside medium? integrate absorption + volumetric emission over the step
		// If Emissive here is per-length volume emission, integrate closed-form
		loss := math.Exp(-sa * dt)
		vis *= loss
		/*
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
		*/
		// Early terminate if basically fully absorbed
		if vis < eps {
			return vis, c // basically fully absorbed
		}
	}
	emission.Mul(vis)
	c.Add(emission)
	return vis, c
}

func (s *Scene) Trace(ray vector.Ray3D, tmax float32) (vis float32, c color.Color) {
	vis = 1.0
	const eps = 1e-4
	emission := color.Color{R: 1.0, G: 1.0, B: 1.0} // fake emission light at end

	dx := 2. / float32(s.Width)
	dz := 0.1 / float32(s.Depth)

	sx := 0
	if ray.Dir.X > 0 {
		sx = 1
	} else if ray.Dir.X < 0 {
		sx = -1
	}
	sy := 0
	if ray.Dir.Y > 0 {
		sy = 1
	} else if ray.Dir.Y < 0 {
		sy = -1
	}
	sz := 0
	if ray.Dir.Z > 0 {
		sz = 1
	} else if ray.Dir.Z < 0 {
		sz = -1
	}

	// Reciprocals (safe for zeros -> +Inf)
	inf := math.Inf(1)
	tDeltaX := inf
	tDeltaY := inf
	tDeltaZ := inf
	if sx != 0 {
		tDeltaX = math.Abs(dx / ray.Dir.X)
	}
	if sy != 0 {
		tDeltaY = math.Abs(dx / ray.Dir.Y)
	}
	if sz != 0 {
		tDeltaZ = math.Abs(dz / ray.Dir.Z)
	}

	ix, iy, iz, outside := s.XYZToIndex(ray.P)
	if outside {
		return vis, c
	}
	// Compute initial t to the first boundary in each axis (tMax*).
	// Using gridMin as the world-space origin of cell (0,0,0).
	cellMin := s.IndexToXYZ(ix, iy, iz)

	var nextBoundaryX, nextBoundaryY, nextBoundaryZ float32
	if sx > 0 {
		nextBoundaryX = cellMin.X + dx
	} else if sx < 0 {
		nextBoundaryX = cellMin.X
	}
	if sy > 0 {
		nextBoundaryY = cellMin.Y + dx
	} else if sy < 0 {
		nextBoundaryY = cellMin.Y
	}
	if sz > 0 {
		nextBoundaryZ = cellMin.Z + dz
	} else if sz < 0 {
		nextBoundaryZ = cellMin.Z
	}

	tMaxX := inf
	tMaxY := inf
	tMaxZ := inf
	if sx != 0 {
		tMaxX = (nextBoundaryX - ray.P.X) / ray.Dir.X
	}
	if sy != 0 {
		tMaxY = (nextBoundaryY - ray.P.Y) / ray.Dir.Y
	}
	if sz != 0 {
		tMaxZ = (nextBoundaryZ - ray.P.Z) / ray.Dir.Z
	}

	// If we start exactly on a boundary, ensure we don't get a tiny negative due to FP
	if tMaxX < 0 && sx != 0 {
		tMaxX = 0
	}
	if tMaxY < 0 && sy != 0 {
		tMaxY = 0
	}
	if tMaxZ < 0 && sz != 0 {
		tMaxZ = 0
	}
	var tPrev float32

	for {
		// Decide which axis boundary we hit next
		tNext := tMaxX
		axis := 0 // 0->X, 1->Y, 2->Z
		if tMaxY < tNext {
			tNext = tMaxY
			axis = 1
		}
		if tMaxZ < tNext {
			tNext = tMaxZ
			axis = 2
		}

		// Record current cell's segment
		segLen := tNext - tPrev
		if segLen < 0 {
			segLen = 0 // guard FP
		}
		//hits = append(hits, CellHit{X: ix, Y: iy, Z: iz, Length: segLen})
		outside = ix < 0 || iy < 0 || iz < 0 || ix >= s.Width || iy >= s.Height || iz >= s.Depth
		if outside {
			break
		}
		v := s.M[iy][ix][iz]
		sa := v.Material.Absorption
		loss := math.Exp(-sa * segLen)
		vis *= loss
		if vis < eps {
			break // basically fully absorbed
		}

		tPrev = tNext

		// Step to next cell across the chosen boundary
		switch axis {
		case 0:
			ix += sx
			tMaxX += tDeltaX
		case 1:
			iy += sy
			tMaxY += tDeltaY
		case 2:
			iz += sz
			tMaxZ += tDeltaZ
		}
	}

	emission.Mul(vis)
	c.Add(emission)
	return vis, c

}
