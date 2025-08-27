package main

import (
	"CoreCascade/primitives"
	"CoreCascade/render/radiance_cascade"
	"CoreCascade/scene/scenes"
	"fmt"
	"os"
)

func PlotSignedDistance() {
	scene := scenes.NewSceneShadows(0)
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
	cc := radiance_cascade.NewCascadeCalculator(WIDTH, HEIGHT)
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

func PlotCascadeBilinearFix() {
	f, err := os.Create("plots/probes_bilinear_fix.data")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	const WIDTH, HEIGHT = 8, 8
	cc := radiance_cascade.NewCascadeCalculator(WIDTH, HEIGHT)
	fmt.Fprintln(f, "# NCascades", cc.NCascades)
	for c := 0; c < cc.NCascades; c++ {
		ciNear := cc.CascadeInfo[c]
		ciFar := cc.CascadeInfo[c+1]
		fmt.Fprintln(f, "# Cascade", c, "total", ciNear.Total(), "dirCount", ciNear.DirCount, "N", ciNear.N, "tStart", ciNear.TStart, "tEnd", ciNear.TEnd)
		for i := 0; i < ciNear.N; i++ {
			for j := 0; j < ciNear.M; j++ {
				for k := 0; k < ciNear.DirCount; k++ {
					probe1 := radiance_cascade.GetBilinearFixProbe(&ciNear, &ciFar, i, j, (i>>1)+0, (j>>1)+0, k)
					probe2 := radiance_cascade.GetBilinearFixProbe(&ciNear, &ciFar, i, j, (i>>1)+1, (j>>1)+0, k)
					probe3 := radiance_cascade.GetBilinearFixProbe(&ciNear, &ciFar, i, j, (i>>1)+1, (j>>1)+1, k)
					probe4 := radiance_cascade.GetBilinearFixProbe(&ciNear, &ciFar, i, j, (i>>1)+0, (j>>1)+1, k)
					fmt.Fprintf(f, "%f %f %f %f\n", probe1.Ray.P.X, probe1.Ray.P.Y, probe1.Ray.Dir.X*probe1.Tmax, probe1.Ray.Dir.Y*probe1.Tmax)
					fmt.Fprintf(f, "%f %f %f %f\n", probe2.Ray.P.X, probe2.Ray.P.Y, probe2.Ray.Dir.X*probe2.Tmax, probe2.Ray.Dir.Y*probe2.Tmax)
					fmt.Fprintf(f, "%f %f %f %f\n", probe3.Ray.P.X, probe3.Ray.P.Y, probe3.Ray.Dir.X*probe3.Tmax, probe3.Ray.Dir.Y*probe3.Tmax)
					fmt.Fprintf(f, "%f %f %f %f\n", probe4.Ray.P.X, probe4.Ray.P.Y, probe4.Ray.Dir.X*probe4.Tmax, probe4.Ray.Dir.Y*probe4.Tmax)
				}
			}
		}
		fmt.Fprintln(f)
	}
}

func PlotCascadeBilinearFixSimple() {
	f, err := os.Create("plots/probes_bilinear_fix_simple.data")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	const WIDTH, HEIGHT = 8, 8
	cc := radiance_cascade.NewCascadeCalculator(WIDTH, HEIGHT)
	ciNear := cc.CascadeInfo[0]
	ciFar := cc.CascadeInfo[1]

	for k := 0; k < ciNear.DirCount; k++ {
		probe := ciNear.GetProbe(0, 0, k)

		probe1 := radiance_cascade.GetBilinearFixProbe(&ciNear, &ciFar, 0, 0, +0, +0, k)
		probe2 := radiance_cascade.GetBilinearFixProbe(&ciNear, &ciFar, 0, 0, +1, +0, k)
		probe3 := radiance_cascade.GetBilinearFixProbe(&ciNear, &ciFar, 0, 0, +1, +1, k)
		probe4 := radiance_cascade.GetBilinearFixProbe(&ciNear, &ciFar, 0, 0, +0, +1, k)
		fmt.Fprintf(f, "%f %f %f %f %f %f\n", probe1.Ray.P.X, probe1.Ray.P.Y, probe1.Ray.Dir.X*probe1.Tmax, probe1.Ray.Dir.Y*probe1.Tmax, probe.Ray.Dir.X*probe.Tmax, probe.Ray.Dir.Y*probe.Tmax)
		fmt.Fprintf(f, "%f %f %f %f %f %f\n", probe2.Ray.P.X, probe2.Ray.P.Y, probe2.Ray.Dir.X*probe2.Tmax, probe2.Ray.Dir.Y*probe2.Tmax, probe.Ray.Dir.X*probe.Tmax, probe.Ray.Dir.Y*probe.Tmax)
		fmt.Fprintf(f, "%f %f %f %f %f %f\n", probe3.Ray.P.X, probe3.Ray.P.Y, probe3.Ray.Dir.X*probe3.Tmax, probe3.Ray.Dir.Y*probe3.Tmax, probe.Ray.Dir.X*probe.Tmax, probe.Ray.Dir.Y*probe.Tmax)
		fmt.Fprintf(f, "%f %f %f %f %f %f\n", probe4.Ray.P.X, probe4.Ray.P.Y, probe4.Ray.Dir.X*probe4.Tmax, probe4.Ray.Dir.Y*probe4.Tmax, probe.Ray.Dir.X*probe.Tmax, probe.Ray.Dir.Y*probe.Tmax)
	}
	fmt.Fprintln(f)

	ciNear = cc.CascadeInfo[1]
	for k := 0; k < ciNear.DirCount; k++ {
		probe := ciNear.GetProbe(0, 0, k)
		fmt.Fprintf(f, "%f %f %f %f\n", probe.Ray.P.X, probe.Ray.P.Y, probe.Ray.Dir.X*probe.Tmax, probe.Ray.Dir.Y*probe.Tmax)
		probe = ciNear.GetProbe(1, 0, k)
		fmt.Fprintf(f, "%f %f %f %f\n", probe.Ray.P.X, probe.Ray.P.Y, probe.Ray.Dir.X*probe.Tmax, probe.Ray.Dir.Y*probe.Tmax)
		probe = ciNear.GetProbe(0, 1, k)
		fmt.Fprintf(f, "%f %f %f %f\n", probe.Ray.P.X, probe.Ray.P.Y, probe.Ray.Dir.X*probe.Tmax, probe.Ray.Dir.Y*probe.Tmax)
		probe = ciNear.GetProbe(1, 1, k)
		fmt.Fprintf(f, "%f %f %f %f\n", probe.Ray.P.X, probe.Ray.P.Y, probe.Ray.Dir.X*probe.Tmax, probe.Ray.Dir.Y*probe.Tmax)
	}

}

func PlotCascade2() {
	f, err := os.Create("plots/plot2.data")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	const WIDTH, HEIGHT = 800, 800
	cc := radiance_cascade.NewCascadeCalculator(WIDTH, HEIGHT)

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
	cc := radiance_cascade.NewCascadeCalculator(WIDTH, HEIGHT)
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
	cc := radiance_cascade.NewCascadeCalculator(WIDTH, HEIGHT)
	fmt.Fprintln(f, "# NCascades", cc.NCascades)
	for c := 0; c < cc.NCascades; c++ {
		ci := cc.CascadeInfo[c]
		fmt.Fprintln(f, "# Cascade", c, "total", ci.Total(), "dirCount", ci.DirCount, "N", ci.N, "tStart", ci.TStart, "tEnd", ci.TEnd)
		for i := 0; i < ci.N; i++ {
			for j := 0; j < ci.M; j++ {
				for k := 0; k < ci.DirCount; k++ {
					probe := ci.GetProbeCenter(i, j)
					fmt.Fprintf(f, "%f %f\n", probe.X, probe.Y)
				}
			}
		}
		fmt.Fprintln(f)
	}
}

func PlotProbeCascadesNonSpatial() {
	f, err := os.Create("plots/probe_cascades_non_spatial.data")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	const WIDTH, HEIGHT = 256, 256
	cc := radiance_cascade.NewCascadeCalculator(WIDTH, HEIGHT)
	fmt.Fprintln(f, "# NCascades", cc.NCascades)
	for c := 0; c < cc.NCascades; c++ {
		ci := cc.CascadeInfo[c]
		fmt.Fprintln(f, "# Cascade", c, "total", ci.Total(), "dirCount", ci.DirCount, "N", ci.N, "tStart", ci.TStart, "tEnd", ci.TEnd)
		probeCenter := ci.GetProbeCenter(0, 0)
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
