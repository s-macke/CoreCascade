package scene

import (
	"CoreCascade/primitives"
	"CoreCascade/scene/signed_distance"
)

func NewSceneBeam() *Scene {
	s := &Scene{}
	s.objects = []sdObject{
		&signed_distance.Box{Center: primitives.Vec2{X: -1., Y: 0.}, HalfSize: primitives.Vec2{X: 0.1, Y: 0.25}, Material: primitives.NewEmissiveMaterial(5., 5., 5.)},
		&signed_distance.Box{Center: primitives.Vec2{X: -0.3, Y: 1.00}, HalfSize: primitives.Vec2{X: 0.5, Y: 0.95}, Material: primitives.NewBlackMaterial()},
		&signed_distance.Box{Center: primitives.Vec2{X: -0.3, Y: -1.00}, HalfSize: primitives.Vec2{X: 0.5, Y: 0.95}, Material: primitives.NewBlackMaterial()},
		&signed_distance.Circle{Center: primitives.Vec2{X: 1., Y: 1.}, Radius: 0.05, Material: primitives.NewEmissiveMaterial(0.1, 0, 0.)},
		&signed_distance.Circle{Center: primitives.Vec2{X: 1., Y: -1.}, Radius: 0.05, Material: primitives.NewEmissiveMaterial(0.1, 0, 0.)},
	}
	return s
}
