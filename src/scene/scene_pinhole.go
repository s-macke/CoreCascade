package scene

import "CoreCascade/primitives"

func NewScenePinhole() *Scene {
	s := &Scene{}
	s.objects = []sdObject{
		&Box{Center: primitives.Vec2{X: -1., Y: -1.00}, HalfSize: primitives.Vec2{X: 0.1, Y: 0.125}, Color: primitives.NewSRGBColor(1., 0., 1.)},
		&Box{Center: primitives.Vec2{X: -1., Y: -0.75}, HalfSize: primitives.Vec2{X: 0.1, Y: 0.125}, Color: primitives.NewSRGBColor(1., 1., 1.)},
		&Box{Center: primitives.Vec2{X: -1., Y: -0.5}, HalfSize: primitives.Vec2{X: 0.1, Y: 0.125}, Color: primitives.NewSRGBColor(1., 1., 0.)},
		&Box{Center: primitives.Vec2{X: -1., Y: -0.25}, HalfSize: primitives.Vec2{X: 0.1, Y: 0.125}, Color: primitives.NewSRGBColor(0., 0., 1.)},
		&Box{Center: primitives.Vec2{X: -1., Y: 0.0}, HalfSize: primitives.Vec2{X: 0.1, Y: 0.125}, Color: primitives.NewSRGBColor(1., 0.5, 0.5)},
		&Box{Center: primitives.Vec2{X: -1., Y: 0.25}, HalfSize: primitives.Vec2{X: 0.1, Y: 0.125}, Color: primitives.NewSRGBColor(0., 1., 0.)},
		&Box{Center: primitives.Vec2{X: -1., Y: 0.5}, HalfSize: primitives.Vec2{X: 0.1, Y: 0.125}, Color: primitives.NewSRGBColor(0.5, 0.5, 1.)},
		&Box{Center: primitives.Vec2{X: -1., Y: 0.75}, HalfSize: primitives.Vec2{X: 0.1, Y: 0.125}, Color: primitives.NewSRGBColor(1., 0., 0.)},
		&Box{Center: primitives.Vec2{X: -1., Y: 1.00}, HalfSize: primitives.Vec2{X: 0.1, Y: 0.125}, Color: primitives.NewSRGBColor(0.5, 1., 0.5)},

		&Box{Center: primitives.Vec2{X: -0., Y: -3.6}, HalfSize: primitives.Vec2{X: 0.02, Y: 3.55}, Color: primitives.Color{R: 0., G: 0, B: 0.0}},
		&Box{Center: primitives.Vec2{X: -0., Y: 3.6}, HalfSize: primitives.Vec2{X: 0.02, Y: 3.55}, Color: primitives.Color{R: 0., G: 0, B: 0.0}},
	}
	return s
}
