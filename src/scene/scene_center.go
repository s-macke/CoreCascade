package scene

import (
	"CoreCascade/primitives"
	"CoreCascade/scene/signed_distance"
)

func NewSceneCenter() *Scene {
	s := &Scene{}
	s.objects = []sdObject{
		&signed_distance.Circle{Center: primitives.Vec2{X: 0., Y: -0.}, Radius: 0.05, Color: primitives.Color{R: 1, G: 1, B: 1.}},
	}
	return s
}
