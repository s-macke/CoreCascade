package scene

import "CoreCascade/primitives"

func NewScenePenumbra() *Scene {
	s := &Scene{}
	s.objects = []sdObject{
		&Box{Center: primitives.Vec2{X: -1., Y: 0.}, HalfSize: primitives.Vec2{X: 0.05, Y: 0.5}, Color: primitives.Color{R: 1., G: 1., B: 1.}},
		&Box{Center: primitives.Vec2{X: -0.3, Y: 0.60}, HalfSize: primitives.Vec2{X: 0.02, Y: 0.5}, Color: primitives.Color{R: 0., G: 0, B: 0}},
	}
	return s
}
