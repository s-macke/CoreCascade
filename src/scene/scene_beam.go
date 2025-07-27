package scene

import (
	"CoreCascade/primitives"
	"CoreCascade/scene/signed_distance"
)

func NewSceneBeam() *Scene {
	s := &Scene{}
	s.objects = []sdObject{
		&signed_distance.Box{Center: primitives.Vec2{X: -1., Y: 0.}, HalfSize: primitives.Vec2{X: 0.1, Y: 0.25}, Color: primitives.Color{R: 5., G: 5., B: 5.}},
		&signed_distance.Box{Center: primitives.Vec2{X: -0.3, Y: 1.00}, HalfSize: primitives.Vec2{X: 0.5, Y: 0.95}, Color: primitives.Color{R: 0., G: 0, B: 0}},
		&signed_distance.Box{Center: primitives.Vec2{X: -0.3, Y: -1.00}, HalfSize: primitives.Vec2{X: 0.5, Y: 0.95}, Color: primitives.Color{R: 0., G: 0, B: 0}},
		&signed_distance.Circle{Center: primitives.Vec2{X: 1., Y: 1.}, Radius: 0.05, Color: primitives.Color{R: 0.1, G: 0, B: 0.}},
		&signed_distance.Circle{Center: primitives.Vec2{X: 1., Y: -1.}, Radius: 0.05, Color: primitives.Color{R: 0.1, G: 0, B: 0.}},
	}
	return s
}
