package scenes

import (
	"CoreCascade3D/primitives"
	"CoreCascade3D/scene/sdf"
	"CoreCascade3D/scene/sdf/signed_distance"
	"color"
	"fmt"
	"vector"
)

func NewSceneHeight(time float32) *sdf.Scene {
	s := &sdf.Scene{}

	absorption := float32(5.)
	center := vector.Vec2{X: 0., Y: 0.}

	c := color.NewRainbowOklabToLinear(0.)
	m1 := primitives.NewAbsorbiveMaterial(absorption, c.R, c.G, c.B)
	c1 := signed_distance.Circle{Center: center, Radius: 1.7, Material: m1}
	c1.Low = 0.0
	c1.High = 0.01

	c = color.NewRainbowOklabToLinear(0.25)
	m2 := primitives.NewAbsorbiveMaterial(absorption, c.R, c.G, c.B)
	c2 := signed_distance.Circle{Center: center, Radius: 0.7, Material: m2}
	c2.Low = 0.01
	c2.High = 0.02

	c = color.NewRainbowOklabToLinear(0.5)
	m3 := primitives.NewAbsorbiveMaterial(absorption, c.R, c.G, c.B)
	c3 := signed_distance.Circle{Center: center, Radius: 0.3, Material: m3}
	c3.Low = 0.02
	c3.High = 0.04

	c = color.NewRainbowOklabToLinear(0.75)
	m4 := primitives.NewAbsorbiveMaterial(absorption*2, c.R, c.G, c.B)
	c4 := signed_distance.Circle{Center: center, Radius: 0.05, Material: m4}
	c4.Low = 0.04
	c4.High = 0.08

	l := primitives.NewEmissiveMaterial(100., 100., 100.)
	cl := signed_distance.Circle{Center: vector.Vec2{X: 0. + time, Y: -1. + time}, Radius: 0.1, Material: l}
	cl.Low = 0.08
	cl.High = 0.1

	s.Objects = []sdf.SdObject{&c1, &c2, &c3, &c4, &cl}
	for i := range s.Objects {
		o := s.Objects[i].(*signed_distance.Circle)
		if o.Low > o.High {
			fmt.Println(o)
			panic("Circle low value cannot be greater than high value")
		}
	}

	return s
}
