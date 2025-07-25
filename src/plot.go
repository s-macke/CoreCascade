package main

import (
	"CoreCascade/primitives"
	"fmt"
	"os"
)

func PlotSignedDistance() {
	scene := NewScene("shadows")
	for x := -2.0; x <= 2.0; x += 0.1 {
		for y := -2.0; y <= 2.0; y += 0.1 {
			d, _ := scene.SignedDistance(primitives.Vec2{X: x, Y: y})
			fmt.Println(x, y, d)
		}
		fmt.Println()
	}

}

func PlotCascade() {
	f, err := os.Create("plots/plot.data")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	const WIDTH, HEIGHT = 4, 4
	cc := NewCascadeCalculator(WIDTH, HEIGHT)
	fmt.Fprintln(f, "# NCascades", cc.NCascades)
	for c := 0; c < 3; c++ {
		ci := cc.cascadeInfo[c]
		fmt.Fprintln(f, "# Cascade", c, "total", ci.Total(), "dirCount", ci.dirCount, "N", ci.N, "tStart", ci.tStart, "tEnd", ci.tEnd)
		for i := 0; i < ci.N; i++ {
			for j := 0; j < ci.M; j++ {
				for k := 0; k < ci.dirCount; k++ {
					probe := cc.GetProbe(c, i, j, k)
					fmt.Fprintf(f, "%f %f %f %f\n", probe.ray.P.X, probe.ray.P.Y, probe.ray.Dir.X*probe.tmax, probe.ray.Dir.Y*probe.tmax)
				}
			}
		}
		fmt.Fprintln(f)
	}
}

func PlotCascade2() {
	f, err := os.Create("plots/plot2.data")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	const WIDTH, HEIGHT = 800, 800
	cc := NewCascadeCalculator(WIDTH, HEIGHT)

	k := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			probe := cc.GetProbe(0, i, j, k)
			fmt.Fprintf(f, "%f %f %f %f\n", probe.ray.P.X, probe.ray.P.Y, probe.ray.Dir.X*probe.tmax, probe.ray.Dir.Y*probe.tmax)
		}
	}
	fmt.Fprintln(f)

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for kk := 0; kk < 4; kk++ {
				d := 4*k + kk
				probe := cc.GetProbe(1, i, j, d)
				fmt.Fprintf(f, "%f %f %f %f\n", probe.ray.P.X, probe.ray.P.Y, probe.ray.Dir.X*probe.tmax, probe.ray.Dir.Y*probe.tmax)
			}
		}
	}
	fmt.Fprintln(f)

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for kk := 0; kk < 4; kk++ {
				d := 4*k + kk
				probe := cc.GetProbe(2, i, j, d)
				fmt.Fprintf(f, "%f %f %f %f\n", probe.ray.P.X, probe.ray.P.Y, probe.ray.Dir.X*probe.tmax, probe.ray.Dir.Y*probe.tmax)
			}

		}
	}

	fmt.Fprintln(f)
}

func PlotCascade3() {
	f, err := os.Create("plots/plot3.data")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	const WIDTH, HEIGHT = 6, 6
	cc := NewCascadeCalculator(WIDTH, HEIGHT)
	fmt.Fprintln(f, "# NCascades", cc.NCascades)
	for c := 0; c < 4; c++ {
		ci := cc.cascadeInfo[c]
		fmt.Fprintln(f, "# Cascade", c, "total", ci.Total(), "dirCount", ci.dirCount, "N", ci.N, "tStart", ci.tStart, "tEnd", ci.tEnd)
		if c == 0 {
			for i := 0; i < ci.N; i++ {
				for j := 0; j < ci.M; j++ {
					for k := 0; k < ci.dirCount; k++ {
						probe := cc.GetProbe(c, i, j, k)
						fmt.Fprintf(f, "%f %f %f %f\n", probe.ray.P.X, probe.ray.P.Y, probe.ray.Dir.X*probe.tmax, probe.ray.Dir.Y*probe.tmax)
					}
				}
			}
		} else {
			for i := 0; i < ci.N; i++ {
				for j := 0; j < ci.M; j++ {
					for k := 0; k < ci.dirCount; k++ {
						probe := cc.GetProbe(c, i, j, k)
						fmt.Fprintf(f, "%f %f %f %f\n", probe.ray.P.X, probe.ray.P.Y, probe.ray.Dir.X*probe.tmax, probe.ray.Dir.Y*probe.tmax)
					}
				}
			}
		}
		fmt.Fprintln(f)
	}
}

func PlotCascade4() {
	f, err := os.Create("plots/plot4.data")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	const WIDTH, HEIGHT = 32, 32
	cc := NewCascadeCalculator(WIDTH, HEIGHT)
	fmt.Fprintln(f, "# NCascades", cc.NCascades)
	for c := 0; c < cc.NCascades; c++ {
		ci := cc.cascadeInfo[c]
		fmt.Fprintln(f, "# Cascade", c, "total", ci.Total(), "dirCount", ci.dirCount, "N", ci.N, "tStart", ci.tStart, "tEnd", ci.tEnd)
		for i := 0; i < ci.N; i++ {
			for j := 0; j < ci.M; j++ {
				for k := 0; k < ci.dirCount; k++ {
					probe := cc.GetProbeCenter(c, i, j)
					fmt.Fprintf(f, "%f %f\n", probe.X, probe.Y)
				}
			}
		}
		fmt.Fprintln(f)
	}
}

func PlotCascade5() {
	f, err := os.Create("plots/plot5.data")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	const WIDTH, HEIGHT = 256, 256
	cc := NewCascadeCalculator(WIDTH, HEIGHT)
	fmt.Fprintln(f, "# NCascades", cc.NCascades)
	for c := 0; c < cc.NCascades; c++ {
		ci := cc.cascadeInfo[c]
		fmt.Fprintln(f, "# Cascade", c, "total", ci.Total(), "dirCount", ci.dirCount, "N", ci.N, "tStart", ci.tStart, "tEnd", ci.tEnd)
		probeCenter := cc.GetProbeCenter(c, 0, 0)
		for k := 0; k < ci.dirCount; k++ {
			probe := cc.GetProbe(c, 0, 0, k)
			probe.ray.P.Sub(probeCenter)
			fmt.Fprintf(f, "%f %f %f %f\n", probe.ray.P.X, probe.ray.P.Y, probe.ray.Dir.X*probe.tmax, probe.ray.Dir.Y*probe.tmax)
		}
		fmt.Fprintln(f)
	}
}

func PlotEnergy(s *primitives.SampledImage) {
	f, err := os.Create("plots/plotEnergy.data")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for x := -350; x < 350; x++ {
		color := s.GetColor(s.Width/2+x, s.Height/2)
		fmt.Fprintln(f, x, color.Intensity())
	}

}
