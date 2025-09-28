package scenes

import (
	"CoreCascade2D/primitives"
	"CoreCascade2D/scene/sdf"
	"CoreCascade2D/scene/sdf/signed_distance"
	"vector"
)

func NewSceneAbsorption(time float64) *sdf.Scene {
	s := &sdf.Scene{}
	s.Objects = []sdf.SdObject{
		&signed_distance.Box{Center: vector.Vec2{X: 0., Y: -0.25}, HalfSize: vector.Vec2{X: 0.05, Y: 0.125}, Material: primitives.NewEmissiveMaterial(0., 0., 10.)},
		&signed_distance.Box{Center: vector.Vec2{X: 0., Y: 0.0}, HalfSize: vector.Vec2{X: 0.05, Y: 0.125}, Material: primitives.NewEmissiveMaterial(10., 0, 0)},
		&signed_distance.Box{Center: vector.Vec2{X: 0., Y: 0.25}, HalfSize: vector.Vec2{X: 0.05, Y: 0.125}, Material: primitives.NewEmissiveMaterial(0., 10., 0.)},
		&signed_distance.Circle{Center: vector.Vec2{X: 0.5 - time, Y: 0.}, Radius: 0.3, Material: primitives.NewAbsorbiveMaterial(5., 0.0, 0.0, 0.0)},
	}
	return s
}
