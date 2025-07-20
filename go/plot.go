package main

import (
	"fmt"
	"os"
)

func PlotSignedDistance() {
	scene := NewScene()
	for x := -20.0; x <= 20.0; x += 0.5 {
		for y := -20.0; y <= 20.0; y += 0.5 {
			p := Vec2{X: x, Y: y}
			d, _ := scene.sd(p)
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

	const WIDTH, HEIGHT = 800, 800
	cc := NewCascadeCalculator(WIDTH, HEIGHT)
	fmt.Fprintln(f, "# NCascades", cc.NCascades)
	for c := 0; c < 3; c++ {
		ci := cc.cascadeInfo[c]
		fmt.Fprintln(f, "# Cascade", c, "total", ci.Total(), "dirCount", ci.dirCount, "N", ci.N, "tStart", ci.tStart, "tEnd", ci.tEnd)
		for i := 0; i < 4>>c; i++ {
			for j := 0; j < 4>>c; j++ {
				for k := 0; k < ci.dirCount; k++ {
					probe := cc.GetProbe(c, i, j, k)
					fmt.Fprintf(f, "%f %f %f %f\n", probe.ray.p.X, probe.ray.p.Y, probe.ray.dir.X*probe.tmax, probe.ray.dir.Y*probe.tmax)
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
			fmt.Fprintf(f, "%f %f %f %f\n", probe.ray.p.X, probe.ray.p.Y, probe.ray.dir.X*probe.tmax, probe.ray.dir.Y*probe.tmax)
		}
	}
	fmt.Fprintln(f)

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for kk := 0; kk < 4; kk++ {
				d := 4*k + kk
				probe := cc.GetProbe(1, i, j, d)
				fmt.Fprintf(f, "%f %f %f %f\n", probe.ray.p.X, probe.ray.p.Y, probe.ray.dir.X*probe.tmax, probe.ray.dir.Y*probe.tmax)
			}
		}
	}
	fmt.Fprintln(f)

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for kk := 0; kk < 4; kk++ {
				d := 4*k + kk
				probe := cc.GetProbe(2, i, j, d)
				fmt.Fprintf(f, "%f %f %f %f\n", probe.ray.p.X, probe.ray.p.Y, probe.ray.dir.X*probe.tmax, probe.ray.dir.Y*probe.tmax)
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
						fmt.Fprintf(f, "%f %f %f %f\n", probe.ray.p.X, probe.ray.p.Y, probe.ray.dir.X*probe.tmax, probe.ray.dir.Y*probe.tmax)
					}
				}
			}
		} else {
			for i := 0; i < ci.N; i++ {
				for j := 0; j < ci.M; j++ {
					for k := 0; k < ci.dirCount; k++ {
						probe := cc.GetProbe(c, i, j, k)
						fmt.Fprintf(f, "%f %f %f %f\n", probe.ray.p.X, probe.ray.p.Y, probe.ray.dir.X*probe.tmax, probe.ray.dir.Y*probe.tmax)
					}
				}
			}
		}
		fmt.Fprintln(f)
	}
}
