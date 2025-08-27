package scenes

import (
	"CoreCascade/primitives"
	"CoreCascade/scene/sdf"
	"CoreCascade/scene/sdf/signed_distance"
	"math"
)

func NewSceneShadows(time float64) *sdf.Scene {
	s := &sdf.Scene{}
	s.Objects = []sdf.SdObject{
		&signed_distance.Circle{Center: primitives.Vec2{X: 0., Y: 0.}, Radius: 0.05, Material: primitives.NewEmissiveMaterial(10, 10, 10)},
	}

	for i := 0; i < 10; i++ {
		x := math.Cos((float64(i)+time/50.)*2.*math.Pi/10.) * 0.5
		y := math.Sin((float64(i)+time/50.)*2.*math.Pi/10.) * 0.5
		s.Objects = append(s.Objects, &signed_distance.Circle{Center: primitives.Vec2{X: x, Y: y}, Radius: float64(i)/100. + 0.01, Material: primitives.NewAbsorbiveMaterial(1000, 0, 0, 0)})
	}
	return s
}
