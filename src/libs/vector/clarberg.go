package vector

import math "github.com/chewxy/math32"

// Clarberg spherical mapping for equal-area mapping between uniform 2D samples and sphere directions.
// Based on Petrik Clarberg's octahedral/hemispherical mapping technique.
// Reference: "Fast Equal-Area Mapping of the (Hemi)Sphere using SIMD" (2008)

// ClarbergToSphere maps uniform 2D coordinates [0,1]x[0,1] to a unit sphere direction.
// Provides equal-area mapping with low distortion using octahedral mapping.
func ClarbergToSphere(u, v float32) Vec3 {
	// Map [0,1]x[0,1] to [-1,1]x[-1,1]
	u = 2*u - 1
	v = 2*v - 1

	var x, y, z float32

	// Octahedral mapping
	z = 1 - math.Abs(u) - math.Abs(v)

	if z >= 0 {
		// z positive hemisphere
		x = u
		y = v
	} else {
		// z negative hemisphere - fold
		x = (1 - math.Abs(v)) * sign(u)
		y = (1 - math.Abs(u)) * sign(v)
	}

	// Normalize to unit sphere
	result := Vec3{X: x, Y: y, Z: z}
	result.Normalize()
	return result
}

type Tile struct{ L, I, J int }

func TileCenterUV(t Tile) (u, v float32) {
	n := float32(int(1) << t.L)
	return (float32(t.I) + 0.5) / n, (float32(t.J) + 0.5) / n
}

// sign returns -1 for negative, +1 for non-negative
func sign(x float32) float32 {
	if x < 0 {
		return -1
	}
	return 1
}
