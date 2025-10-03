package vector

import math "github.com/chewxy/math32"

type Vec2 struct {
	X, Y float32
}

func NewVec2fromAngle(angle float32) Vec2 {
	return Vec2{
		X: math.Cos(angle),
		Y: math.Sin(angle),
	}
}

func (v *Vec2) Length() float32 {
	return math.Hypot(v.X, v.Y)
}

func (v *Vec2) Normalize() float32 {
	l := v.Length()
	v.X /= l
	v.Y /= l
	return l

}

func (v *Vec2) Add(w Vec2) {
	v.X += w.X
	v.Y += w.Y
}

func (v *Vec2) Sub(w Vec2) {
	v.X -= w.X
	v.Y -= w.Y
}

type Ray2D struct {
	P   Vec2
	Dir Vec2
}

func (r *Ray2D) Trace(t float32) Vec2 {
	// Move the ray's point along its direction vector by distance d
	return Vec2{
		r.P.X + r.Dir.X*t,
		r.P.Y + r.Dir.Y*t,
	}
}
