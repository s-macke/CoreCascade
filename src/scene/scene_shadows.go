package scene

import (
	"CoreCascade/primitives"
	"math"
)

func NewSceneShadows() *Scene {
	s := &Scene{}
	s.objects = []sdObject{
		&Circle{Center: primitives.Vec2{X: 0., Y: 0.}, Radius: 0.05, Color: primitives.Color{R: 1., G: 1, B: 1.}},
	}

	for i := 0; i < 10; i++ {
		x := math.Cos(float64(i)*2.*math.Pi/10.) * 0.5
		y := math.Sin(float64(i)*2.*math.Pi/10.) * 0.5
		s.objects = append(s.objects, &Circle{Center: primitives.Vec2{X: x, Y: y}, Radius: float64(i)/100. + 0.01, Color: primitives.Color{R: 0., G: 0, B: 0.}})
	}
	return s
}
