package scenes

import (
	"CoreCascade3D/primitives"
	"CoreCascade3D/scene/grid"
	"color"
	"fmt"
	math "github.com/chewxy/math32"
	"image"
	"os"
)

func NewSceneFluidHeight(time float32) *grid.Scene {

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

	s := grid.NewScene(img.Bounds().Dx(), img.Bounds().Dy(), 40)

	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {

			for z := 0; z < s.Depth; z++ {
				s.M[y][x][z].Material = primitives.VoidMaterial
			}

			r, g, _, _ := img.At(x, y).RGBA()
			densityA := float32(r)
			densityB := float32(g)

			//densityA := float32(r) / float32(idx)
			//densityB := float32(g) / float32(idx)
			//densityA := max(float32(r)-float32(idx)*5000, 0.)
			//densityB := max(float32(g)-float32(idx)*5000, 0.)

			/*
				densityA = min(densityA/8., 1000.)
				densityB = min(densityB/8., 1000.)
			*/
			/*
				if densityA >= float32(idx)*10000. && densityA < float32(idx+1)*10000. {
					densityA = densityA - float32(idx)*10000.
				} else {
					densityA = 0.
				}
				if densityB >= float32(idx)*10000. && densityB < float32(idx+1)*10000. {
					densityB = densityB - float32(idx)*10000.
				} else {
					densityB = 0.
				}
			*/
			//densityA = max(densityA-float32(idx)*2000., 0.)
			//densityB = max(densityB-float32(idx)*2000., 0.)

			// x*x + y*y + z*z = r*r
			// radius + z*z = r*r
			/*
				radius := (x-64)*(x-64) + (y-64)*(y-64)
				//if radius < 30*30 {
				//radius = radius / (30. * 30)
				//absorption := max(1.-float32(radius), 0.)
				absorption := math.Sqrt(max(30.*30.-float32(radius), 0.)) / 30.
				//absorption = max(absorption-float32(idx)*0.3, 0.)
				absorption = max(absorption/float32(idx*2), 0.)
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
				L: 0.1,
				A: 0.4 * math.Cos(0.+0.9),
				B: 0.4 * math.Sin(0.+0.9),
			}
			c1 := lab1.ToLinear()

			lab2 := color.Oklab{
				L: 0.1,
				A: 0.4 * math.Cos(math.Pi+0.9),
				B: 0.4 * math.Sin(math.Pi+0.9),
			}
			c2 := lab2.ToLinear()

			height1 := min(int(densityA/1000.), s.Depth-2) + 1
			height2 := min(int(densityB/1000.), s.Depth-2) + 1
			/*
				if height1 == s.Depth-1 {
					fmt.Println(height1, height2, c1, c2)
				}
			*/
			for z := 1; z < height1; z++ {
				s.M[y][x][z>>2].Material.Absorption += 2.
				s.M[y][x][z>>2].Material.Diffuse.R += c1.R * 0.25
				s.M[y][x][z>>2].Material.Diffuse.G += c1.G * 0.25
				s.M[y][x][z>>2].Material.Diffuse.B += c1.B * 0.25
			}
			for z := 1; z < height2; z++ {
				s.M[y][x][z>>2].Material.Absorption += 2.
				s.M[y][x][z>>2].Material.Diffuse.R += c2.R * 0.25
				s.M[y][x][z>>2].Material.Diffuse.G += c2.G * 0.25
				s.M[y][x][z>>2].Material.Diffuse.B += c2.B * 0.25
			}

			s.M[y][x][0].Material.Absorption = 50.
			s.M[y][x][0].Material.Diffuse.R = 0.1
			s.M[y][x][0].Material.Diffuse.G = 0.1
			s.M[y][x][0].Material.Diffuse.B = 0.1

			/*
				s.M[y][x].Material.Absorption = (densityA + densityB) / 50000.
				s.M[y][x].Material.Diffuse.R += c1.R * 0.000001 * densityA
				s.M[y][x].Material.Diffuse.G += c1.G * 0.000001 * densityA
				s.M[y][x].Material.Diffuse.B += c1.B * 0.000001 * densityA
				s.M[y][x].Material.Diffuse.R += c2.R * 0.000001 * densityB
				s.M[y][x].Material.Diffuse.G += c2.G * 0.000001 * densityB
				s.M[y][x].Material.Diffuse.B += c2.B * 0.000001 * densityB
				s.M[y][x].Material.Low = 0.0
				s.M[y][x].Material.High = 0.05
			*/
			/*
				s.M[y][x].Material.Absorption = 5.0
				s.M[y][x].Material.Diffuse.R = 1.
			*/

		}
	}

	return s
}
