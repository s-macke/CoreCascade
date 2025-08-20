package scene

import (
	"CoreCascade/primitives"
	"CoreCascade/scene/signed_distance"
)

func NewScenePinhole() *Scene {
	s := &Scene{}
	s.objects = []sdObject{
		&signed_distance.Box{Center: primitives.Vec2{X: -1., Y: -1.00}, HalfSize: primitives.Vec2{X: 0.1, Y: 0.125}, Material: primitives.NewEmissiveSRGBMaterial(1., 0., 1.)},
		&signed_distance.Box{Center: primitives.Vec2{X: -1., Y: -0.75}, HalfSize: primitives.Vec2{X: 0.1, Y: 0.125}, Material: primitives.NewEmissiveSRGBMaterial(1., 1., 1.)},
		&signed_distance.Box{Center: primitives.Vec2{X: -1., Y: -0.5}, HalfSize: primitives.Vec2{X: 0.1, Y: 0.125}, Material: primitives.NewEmissiveSRGBMaterial(1., 1., 0.)},
		&signed_distance.Box{Center: primitives.Vec2{X: -1., Y: -0.25}, HalfSize: primitives.Vec2{X: 0.1, Y: 0.125}, Material: primitives.NewEmissiveSRGBMaterial(0., 0., 1.)},
		&signed_distance.Box{Center: primitives.Vec2{X: -1., Y: 0.0}, HalfSize: primitives.Vec2{X: 0.1, Y: 0.125}, Material: primitives.NewEmissiveSRGBMaterial(1., 0.5, 0.5)},
		&signed_distance.Box{Center: primitives.Vec2{X: -1., Y: 0.25}, HalfSize: primitives.Vec2{X: 0.1, Y: 0.125}, Material: primitives.NewEmissiveSRGBMaterial(0., 1., 0.)},
		&signed_distance.Box{Center: primitives.Vec2{X: -1., Y: 0.5}, HalfSize: primitives.Vec2{X: 0.1, Y: 0.125}, Material: primitives.NewEmissiveSRGBMaterial(0.5, 0.5, 1.)},
		&signed_distance.Box{Center: primitives.Vec2{X: -1., Y: 0.75}, HalfSize: primitives.Vec2{X: 0.1, Y: 0.125}, Material: primitives.NewEmissiveSRGBMaterial(1., 0., 0.)},
		&signed_distance.Box{Center: primitives.Vec2{X: -1., Y: 1.00}, HalfSize: primitives.Vec2{X: 0.1, Y: 0.125}, Material: primitives.NewEmissiveSRGBMaterial(0.5, 1., 0.5)},

		&signed_distance.Box{Center: primitives.Vec2{X: -0., Y: -3.6}, HalfSize: primitives.Vec2{X: 0.02, Y: 3.55}, Material: primitives.NewBlackMaterial()},
		&signed_distance.Box{Center: primitives.Vec2{X: -0., Y: 3.6}, HalfSize: primitives.Vec2{X: 0.02, Y: 3.55}, Material: primitives.NewBlackMaterial()},
	}
	/*
		for i := 0; i < len(s.objects); i++ {
			if obj, ok := s.objects[i].(*signed_distance.Box); ok {
				obj.Color.Mul(1.0)
			}
		}
	*/
	return s
}
