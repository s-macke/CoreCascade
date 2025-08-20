package signed_distance

import (
	"CoreCascade/primitives"
	"math"
)

// Circle represents a circle with a center and radius.
type Circle struct {
	Center   primitives.Vec2
	Radius   float64
	Material primitives.Material
}

// sdCircle calculates the signed distance from a point p to a circle c.
// It returns a negative value if the point is inside the circle,
// a positive value if it is outside, and 0 if it is on the circle.
func (c *Circle) Distance(p primitives.Vec2) float64 {
	distance := math.Sqrt((p.X-c.Center.X)*(p.X-c.Center.X) + (p.Y-c.Center.Y)*(p.Y-c.Center.Y))
	return distance - c.Radius
}

func (c *Circle) GetMaterial() primitives.Material {
	return c.Material
}
