package scenes

import (
	"CoreCascade2D/primitives"
	"CoreCascade2D/scene/sdf"
	"CoreCascade2D/scene/sdf/signed_distance"
	math "github.com/chewxy/math32"
	"vector"
)

func NewSceneShadows(time float32) *sdf.Scene {
	s := &sdf.Scene{}
	s.Objects = []sdf.SdObject{
		&signed_distance.Circle{Center: vector.Vec2{X: 0., Y: 0.}, Radius: 0.05, Material: primitives.NewEmissiveMaterial(10, 10, 10)},
	}

	for i := 0; i < 10; i++ {
		x := math.Cos((float32(i)+time/50.)*2.*math.Pi/10.) * 0.5
		y := math.Sin((float32(i)+time/50.)*2.*math.Pi/10.) * 0.5
		s.Objects = append(s.Objects, &signed_distance.Circle{Center: vector.Vec2{X: x, Y: y}, Radius: float32(i)/100. + 0.01, Material: primitives.NewAbsorbiveMaterial(1000, 0, 0, 0)})
	}
	return s
}
