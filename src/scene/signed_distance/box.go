package signed_distance

import (
	"CoreCascade/primitives"
	"math"
)

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
