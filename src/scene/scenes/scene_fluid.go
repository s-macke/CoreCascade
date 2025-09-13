package scenes

import (
	"CoreCascade/scene/grid"
	"fmt"
	"image"
	"os"
)

func NewSceneFluid(time float64, idx int) *grid.Scene {

	infile, err := os.Open(fmt.Sprintf("src/scene/scenes/images/img%05d.png", int(time)))
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
			r, g, _, _ := img.At(x, y).RGBA()
			s.M[y][x].Material.Diffuse.R = 0.
			s.M[y][x].Material.Diffuse.G = 0.
			s.M[y][x].Material.Diffuse.B = 0.
			s.M[y][x].Material.Emissive.R = 0.
			s.M[y][x].Material.Emissive.G = 0.
			s.M[y][x].Material.Emissive.B = 0.
			s.M[y][x].Material.Absorption = 0.

			densityA := float64(r)
			densityB := float64(g)
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

			/*
				s.M[y][x].Material.Absorption = (densityA + densityB) / 5000.
				s.M[y][x].Material.Diffuse.R += 0.001 * float64(densityA)
				s.M[y][x].Material.Diffuse.G += 0.001 * float64(densityB)
			*/

			s.M[y][x].Material.Absorption = (densityA + densityB) / 500.
			s.M[y][x].Material.Emissive.R += 0.0002 * float64(densityA)
			s.M[y][x].Material.Emissive.G += 0.0002 * float64(densityA)
			s.M[y][x].Material.Emissive.B += 0.0002 * float64(densityA)
			//s.M[y][x].Material.Emissive.G += 0.0001 * float64(densityB)

			/*
				s.M[y][x].Material.Absorption = 3.
				s.M[y][x].Material.Diffuse.R = 100
				s.M[y][x].Material.Diffuse.G = 0.
			*/
		}
	}
	/*
		for x := 60; x < img.Bounds().Dx()-60; x++ {
			s.M[2][x].Material.Emissive.R = 50.
			s.M[2][x].Material.Emissive.G = 50.
			s.M[2][x].Material.Emissive.B = 50.
			s.M[2][x].Material.Absorption = 0.
			s.M[3][x].Material.Emissive.R = 50.
			s.M[3][x].Material.Emissive.G = 50.
			s.M[3][x].Material.Emissive.B = 50.
			s.M[3][x].Material.Absorption = 0.
		}
	*/
	return s
}
