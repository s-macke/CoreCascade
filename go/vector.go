package main

import "math"

type Vec2 struct {
	X, Y float64
}

func NewVec2fromAngle(angle float64) Vec2 {
	return Vec2{
		X: math.Cos(angle),
		Y: math.Sin(angle),
	}
}

func (v *Vec2) Length() float64 {
	return math.Hypot(v.X, v.Y)
}

func (v *Vec2) Add(w Vec2) {
	v.X += w.X
	v.Y += w.Y
}

func (v *Vec2) Sub(w Vec2) {
	v.X -= w.X
	v.Y -= w.Y
}

type Ray struct {
	p   Vec2
	dir Vec2
}

func (r *Ray) Trace(t float64) Vec2 {
	// Move the ray's point along its direction vector by distance d
	return Vec2{
		r.p.X + r.dir.X*t,
		r.p.Y + r.dir.Y*t,
	}
}
