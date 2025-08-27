package scenes

import (
	"CoreCascade/primitives"
	"CoreCascade/scene/sdf"
	"CoreCascade/scene/sdf/signed_distance"
)

func NewScenePenumbra() *sdf.Scene {
	s := &sdf.Scene{}
	s.Objects = []sdf.SdObject{
		&signed_distance.Box{Center: primitives.Vec2{X: -1., Y: 0.}, HalfSize: primitives.Vec2{X: 0.05, Y: 0.5}, Material: primitives.NewEmissiveMaterial(3., 3., 3.)},
		&signed_distance.Box{Center: primitives.Vec2{X: -0.3, Y: 0.60}, HalfSize: primitives.Vec2{X: 0.02, Y: 0.6}, Material: primitives.NewAbsorbiveMaterial(500, 0, 0, 0)},
	}
	return s
}
