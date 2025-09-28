package grid

import (
	"math"
	"vector"
)

type AABBHit struct {
	Hit           bool
	TEnter, TExit float64
}

// IntersectRayAABB tests a ray against an axis-aligned box [min,max].
// Returns (hit, tEnter, tExit, hitPoint).
// - If no hit for t >= 0, hit is false.
func IntersectRayAABB(ray vector.Ray2D, min, max vector.Vec2) (hit AABBHit) {
	hit.Hit = false
	hit.TEnter = 0.0
	hit.TExit = math.Inf(1)

	// X slab
	if ray.Dir.X != 0 {
		t1 := (min.X - ray.P.X) / ray.Dir.X
		t2 := (max.X - ray.P.X) / ray.Dir.X
		if t1 > t2 {
			t1, t2 = t2, t1
		}

		hit.TEnter = math.Max(hit.TEnter, t1)
		hit.TExit = math.Min(hit.TExit, t2)
		if hit.TExit < hit.TEnter {
			return hit
		}
	} else {
		// Parallel to X slabs: must be inside them
		if ray.P.X < min.X || ray.P.X > max.X {
			return hit
		}
	}

	// Y slab
	if ray.Dir.Y != 0 {
		t1 := (min.Y - ray.P.Y) / ray.Dir.Y
		t2 := (max.Y - ray.P.Y) / ray.Dir.Y
		if t1 > t2 {
			t1, t2 = t2, t1
		}
		hit.TEnter = math.Max(hit.TEnter, t1)
		hit.TExit = math.Min(hit.TExit, t2)
		if hit.TExit < hit.TEnter {
			return hit
		}
	} else {
		// Parallel to Y slabs: must be inside them
		if ray.P.Y < min.Y || ray.P.Y > max.Y {
			return hit
		}
	}

	// Valid only for t >= 0 (forward along the ray)
	if hit.TExit < 0 {
		return hit
	}

	hit.Hit = true
	return hit
}
