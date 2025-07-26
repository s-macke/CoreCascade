package main

import (
	"CoreCascade/primitives"
	"CoreCascade/render/vanilla_rc"
	"fmt"
	"os"
)

func PlotSignedDistance() {
	scene := NewScene("shadows", 0.)
	for x := -2.0; x <= 2.0; x += 0.1 {
		for y := -2.0; y <= 2.0; y += 0.1 {
			d, _ := scene.SignedDistance(primitives.Vec2{X: x, Y: y})
			fmt.Println(x, y, d)
		}
		fmt.Println()
	}

}

func PlotCascade() {
	f, err := os.Create("plots/probes.data")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	const WIDTH, HEIGHT = 8, 8
	cc := vanilla_rc.NewCascadeCalculator(WIDTH, HEIGHT)
	fmt.Fprintln(f, "# NCascades", cc.NCascades)
	for c := 0; c < cc.NCascades; c++ {
		ci := cc.CascadeInfo[c]
		fmt.Fprintln(f, "# Cascade", c, "total", ci.Total(), "dirCount", ci.DirCount, "N", ci.N, "tStart", ci.TStart, "tEnd", ci.TEnd)
		for i := 0; i < ci.N; i++ {
			for j := 0; j < ci.M; j++ {
				for k := 0; k < ci.DirCount; k++ {
					probe := ci.GetProbe(i, j, k)
					fmt.Fprintf(f, "%f %f %f %f\n", probe.Ray.P.X, probe.Ray.P.Y, probe.Ray.Dir.X*probe.Tmax, probe.Ray.Dir.Y*probe.Tmax)
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
	cc := vanilla_rc.NewCascadeCalculator(WIDTH, HEIGHT)

	k := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			probe := cc.CascadeInfo[0].GetProbe(i, j, k)
			fmt.Fprintf(f, "%f %f %f %f\n", probe.Ray.P.X, probe.Ray.P.Y, probe.Ray.Dir.X*probe.Tmax, probe.Ray.Dir.Y*probe.Tmax)
		}
	}
	fmt.Fprintln(f)

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for kk := 0; kk < 4; kk++ {
				d := 4*k + kk
				probe := cc.CascadeInfo[1].GetProbe(i, j, d)
				fmt.Fprintf(f, "%f %f %f %f\n", probe.Ray.P.X, probe.Ray.P.Y, probe.Ray.Dir.X*probe.Tmax, probe.Ray.Dir.Y*probe.Tmax)
			}
		}
	}
	fmt.Fprintln(f)

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for kk := 0; kk < 4; kk++ {
				d := 4*k + kk
				probe := cc.CascadeInfo[2].GetProbe(i, j, d)
				fmt.Fprintf(f, "%f %f %f %f\n", probe.Ray.P.X, probe.Ray.P.Y, probe.Ray.Dir.X*probe.Tmax, probe.Ray.Dir.Y*probe.Tmax)
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
	cc := vanilla_rc.NewCascadeCalculator(WIDTH, HEIGHT)
	fmt.Fprintln(f, "# NCascades", cc.NCascades)
	for c := 0; c < 4; c++ {
		ci := cc.CascadeInfo[c]
		fmt.Fprintln(f, "# Cascade", c, "total", ci.Total(), "dirCount", ci.DirCount, "N", ci.N, "tStart", ci.TStart, "tEnd", ci.TEnd)
		if c == 0 {
			for i := 0; i < ci.N; i++ {
				for j := 0; j < ci.M; j++ {
					for k := 0; k < ci.DirCount; k++ {
						probe := ci.GetProbe(i, j, k)
						fmt.Fprintf(f, "%f %f %f %f\n", probe.Ray.P.X, probe.Ray.P.Y, probe.Ray.Dir.X*probe.Tmax, probe.Ray.Dir.Y*probe.Tmax)
					}
				}
			}
		} else {
			for i := 0; i < ci.N; i++ {
				for j := 0; j < ci.M; j++ {
					for k := 0; k < ci.DirCount; k++ {
						probe := ci.GetProbe(i, j, k)
						fmt.Fprintf(f, "%f %f %f %f\n", probe.Ray.P.X, probe.Ray.P.Y, probe.Ray.Dir.X*probe.Tmax, probe.Ray.Dir.Y*probe.Tmax)
					}
				}
			}
		}
		fmt.Fprintln(f)
	}
}

func PlotProbeCenter() {
	f, err := os.Create("plots/probe_center.data")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	const WIDTH, HEIGHT = 32, 32
	cc := vanilla_rc.NewCascadeCalculator(WIDTH, HEIGHT)
	fmt.Fprintln(f, "# NCascades", cc.NCascades)
	for c := 0; c < cc.NCascades; c++ {
		ci := cc.CascadeInfo[c]
		fmt.Fprintln(f, "# Cascade", c, "total", ci.Total(), "dirCount", ci.DirCount, "N", ci.N, "tStart", ci.TStart, "tEnd", ci.TEnd)
		for i := 0; i < ci.N; i++ {
			for j := 0; j < ci.M; j++ {
				for k := 0; k < ci.DirCount; k++ {
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
	cc := vanilla_rc.NewCascadeCalculator(WIDTH, HEIGHT)
	fmt.Fprintln(f, "# NCascades", cc.NCascades)
	for c := 0; c < cc.NCascades; c++ {
		ci := cc.CascadeInfo[c]
		fmt.Fprintln(f, "# Cascade", c, "total", ci.Total(), "dirCount", ci.DirCount, "N", ci.N, "tStart", ci.TStart, "tEnd", ci.TEnd)
		probeCenter := cc.GetProbeCenter(c, 0, 0)
		for k := 0; k < ci.DirCount; k++ {
			probe := ci.GetProbe(0, 0, k)
			probe.Ray.P.Sub(probeCenter)
			fmt.Fprintf(f, "%f %f %f %f\n", probe.Ray.P.X, probe.Ray.P.Y, probe.Ray.Dir.X*probe.Tmax, probe.Ray.Dir.Y*probe.Tmax)
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
