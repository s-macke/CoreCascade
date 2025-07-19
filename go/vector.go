package main

type Vec2 struct {
	X, Y float64
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
