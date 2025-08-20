package scene

import (
	"CoreCascade/primitives"
	"CoreCascade/scene/signed_distance"
)

func NewSceneAbsorption(time float64) *Scene {
	s := &Scene{}
	s.objects = []sdObject{
		//&signed_distance.Circle{Center: primitives.Vec2{X: 0., Y: 0.}, Radius: 0.05, Material: primitives.NewEmissiveMaterial(10., 10., 10.)},
		&signed_distance.Box{Center: primitives.Vec2{X: 0., Y: -0.25}, HalfSize: primitives.Vec2{X: 0.05, Y: 0.125}, Material: primitives.NewEmissiveMaterial(0., 0., 10.)},
		&signed_distance.Box{Center: primitives.Vec2{X: 0., Y: 0.0}, HalfSize: primitives.Vec2{X: 0.05, Y: 0.125}, Material: primitives.NewEmissiveMaterial(10., 0, 0)},
		&signed_distance.Box{Center: primitives.Vec2{X: 0., Y: 0.25}, HalfSize: primitives.Vec2{X: 0.05, Y: 0.125}, Material: primitives.NewEmissiveMaterial(0., 10., 0.)},
		&signed_distance.Circle{Center: primitives.Vec2{X: 0.5 - time, Y: 0.}, Radius: 0.3, Material: primitives.NewAbsorbiveMaterial(5.)},
	}
	return s
}
