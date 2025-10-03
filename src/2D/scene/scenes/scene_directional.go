package scenes

import (
	"CoreCascade2D/primitives"
	"CoreCascade2D/scene/sdf"
	"CoreCascade2D/scene/sdf/signed_distance"
	"color"
	math "github.com/chewxy/math32"
	"vector"
)

func NewSceneDirectional(time float32) *sdf.Scene {
	s := &sdf.Scene{}
	/*
		m := primitives.NewEmissiveMaterial(10, 10, 10)
		m.EmissiveDirection = primitives.NewVec2fromAngle(time * (2. * math.Pi) / 60.)
		m.EmissiveAngle = math.Cos(math.Pi / 8.)
		m.DirectionEnabled = true
		s.Objects = []sdf.SdObject{
			&signed_distance.Circle{Center: primitives.Vec2{X: 0., Y: -0.}, Radius: 0.05, Material: m},
		}
	*/
	for i := 0; i < 10; i++ {
		angle := float32(i) * (2. * math.Pi) / 10.
		lab := color.Oklab{
			L: 1.,
			A: 0.4 * math.Cos(angle),
			B: 0.4 * math.Sin(angle),
		}
		c := lab.ToLinear()
		c.Mul(5.)
		m := primitives.NewEmissiveMaterial(c.R, c.G, c.B)
		m.Absorption = 3.0
		m.DirectionEnabled = true
		m.EmissiveDirection = vector.NewVec2fromAngle(angle + time*(2.*math.Pi)/60.)
		m.EmissiveAngle = math.Cos(math.Pi / 8.)

		s.Objects = append(s.Objects,
			&signed_distance.Circle{
				Center: vector.Vec2{X: 0.5 * math.Cos(angle), Y: 0.5 * math.Sin(angle)},
				Radius: 0.07,
				//Material: primitives.NewAbsorbiveMaterial(5.0, c.R, c.G, c.B)})
				Material: m})
	}
	return s
}
