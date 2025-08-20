package scene

import (
	"CoreCascade/primitives"
	"CoreCascade/scene/signed_distance"
)

func NewSceneCenter() *Scene {
	s := &Scene{}
	s.objects = []sdObject{
		&signed_distance.Circle{Center: primitives.Vec2{X: 0., Y: -0.}, Radius: 0.05, Material: primitives.NewEmissiveMaterial(1, 1, 1.)},
	}
	return s
}
