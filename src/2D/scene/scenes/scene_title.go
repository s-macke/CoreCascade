package scenes

import (
	"CoreCascade2D/primitives"
	"CoreCascade2D/scene/sdf"
	"CoreCascade2D/scene/sdf/signed_distance"
	"color"
	"fmt"
	"vector"
)

func NewSceneTitle(time float32) *sdf.Scene {

	title := []string{
		`
 #####                       
#     #   ####   #####   ###### 
#        #    #  #    #  #      
#        #    #  #    #  #####  
#        #    #  #####   #      
#     #  #    #  #   #   #      
 #####    ####   #    #  ######
`,
		`
 #####
#     #    ##     ####    ####     ##    #####   ######
#         #  #   #       #    #   #  #   #    #  #
#        #    #   ####   #       #    #  #    #  #####
#        ######       #  #       ######  #    #  #
#     #  #    #  #    #  #    #  #    #  #    #  #
 #####   #    #   ####    ####   #    #  #####   ######
`,
	}

	s := &sdf.Scene{}
	s.Objects = []sdf.SdObject{
		&signed_distance.Box{Center: vector.Vec2{X: 0., Y: 0.5}, HalfSize: vector.Vec2{X: 0.5, Y: 0.02}, Material: primitives.NewAbsorbiveMaterial(20., 0., 0., 0.)},
		&signed_distance.Box{Center: vector.Vec2{X: 0., Y: -0.5}, HalfSize: vector.Vec2{X: 0.3, Y: 0.02}, Material: primitives.NewAbsorbiveMaterial(20., 0., 0., 0.)},
	}

	l := float32(0.03)
	x := float32(-0.5)
	y := float32(-0.9 + 0.56)
	for _, c := range title[0] {
		x += l

		switch c {
		case ' ':
			// Skip spaces
		case '#':
			col := color.NewRainbowOklabToLinear((x + 0.5) * 0.5)
			col.Mul(20.)
			s.Objects = append(s.Objects, &signed_distance.Circle{
				Center:   vector.Vec2{X: x, Y: y},
				Radius:   0.007,
				Material: primitives.NewEmissiveMaterial(col.R, col.G, col.B),
			})
		case '\n':
			x = -0.5
			y += l
		default:
			fmt.Println("Unknown character in title:", string(c))
		}
	}

	l = 0.03
	x = -0.85
	y = -0.5 + 0.56
	for _, c := range title[1] {
		x += l

		switch c {
		case ' ':
			// Skip spaces
		case '#':
			col := color.NewRainbowOklabToLinear(1. - (x+0.5)/2.)
			col.Mul(20.)
			s.Objects = append(s.Objects, &signed_distance.Circle{
				Center:   vector.Vec2{X: x, Y: y},
				Radius:   0.007,
				Material: primitives.NewEmissiveMaterial(col.R, col.G, col.B),
			})
		case '\n':
			x = -0.85
			y += l
		default:
			fmt.Println("Unknown character in title:", string(c))
		}
	}

	return s
}
