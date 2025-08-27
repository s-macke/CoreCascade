package scenes

import (
	"CoreCascade/primitives"
	"CoreCascade/scene/sdf"
	"CoreCascade/scene/sdf/signed_distance"
)

func NewSceneBeam() *sdf.Scene {
	s := &sdf.Scene{}
	s.Objects = []sdf.SdObject{
		&signed_distance.Box{Center: primitives.Vec2{X: -1., Y: 0.}, HalfSize: primitives.Vec2{X: 0.1, Y: 0.25}, Material: primitives.NewEmissiveMaterial(5., 5., 5.)},
		&signed_distance.Box{Center: primitives.Vec2{X: -0.3, Y: 1.00}, HalfSize: primitives.Vec2{X: 0.5, Y: 0.95}, Material: primitives.NewAbsorbiveMaterial(500, 0, 0, 0)},
		&signed_distance.Box{Center: primitives.Vec2{X: -0.3, Y: -1.00}, HalfSize: primitives.Vec2{X: 0.5, Y: 0.95}, Material: primitives.NewAbsorbiveMaterial(500, 0, 0, 0)},
		&signed_distance.Circle{Center: primitives.Vec2{X: 1., Y: 1.}, Radius: 0.05, Material: primitives.NewEmissiveMaterial(0.1, 0, 0.)},
		&signed_distance.Circle{Center: primitives.Vec2{X: 1., Y: -1.}, Radius: 0.05, Material: primitives.NewEmissiveMaterial(0.1, 0, 0.)},
	}
	return s
}
