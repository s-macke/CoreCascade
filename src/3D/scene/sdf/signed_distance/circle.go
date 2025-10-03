package signed_distance

import (
	"CoreCascade3D/primitives"
	math "github.com/chewxy/math32"
	"vector"
)

// Circle represents a circle with a center and radius.
type Circle struct {
	Center   vector.Vec2
	Radius   float32
	Material primitives.Material
	Low      float32 // Height of the material in scene units. Used in 2.5D-3D rendering.
	High     float32 // Height of the material in scene units. Used in 2.5D-3D rendering.
}

// sdCircle calculates the signed distance from a point p to a circle c.
// It returns a negative value if the point is inside the circle,
// a positive value if it is outside, and 0 if it is on the circle.
func (c *Circle) Distance2D(p vector.Vec2) float32 {
	distance := math.Sqrt((p.X-c.Center.X)*(p.X-c.Center.X) + (p.Y-c.Center.Y)*(p.Y-c.Center.Y))
	return distance - c.Radius
}

func (c *Circle) GetMaterial() *primitives.Material {
	return &c.Material
}

func (c *Circle) Distance(p vector.Vec3) float32 {
	d := c.Distance2D(p.ToVec2())

	h := 0.5 * (c.High - c.Low)
	zc := 0.5 * (c.Low + c.High)
	d2 := math.Abs(p.Z-zc) - h
	//return max(d, d2)
	qx := math.Max(d, 0)
	qy := math.Max(d2, 0)
	outside := math.Hypot(qx, qy)
	inside := math.Min(math.Max(d, d2), 0)
	return outside + inside

	/*
			if d < 0 {
				if p.Z < c.Material.Low {
					return min((c.Material.Low - p.Z), d)
				}
				if p.Z > c.Material.High {
					return min((p.Z - c.Material.High), d)
				}
			}
		return d

	*/
}
