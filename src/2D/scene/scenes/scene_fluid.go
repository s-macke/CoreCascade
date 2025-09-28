package scenes

import (
	"CoreCascade2D/primitives"
	"CoreCascade2D/scene/grid"
	"color"
	"fmt"
	"image"
	"math"
	"os"
)

func NewSceneFluid(time float64, idx int) *grid.Scene {

	infile, err := os.Open(fmt.Sprintf("src/scene/scenes/images/img%05d.png", int(time)))
	//infile, err := os.Open(fmt.Sprintf("src/scene/scenes/images/dam%05d.png", int(time)))
	if err != nil {
		panic(err)
	}
	defer infile.Close()

	img, _, err := image.Decode(infile)
	if err != nil {
		panic(err)
	}

	s := grid.NewScene(img.Bounds().Dx(), img.Bounds().Dy())

	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			s.M[y][x].Material = primitives.VoidMaterial

			r, g, _, _ := img.At(x, y).RGBA()
			densityA := float64(r)
			densityB := float64(g)

			//densityA := float64(r) / float64(idx)
			//densityB := float64(g) / float64(idx)
			//densityA := max(float64(r)-float64(idx)*5000, 0.)
			//densityB := max(float64(g)-float64(idx)*5000, 0.)

			/*
				densityA = min(densityA/8., 1000.)
				densityB = min(densityB/8., 1000.)
			*/
			/*
				if densityA >= float64(idx)*10000. && densityA < float64(idx+1)*10000. {
					densityA = densityA - float64(idx)*10000.
				} else {
					densityA = 0.
				}
				if densityB >= float64(idx)*10000. && densityB < float64(idx+1)*10000. {
					densityB = densityB - float64(idx)*10000.
				} else {
					densityB = 0.
				}
			*/
			//densityA = max(densityA-float64(idx)*2000., 0.)
			//densityB = max(densityB-float64(idx)*2000., 0.)

			// x*x + y*y + z*z = r*r
			// radius + z*z = r*r
			/*
				radius := (x-64)*(x-64) + (y-64)*(y-64)
				//if radius < 30*30 {
				//radius = radius / (30. * 30)
				//absorption := max(1.-float64(radius), 0.)
				absorption := math.Sqrt(max(30.*30.-float64(radius), 0.)) / 30.
				//absorption = max(absorption-float64(idx)*0.3, 0.)
				absorption = max(absorption/float64(idx*2), 0.)
				s.M[y][x].Material.Absorption = absorption * 5. * 5.
				s.M[y][x].Material.Diffuse.R = 2. * absorption * 5.
				//}
			*/
			/*
				if r > 0 {
					s.M[y][x].Material.Absorption = 10.
					s.M[y][x].Material.Diffuse.R += 0.5
					s.M[y][x].Material.Diffuse.G += 0.5
					s.M[y][x].Material.Diffuse.B += 3.
				}
			*/

			lab1 := color.Oklab{
				L: 1.,
				A: 0.4 * math.Cos(0.+0.1),
				B: 0.4 * math.Sin(0.+0.1),
			}
			c1 := lab1.ToLinear()
			lab2 := color.Oklab{
				L: 1.,
				A: 0.4 * math.Cos(math.Pi+0.1),
				B: 0.4 * math.Sin(math.Pi+0.1),
			}
			c2 := lab2.ToLinear()

			s.M[y][x].Material.Absorption = (densityA + densityB) / 5000. * 2.
			s.M[y][x].Material.Diffuse.R += c1.R * 0.00005 * densityA
			s.M[y][x].Material.Diffuse.G += c1.G * 0.00005 * densityA
			s.M[y][x].Material.Diffuse.B += c1.B * 0.00005 * densityA
			s.M[y][x].Material.Diffuse.R += c2.R * 0.00005 * densityB
			s.M[y][x].Material.Diffuse.G += c2.G * 0.00005 * densityB
			s.M[y][x].Material.Diffuse.B += c2.B * 0.00005 * densityB

			/*
				if s.M[y][x].Material.Absorption > 3. {
					fmt.Println(s.M[y][x].Material)
				}
			*/
			/*
				s.M[y][x].Material.Absorption = (densityA + densityB) / 500.
				s.M[y][x].Material.Emissive.R += 0.0002 * float64(densityA)
				s.M[y][x].Material.Emissive.G += 0.0002 * float64(densityA)
				s.M[y][x].Material.Emissive.B += 0.0002 * float64(densityA)
				//s.M[y][x].Material.Emissive.G += 0.0001 * float64(densityB)
			*/
			/*
				s.M[y][x].Material.Absorption = 3.
				s.M[y][x].Material.Diffuse.R = 100
				s.M[y][x].Material.Diffuse.G = 0.
			*/
		}
	}

	//for x := 50; x < img.Bounds().Dx()-50; x++ {
	for x := 0; x < img.Bounds().Dx()-0; x++ {
		s.M[2][x].Material.Emissive.R = 50. / 30.
		s.M[2][x].Material.Emissive.G = 50. / 30.
		s.M[2][x].Material.Emissive.B = 50. / 30.
		s.M[3][x].Material.Emissive.R = 50. / 30.
		s.M[3][x].Material.Emissive.G = 50. / 30.
		s.M[3][x].Material.Emissive.B = 50. / 30.
	}

	return s
}
