package main

import (
	"fmt"
	"math"
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
	sceneWidth := 2.
	sceneHeight := 2.
	sceneDiagonal := math.Sqrt(sceneWidth*sceneWidth + sceneHeight*sceneHeight)
	//sceneDiagonal := max(sceneWidth, sceneHeight) // Use the maximum of width and height for the diagonal.

	cc.dirCountC0 = 4 // Initial number of directions in the first cascade.
	cc.dirCountMultiplier = 4
	cc.RAY_INTERVAL_LENGTH_MULTIPLIER = 4.0

	cc.PROBE_SPACING_MULTIPLIER = 2.0

	cc.cellSizeC0 = sceneWidth / float64(width) // from -1 to 1 in normalized device coordinates, so cell size is 2/width.
	//lengthC0 := cc.cellSizeC0 * 0.5
	lengthC0 := cc.cellSizeC0 * 1.0

	// determine the number of cascades based on the scene width and the length of the first cascade.
	iFloat := math.Log(sceneDiagonal/lengthC0) / math.Log(cc.RAY_INTERVAL_LENGTH_MULTIPLIER)
	fmt.Println("iFloat:", iFloat)
	fmt.Println("diagonal:", sceneDiagonal)
	cc.NCascades = int(math.Ceil(iFloat)) + 1 // add one, because we iterate to NCascades - 1 in the loop.
	fmt.Println("maximum length:", lengthC0*math.Pow(cc.RAY_INTERVAL_LENGTH_MULTIPLIER, float64(cc.NCascades-1)))

	cc.cascadeInfo = make([]CascadeInfo, cc.NCascades+1) // add one additional cascade into array to prevent out of bounds.
	cc.cascadeInfo[0].dirCount = cc.dirCountC0
	cc.cascadeInfo[0].spacing = cc.cellSizeC0
	cc.cascadeInfo[0].N = width
	cc.cascadeInfo[0].M = height
	for i := 1; i < cc.NCascades+1; i++ {
		cc.cascadeInfo[i].dirCount = cc.cascadeInfo[i-1].dirCount * cc.dirCountMultiplier
		cc.cascadeInfo[i].spacing = cc.cascadeInfo[i-1].spacing * cc.PROBE_SPACING_MULTIPLIER
		cc.cascadeInfo[i].N = int(math.Ceil(float64(cc.cascadeInfo[i-1].N)/cc.PROBE_SPACING_MULTIPLIER)) + 1
		cc.cascadeInfo[i].M = int(math.Ceil(float64(cc.cascadeInfo[i-1].M)/cc.PROBE_SPACING_MULTIPLIER)) + 1
	}

	for i := 0; i < cc.NCascades+1; i++ {
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

		//cc.cascadeInfo[i].tStart = 0.
		//cc.cascadeInfo[i].tEnd = 0.1
	}

	fmt.Println("Number of cascades:", cc.NCascades)
	N := 0
	for _, ci := range cc.cascadeInfo {
		N += ci.Total()
	}
	fmt.Println("Total number of probes:", N)

	return cc
}

func (cc *CascadeCalculator) GetProbe(cascade int, i int, j int, index int) CascadeProbe {
	ci := cc.cascadeInfo[cascade]
	angle := ci.angleStart + cc.cascadeInfo[cascade].deltaAngle*float64(index)
	dir := NewVec2fromAngle(angle)
	ray := Ray{p: Vec2{
		ci.p0.X + float64(i)*ci.spacing + dir.X*ci.tStart,
		ci.p0.Y + float64(j)*ci.spacing + dir.Y*ci.tStart},
		dir: dir}

	return CascadeProbe{
		ray:  ray,
		tmax: ci.tEnd - ci.tStart, // Maximum distance for the ray to travel.
	}
}

func (cc *CascadeCalculator) GetProbeCenter(cascade int, i int, j int) Vec2 {
	ci := cc.cascadeInfo[cascade]
	p := Vec2{
		ci.p0.X + float64(i)*ci.spacing,
		ci.p0.Y + float64(j)*ci.spacing,
	}
	return p
}
