package main

import (
	"fmt"
	"math"
	"os"
)

type CascadeInfo struct {
	dirCount   int
	angleStart float64 // Start of the phase function for this cascade.
	deltaAngle float64 // Angular resolution for this cascade.
	tStart     float64 // Start of the ray.
	tEnd       float64 // End of the ray.
	spacing    float64 // Spacing of the probes in this cascade.
	p0         Vec2    // Position of the first probe in this cascade.
	N, M       int     // N and M are used for the grid size in the cascade.
}

func (ci *CascadeInfo) Total() int {
	return ci.dirCount * ci.N * ci.M
}

type CascadeCalculator struct {
	width, height int
	NCascades     int

	dirCountC0                     int     // Number of directions in cascade 0 (angular resolution).
	dirCountMultiplier             int     // Each next cascade has 4 times more directions.
	cellSizeC0                     float64 // Size of a cell in cascade 0, used for spacing.
	RAY_INTERVAL_LENGTH_MULTIPLIER float64 // Multiplier for ray interval length in cascades.
	PROBE_SPACING_MULTIPLIER       float64

	cascadeInfo []CascadeInfo
}

type CascadeProbe struct {
	ray  Ray
	tmax float64 // Maximum distance for the ray to travel.
}

func NewCascadeCalculator(width, height int) *CascadeCalculator {
	cc := &CascadeCalculator{
		width:  width,
		height: height,
	}
	cc.dirCountC0 = 4
	cc.dirCountMultiplier = 4
	cc.RAY_INTERVAL_LENGTH_MULTIPLIER = 4.0

	cc.PROBE_SPACING_MULTIPLIER = 2.0
	cc.cellSizeC0 = 2. / float64(width) // from -1 to 1 in normalized device coordinates, so cell size is 2/width.
	//lengthC0 := cc.cellSizeC0 * 0.5
	lengthC0 := cc.cellSizeC0 * 1.0

	cc.NCascades = 6

	cc.cascadeInfo = make([]CascadeInfo, cc.NCascades)
	cc.cascadeInfo[0].dirCount = cc.dirCountC0
	cc.cascadeInfo[0].spacing = cc.cellSizeC0
	cc.cascadeInfo[0].N = width
	cc.cascadeInfo[0].M = height
	for i := 1; i < cc.NCascades; i++ {
		cc.cascadeInfo[i].dirCount = cc.cascadeInfo[i-1].dirCount * cc.dirCountMultiplier
		cc.cascadeInfo[i].spacing = cc.cascadeInfo[i-1].spacing * cc.PROBE_SPACING_MULTIPLIER
		cc.cascadeInfo[i].N = int(math.Ceil(float64(cc.cascadeInfo[i-1].N)/cc.PROBE_SPACING_MULTIPLIER)) + 1
		cc.cascadeInfo[i].M = int(math.Ceil(float64(cc.cascadeInfo[i-1].M)/cc.PROBE_SPACING_MULTIPLIER)) + 1
	}

	for i := 0; i < cc.NCascades; i++ {
		cc.cascadeInfo[i].angleStart = 0.5 / float64(cc.cascadeInfo[i].dirCount) * math.Pi * 2.
		cc.cascadeInfo[i].deltaAngle = 1. / float64(cc.cascadeInfo[i].dirCount) * math.Pi * 2.

		if i == 0 {
			cc.cascadeInfo[i].tStart = 0.
			cc.cascadeInfo[i].p0 = Vec2{
				X: 0.5*cc.cascadeInfo[i].spacing - 1.0,
				Y: 0.5*cc.cascadeInfo[i].spacing - 1.0,
			}
		} else {
			cc.cascadeInfo[i].tStart = lengthC0 * math.Pow(cc.RAY_INTERVAL_LENGTH_MULTIPLIER, float64(i)-1.0)
			cc.cascadeInfo[i].p0 = Vec2{
				X: cc.cascadeInfo[i-1].p0.X - 0.25*cc.cascadeInfo[i].spacing,
				Y: cc.cascadeInfo[i-1].p0.Y - 0.25*cc.cascadeInfo[i].spacing,
			}
		}
		cc.cascadeInfo[i].tEnd = lengthC0 * math.Pow(cc.RAY_INTERVAL_LENGTH_MULTIPLIER, float64(i))
		/*
			cc.cascadeInfo[i].tStart = 0.
			cc.cascadeInfo[i].tEnd = 0.1
		*/
	}
	return cc
}

func (cc *CascadeCalculator) GetProbe(cascade int, i int, j int, index int) CascadeProbe {
	// Calculate the probe position and spacing based on the cascade and index.
	// This is a placeholder implementation; actual logic will depend on your scene setup.
	ci := cc.cascadeInfo[cascade]
	angle := ci.angleStart + cc.cascadeInfo[cascade].deltaAngle*float64(index)
	dir := Vec2{
		X: math.Cos(angle),
		Y: math.Sin(angle),
	}
	ray := Ray{Vec2{
		ci.p0.X + float64(i)*ci.spacing + dir.X*ci.tStart,
		ci.p0.Y + float64(j)*ci.spacing + dir.Y*ci.tStart},
		Vec2{
			X: dir.X,
			Y: dir.Y,
		}}

	return CascadeProbe{
		ray:  ray,
		tmax: ci.tEnd - ci.tStart, // Maximum distance for the ray to travel.
	}
}

func PlotCascade() {
	f, err := os.Create("plot.data")
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
	f, err := os.Create("plot2.data")
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
	f, err := os.Create("plot3.data")
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
