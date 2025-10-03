package main

import (
	"CoreCascade3D/render/radiance_cascade"
	"fmt"
	"os"
)

func PlotAngleDistribution() {
	f, err := os.Create("plots/probes3D.data")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	const WIDTH, HEIGHT, DEPTH = 8, 8, 8
	cc := radiance_cascade.NewCascadeCalculator(WIDTH, HEIGHT, DEPTH)
	fmt.Fprintln(f, "# NCascades", cc.NCascades)

	for c := 0; c < cc.NCascades; c++ {
		ci := cc.CascadeInfo[c]
		fmt.Fprintln(f, "# Cascade", c, "total", ci.Total(), "dirCount", ci.DirCount, "N", ci.N, "tStart", ci.TStart, "tEnd", ci.TEnd)

		for j := 0; j < ci.DirCount; j++ {
			for i := 0; i < ci.DirCount; i++ {
				fmt.Fprintf(f, "%f %f\n", ci.AngleStart+float32(i)*ci.DeltaAngle, ci.AngleStart+float32(j)*ci.DeltaAngle)
			}
		}
		fmt.Fprintln(f)
	}
}
