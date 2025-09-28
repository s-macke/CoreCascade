package scenes

import (
	"CoreCascade2D/primitives"
	"CoreCascade2D/scene/sdf"
	"CoreCascade2D/scene/sdf/signed_distance"
	"vector"
)

func NewScenePenumbra() *sdf.Scene {
	s := &sdf.Scene{}
	s.Objects = []sdf.SdObject{
		&signed_distance.Box{Center: vector.Vec2{X: -1., Y: 0.}, HalfSize: vector.Vec2{X: 0.05, Y: 0.5}, Material: primitives.NewEmissiveMaterial(3., 3., 3.)},
		&signed_distance.Box{Center: vector.Vec2{X: -0.3, Y: 0.60}, HalfSize: vector.Vec2{X: 0.02, Y: 0.6}, Material: primitives.NewAbsorbiveMaterial(500, 0, 0, 0)},
	}
	return s
}
