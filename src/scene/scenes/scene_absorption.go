package scenes

import (
	"CoreCascade/primitives"
	"CoreCascade/scene/sdf"
	"CoreCascade/scene/sdf/signed_distance"
)

func NewSceneAbsorption(time float64) *sdf.Scene {
	s := &sdf.Scene{}
	s.Objects = []sdf.SdObject{
		&signed_distance.Box{Center: primitives.Vec2{X: 0., Y: -0.25}, HalfSize: primitives.Vec2{X: 0.05, Y: 0.125}, Material: primitives.NewEmissiveMaterial(0., 0., 10.)},
		&signed_distance.Box{Center: primitives.Vec2{X: 0., Y: 0.0}, HalfSize: primitives.Vec2{X: 0.05, Y: 0.125}, Material: primitives.NewEmissiveMaterial(10., 0, 0)},
		&signed_distance.Box{Center: primitives.Vec2{X: 0., Y: 0.25}, HalfSize: primitives.Vec2{X: 0.05, Y: 0.125}, Material: primitives.NewEmissiveMaterial(0., 10., 0.)},
		&signed_distance.Circle{Center: primitives.Vec2{X: 0.5 - time, Y: 0.}, Radius: 0.3, Material: primitives.NewAbsorbiveMaterial(5., 3.0, 3.0, 3.0)},
	}
	return s
}
