package scenes

import (
	"CoreCascade2D/primitives"
	"CoreCascade2D/scene/sdf"
	"CoreCascade2D/scene/sdf/signed_distance"
	"vector"
)

func NewScenePinhole() *sdf.Scene {
	s := &sdf.Scene{}
	scale := 2.
	s.Objects = []sdf.SdObject{
		&signed_distance.Box{Center: vector.Vec2{X: -0.9, Y: -1.00}, HalfSize: vector.Vec2{X: 0.02, Y: 0.125}, Material: primitives.NewEmissiveSRGBMaterial(1.*scale, 0.*scale, 1.*scale)},
		&signed_distance.Box{Center: vector.Vec2{X: -0.9, Y: -0.75}, HalfSize: vector.Vec2{X: 0.02, Y: 0.125}, Material: primitives.NewEmissiveSRGBMaterial(1.*scale, 1.*scale, 1.*scale)},
		&signed_distance.Box{Center: vector.Vec2{X: -0.9, Y: -0.5}, HalfSize: vector.Vec2{X: 0.02, Y: 0.125}, Material: primitives.NewEmissiveSRGBMaterial(1.*scale, 1.*scale, 0.*scale)},
		&signed_distance.Box{Center: vector.Vec2{X: -0.9, Y: -0.25}, HalfSize: vector.Vec2{X: 0.02, Y: 0.125}, Material: primitives.NewEmissiveSRGBMaterial(0.*scale, 0.*scale, 1.*scale)},
		&signed_distance.Box{Center: vector.Vec2{X: -0.9, Y: 0.0}, HalfSize: vector.Vec2{X: 0.02, Y: 0.125}, Material: primitives.NewEmissiveSRGBMaterial(1.*scale, 0.5*scale, 0.5*scale)},
		&signed_distance.Box{Center: vector.Vec2{X: -0.9, Y: 0.25}, HalfSize: vector.Vec2{X: 0.02, Y: 0.125}, Material: primitives.NewEmissiveSRGBMaterial(0.*scale, 1.*scale, 0.*scale)},
		&signed_distance.Box{Center: vector.Vec2{X: -0.9, Y: 0.5}, HalfSize: vector.Vec2{X: 0.02, Y: 0.125}, Material: primitives.NewEmissiveSRGBMaterial(0.5*scale, 0.5*scale, 1.*scale)},
		&signed_distance.Box{Center: vector.Vec2{X: -0.9, Y: 0.75}, HalfSize: vector.Vec2{X: 0.02, Y: 0.125}, Material: primitives.NewEmissiveSRGBMaterial(1.*scale, 0.*scale, 0.*scale)},
		&signed_distance.Box{Center: vector.Vec2{X: -0.9, Y: 1.00}, HalfSize: vector.Vec2{X: 0.02, Y: 0.125}, Material: primitives.NewEmissiveSRGBMaterial(0.5*scale, 1.*scale, 0.5*scale)},
		&signed_distance.Box{Center: vector.Vec2{X: -0.9, Y: 0.00}, HalfSize: vector.Vec2{X: 0.02, Y: 1.}, Material: primitives.NewAbsorbiveMaterial(100, 0, 0, 0)},

		&signed_distance.Box{Center: vector.Vec2{X: -0., Y: -3.6}, HalfSize: vector.Vec2{X: 0.02, Y: 3.55}, Material: primitives.NewAbsorbiveMaterial(500, 0, 0, 0)},
		&signed_distance.Box{Center: vector.Vec2{X: -0., Y: 3.6}, HalfSize: vector.Vec2{X: 0.02, Y: 3.55}, Material: primitives.NewAbsorbiveMaterial(500, 0, 0, 0)},
	}
	return s
}
