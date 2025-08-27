package scenes

import (
	"CoreCascade/primitives"
	"CoreCascade/scene/sdf"
	"CoreCascade/scene/sdf/signed_distance"
)

func NewSceneCenter() *sdf.Scene {
	s := &sdf.Scene{}
	s.Objects = []sdf.SdObject{
		&signed_distance.Circle{Center: primitives.Vec2{X: 0., Y: -0.}, Radius: 0.05, Material: primitives.NewEmissiveMaterial(10, 10, 10)},
	}
	return s
}
