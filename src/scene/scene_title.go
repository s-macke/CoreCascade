package scene

import (
	"CoreCascade/primitives"
	"CoreCascade/scene/signed_distance"
	"fmt"
)

func NewSceneTitle(time float64) *Scene {

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
	s := &Scene{}
	s.objects = []sdObject{
		&signed_distance.Box{Center: primitives.Vec2{X: 0., Y: 0.5}, HalfSize: primitives.Vec2{X: 0.5, Y: 0.02}, Material: primitives.NewEmissiveMaterial(0.01, 0.01, 0.04)},
		&signed_distance.Box{Center: primitives.Vec2{X: 0., Y: -0.5}, HalfSize: primitives.Vec2{X: 0.3, Y: 0.02}, Material: primitives.NewEmissiveMaterial(0.03, 0.02, 0.01)},
	}
	/*
		for i := 0; i < 20; i++ {
			x := -1.0 + float64(i)*0.1
			y := 0.
			s.objects = append(s.objects, &signed_distance.Circle{
				Center: primitives.Vec2{X: x, Y: y},
				Radius: 0.01,
				Color:  primitives.Color{R: 0.0, G: 0.0, B: 0.0},
			})
		}
	*/
	l := 0.03
	x := -0.5
	y := -0.9 + 0.56
	for _, c := range title[0] {
		x += l

		switch c {
		case ' ':
			// Skip spaces
		case '#':
			s.objects = append(s.objects, &signed_distance.Circle{
				Center:   primitives.Vec2{X: x, Y: y},
				Radius:   0.007,
				Material: primitives.NewEmissiveMaterial(0.3, 0.2, 0.1),
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
			s.objects = append(s.objects, &signed_distance.Circle{
				Center:   primitives.Vec2{X: x, Y: y},
				Radius:   0.007,
				Material: primitives.NewEmissiveMaterial(0.15, 0.15, 0.6),
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
