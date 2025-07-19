package main

import "math"

type sdObject interface {
	Distance(p Vec2) float64
	GetColor() Color
}

// Circle represents a circle with a center and radius.
type Circle struct {
	Center Vec2
	Radius float64
	Color  Color
}

// sdCircle calculates the signed distance from a point p to a circle c.
// It returns a negative value if the point is inside the circle,
// a positive value if it is outside, and 0 if it is on the circle.
func (c *Circle) Distance(p Vec2) float64 {
	distance := math.Sqrt(((p.X-c.Center.X)*(p.X-c.Center.X) + (p.Y-c.Center.Y)*(p.Y-c.Center.Y)))
	return distance - c.Radius
}

func (c *Circle) GetColor() Color {
	return c.Color
}

// Box represents an axis-aligned bounding box.
// Center is the geometric center of the box.
// HalfSize is a vector representing half of the width and height.
type Box struct {
	Center   Vec2
	HalfSize Vec2
	Color    Color
}

func (b *Box) GetColor() Color {
	return b.Color
}

// sdBox calculates the signed distance from a point p to an axis-aligned box b.
// It returns a negative value if the point is inside the box,
// a positive value if it is outside, and 0 if it is on the boundary.
func (b *Box) Distance(p Vec2) float64 {
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

func NewScene() *Scene {
	s := &Scene{}
	/*
		s.objects = []sdObject{
			//&Circle{Center: Vec2{X: -0.0, Y: 0}, Radius: 0.1, Color: Color{R: 1, G: 1, B: 1}},
			//&Box{Center: Vec2{X: -1., Y: 0.}, HalfSize: Vec2{X: 0.1, Y: 1.0}, Color: Color{R: 1., G: 1., B: 1.}},
			//&Box{Center: Vec2{X: -0.4, Y: 0.5}, HalfSize: Vec2{X: 0.1, Y: 0.5}, Color: Color{R: 0., G: 0, B: 0.0}},
		}
	*/
	/*
		s.objects = []sdObject{
			//&Box{Center: Vec2{X: -1., Y: 0.00}, HalfSize: Vec2{X: 0.1, Y: 4.00}, Color: NewSRGBColor(1., 0., 1.)},

			&Box{Center: Vec2{X: -1., Y: -1.00}, HalfSize: Vec2{X: 0.1, Y: 0.125}, Color: NewSRGBColor(1., 0., 1.)},
			&Box{Center: Vec2{X: -1., Y: -0.75}, HalfSize: Vec2{X: 0.1, Y: 0.125}, Color: NewSRGBColor(1., 1., 1.)},
			&Box{Center: Vec2{X: -1., Y: -0.5}, HalfSize: Vec2{X: 0.1, Y: 0.125}, Color: NewSRGBColor(1., 1., 0.)},
			&Box{Center: Vec2{X: -1., Y: -0.25}, HalfSize: Vec2{X: 0.1, Y: 0.125}, Color: NewSRGBColor(0., 0., 1.)},
			&Box{Center: Vec2{X: -1., Y: 0.0}, HalfSize: Vec2{X: 0.1, Y: 0.125}, Color: NewSRGBColor(1., 0.5, 0.5)},
			&Box{Center: Vec2{X: -1., Y: 0.25}, HalfSize: Vec2{X: 0.1, Y: 0.125}, Color: NewSRGBColor(0., 1., 0.)},
			&Box{Center: Vec2{X: -1., Y: 0.5}, HalfSize: Vec2{X: 0.1, Y: 0.125}, Color: NewSRGBColor(0.5, 0.5, 1.)},
			&Box{Center: Vec2{X: -1., Y: 0.75}, HalfSize: Vec2{X: 0.1, Y: 0.125}, Color: NewSRGBColor(1., 0., 0.)},
			&Box{Center: Vec2{X: -1., Y: 1.00}, HalfSize: Vec2{X: 0.1, Y: 0.125}, Color: NewSRGBColor(0.5, 1., 0.5)},

			&Box{Center: Vec2{X: -0., Y: -3.6}, HalfSize: Vec2{X: 0.02, Y: 3.55}, Color: Color{R: 0., G: 0, B: 0.0}},
			&Box{Center: Vec2{X: -0., Y: 3.6}, HalfSize: Vec2{X: 0.02, Y: 3.55}, Color: Color{R: 0., G: 0, B: 0.0}},
		}
	*/

	s.objects = []sdObject{
		&Circle{Center: Vec2{X: 0., Y: 0.}, Radius: 0.05, Color: Color{R: 1., G: 1, B: 1.}},
	}

	for i := 0; i < 10; i++ {
		x := math.Cos(float64(i)*2.*math.Pi/10.) * 0.5
		y := math.Sin(float64(i)*2.*math.Pi/10.) * 0.5
		s.objects = append(s.objects, &Circle{Center: Vec2{X: x, Y: y}, Radius: float64(i)/100. + 0.01, Color: Color{R: 0., G: 0, B: 0.}})
	}

	/*
		for i := 0; i < len(s.objects); i++ {
			switch obj := s.objects[i].(type) {
			case *Circle:
				obj.Color.Mul(4.)
			case *Box:
				obj.Color.Mul(4.)
			}
		}
	*/
	return s
}

func (s *Scene) sd(p Vec2) (float64, Color) {
	// Calculate the signed distance to the circle and box
	c := Color{R: 0., G: 0., B: 0.}
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

func (s *Scene) Intersect(r Ray, tmax float64) (float64, Color) {
	black := Color{R: 0., G: 0., B: 0.}
	t := 0.
	for j := 0; j < 50; j++ {
		p := r.Trace(t)
		if p.X < -2.1 || p.X > 2.1 || p.Y < -2.1 || p.Y > 2.1 {
			return 0., black // Out of bounds
		}
		d, c := s.sd(p)
		if d < 1e-3 {
			return 1., c
		}
		t += d
		if t > tmax {
			return 0., black
		}
	}
	return 0., black
}
