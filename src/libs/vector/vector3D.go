package vector

import (
	"math"
	"math/rand/v2"
)

type Vec3 struct {
	X, Y, Z float64
}

// RandomUnitVec3 returns a random unit vector uniformly distributed on the sphere
func NewRandomUnitVec3() Vec3 {
	// Generate random angle (azimuthal) and z-value
	z := rand.Float64()*2 - 1 // uniform in [-1,1]
	theta := rand.Float64() * 2 * math.Pi

	r := math.Sqrt(1 - z*z)
	x := r * math.Cos(theta)
	y := r * math.Sin(theta)

	return Vec3{X: x, Y: y, Z: z}
}

func (v *Vec3) ToVec2() Vec2 {
	return Vec2{
		v.X,
		v.Y,
	}
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vec3) Normalize() float64 {
	l := v.Length()
	v.X /= l
	v.Y /= l
	v.Z /= l
	return l

}

func (v *Vec3) Add(w Vec3) {
	v.X += w.X
	v.Y += w.Y
	v.Z += w.Z
}

func (v *Vec3) Sub(w Vec3) {
	v.X -= w.X
	v.Y -= w.Y
	v.Z -= w.Z
}

type Ray3D struct {
	P   Vec3
	Dir Vec3
}

func (r *Ray3D) Trace(t float64) Vec3 {
	// Move the ray's point along its direction vector by distance d
	return Vec3{
		X: r.P.X + r.Dir.X*t,
		Y: r.P.Y + r.Dir.Y*t,
		Z: r.P.Z + r.Dir.Z*t,
	}
}
