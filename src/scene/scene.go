package scene

import (
	"CoreCascade/primitives"
	"math"
)

type sdObject interface {
	Distance(p primitives.Vec2) float64
	GetColor() primitives.Color
}

// Circle represents a circle with a center and radius.
type Circle struct {
	Center primitives.Vec2
	Radius float64
	Color  primitives.Color
}

// sdCircle calculates the signed distance from a point p to a circle c.
// It returns a negative value if the point is inside the circle,
// a positive value if it is outside, and 0 if it is on the circle.
func (c *Circle) Distance(p primitives.Vec2) float64 {
	distance := math.Sqrt(((p.X-c.Center.X)*(p.X-c.Center.X) + (p.Y-c.Center.Y)*(p.Y-c.Center.Y)))
	return distance - c.Radius
}

func (c *Circle) GetColor() primitives.Color {
	return c.Color
}

// Box represents an axis-aligned bounding box.
// Center is the geometric center of the box.
// HalfSize is a vector representing half of the width and height.
type Box struct {
	Center   primitives.Vec2
	HalfSize primitives.Vec2
	Color    primitives.Color
}

func (b *Box) GetColor() primitives.Color {
	return b.Color
}

// sdBox calculates the signed distance from a point p to an axis-aligned box b.
// It returns a negative value if the point is inside the box,
// a positive value if it is outside, and 0 if it is on the boundary.
func (b *Box) Distance(p primitives.Vec2) float64 {
	// 1. Translate the point so the box is centered at the origin
	p.X -= b.Center.X
	p.Y -= b.Center.Y

	// 2. Calculate the component-wise distance from the point to the box's surface
	dx := math.Abs(p.X) - b.HalfSize.X
	dy := math.Abs(p.Y) - b.HalfSize.Y

	// 3. Calculate the signed distance
	// The distance from the origin to the closest point on the box's surface.
	// We use max(dx, 0) and max(dy, 0) to only consider distances for axes where the point is outside the box.
	outsideDistance := math.Sqrt(math.Max(dx, 0)*math.Max(dx, 0) + math.Max(dy, 0)*math.Max(dy, 0))
	// The distance for a point inside the box is the largest of the negative distances to the edges.
	insideDistance := math.Min(math.Max(dx, dy), 0.0)
	return outsideDistance + insideDistance
}

type Scene struct {
	objects []sdObject
}

func (s *Scene) GetExtent() (primitives.Vec2, primitives.Vec2) {
	// Return the extent of the scene
	// This is used to define the bounds for rendering or ray tracing
	return primitives.Vec2{X: -1.0, Y: -1.0}, primitives.Vec2{X: 1.0, Y: 1.0}
}

func (s *Scene) SignedDistance(p primitives.Vec2) (float64, primitives.Color) {
	// Calculate the signed distance to the circle and box
	c := primitives.Color{R: 0., G: 0., B: 0.}
	d := 1e99 // Initialize with a large distance
	for _, obj := range s.objects {
		distance := obj.Distance(p)
		if distance < d {
			d = distance
			c = obj.GetColor()
		}
	}
	return d, c
}

func (s *Scene) Intersect(r primitives.Ray, tmax float64) (bool, primitives.Color) {
	black := primitives.Color{R: 0., G: 0., B: 0.}
	t := 0.
	for j := 0; j < 50; j++ {
		p := r.Trace(t)
		if p.X < -2.1 || p.X > 2.1 || p.Y < -2.1 || p.Y > 2.1 {
			return false, black // Out of bounds
		}
		d, c := s.SignedDistance(p)
		if d < 1e-3 {
			return true, c
		}
		t += max(d, 0.01) // define some minimum step size, which is determined by the smallest object in the scene
		if t > tmax {
			return false, black
		}
	}
	return false, black
}
