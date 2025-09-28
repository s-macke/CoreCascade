package scenes

import (
	"CoreCascade2D/primitives"
	"CoreCascade2D/scene/sdf"
	"CoreCascade2D/scene/sdf/signed_distance"
	"vector"
)

func NewSceneBeam() *sdf.Scene {
	s := &sdf.Scene{}
	s.Objects = []sdf.SdObject{
		&signed_distance.Box{Center: vector.Vec2{X: -1., Y: 0.}, HalfSize: vector.Vec2{X: 0.1, Y: 0.25}, Material: primitives.NewEmissiveMaterial(5., 5., 5.)},
		&signed_distance.Box{Center: vector.Vec2{X: -0.3, Y: 1.00}, HalfSize: vector.Vec2{X: 0.5, Y: 0.95}, Material: primitives.NewAbsorbiveMaterial(500, 0, 0, 0)},
		&signed_distance.Box{Center: vector.Vec2{X: -0.3, Y: -1.00}, HalfSize: vector.Vec2{X: 0.5, Y: 0.95}, Material: primitives.NewAbsorbiveMaterial(500, 0, 0, 0)},
		&signed_distance.Circle{Center: vector.Vec2{X: 1., Y: 1.}, Radius: 0.05, Material: primitives.NewEmissiveMaterial(0.1, 0, 0.)},
		&signed_distance.Circle{Center: vector.Vec2{X: 1., Y: -1.}, Radius: 0.05, Material: primitives.NewEmissiveMaterial(0.1, 0, 0.)},
	}
	return s
}
