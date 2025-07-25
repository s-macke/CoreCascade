package scene

import "CoreCascade/primitives"

func NewSceneBeam() *Scene {
	s := &Scene{}
	s.objects = []sdObject{
		&Box{Center: primitives.Vec2{X: -1., Y: 0.}, HalfSize: primitives.Vec2{X: 0.1, Y: 0.25}, Color: primitives.Color{R: 2., G: 2., B: 2.}},
		&Box{Center: primitives.Vec2{X: -0.3, Y: 1.00}, HalfSize: primitives.Vec2{X: 0.5, Y: 0.95}, Color: primitives.Color{R: 0., G: 0, B: 0}},
		&Box{Center: primitives.Vec2{X: -0.3, Y: -1.00}, HalfSize: primitives.Vec2{X: 0.5, Y: 0.95}, Color: primitives.Color{R: 0., G: 0, B: 0}},
		&Circle{Center: primitives.Vec2{X: 1., Y: 1.}, Radius: 0.05, Color: primitives.Color{R: 0.1, G: 0, B: 0.}},
		&Circle{Center: primitives.Vec2{X: 1., Y: -1.}, Radius: 0.05, Color: primitives.Color{R: 0.1, G: 0, B: 0.}},
	}
	return s
}
