package scenes

import (
	"CoreCascade2D/primitives"
	"CoreCascade2D/scene/sdf"
	"CoreCascade2D/scene/sdf/signed_distance"
	"color"
	"vector"
)

func NewScenePinhole() *sdf.Scene {
	s := &sdf.Scene{}
	s.Objects = []sdf.SdObject{
		//&signed_distance.Box{Center: vector.Vec2{X: -0.9, Y: 0.00}, HalfSize: vector.Vec2{X: 0.02, Y: 1.}, Material: primitives.NewAbsorbiveMaterial(100, 0, 0, 0)},
		&signed_distance.Box{Center: vector.Vec2{X: -0., Y: -3.6}, HalfSize: vector.Vec2{X: 0.02, Y: 3.55}, Material: primitives.NewAbsorbiveMaterial(500, 0, 0, 0)},
		&signed_distance.Box{Center: vector.Vec2{X: -0., Y: 3.6}, HalfSize: vector.Vec2{X: 0.02, Y: 3.55}, Material: primitives.NewAbsorbiveMaterial(500, 0, 0, 0)},
	}

	for i := -10; i <= 10; i++ {
		c := color.NewRainbowOklabToLinear((float64(i) + 10) / 20.)
		c.Mul(3.)
		s.Objects = append(s.Objects,
			&signed_distance.Box{
				Center:   vector.Vec2{X: -0.9, Y: -float64(i) * 0.1},
				HalfSize: vector.Vec2{X: 0.02, Y: 0.05},
				Material: primitives.NewEmissiveMaterial(c.R, c.G, c.B),
			})
	}

	return s
}
