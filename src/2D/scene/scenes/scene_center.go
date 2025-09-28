package scenes

import (
	"CoreCascade2D/primitives"
	"CoreCascade2D/scene/sdf"
	"CoreCascade2D/scene/sdf/signed_distance"
	"vector"
)

func NewSceneCenter() *sdf.Scene {
	s := &sdf.Scene{}
	s.Objects = []sdf.SdObject{
		&signed_distance.Circle{Center: vector.Vec2{X: 0., Y: -0.}, Radius: 0.05, Material: primitives.NewEmissiveMaterial(10, 10, 10)},
	}
	return s
}
