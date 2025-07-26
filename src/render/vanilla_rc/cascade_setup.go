package vanilla_rc

import (
	"CoreCascade/primitives"
	"fmt"
	"math"
)

type CascadeInfo struct {
	Cascade    int  // the cascade number, starting from 0.
	isLast     bool // true if this is the last cascade
	DirCount   int
	angleStart float64         // Start of the phase function for this cascade.
	deltaAngle float64         // Angular resolution for this cascade.
	TStart     float64         // Start of the Ray.
	TEnd       float64         // End of the Ray.
	spacing    float64         // Spacing of the probes in this cascade.
	p0         primitives.Vec2 // Position of the first probe in this cascade.
	N, M       int             // N and M are used for the grid size in the cascade.
}

func (ci *CascadeInfo) Total() int {
	return ci.DirCount * ci.N * ci.M
}

type CascadeCalculator struct {
	width, height int
	NCascades     int

	dirCountC0                     int     // Number of directions in cascade 0 (angular resolution).
	dirCountMultiplier             int     // Each next cascade has 4 times more directions.
	cellSizeC0                     float64 // Size of a cell in cascade 0, used for spacing.
	lengthC0                       float64 // Length of the first Ray in cascade 0.
	RAY_INTERVAL_LENGTH_MULTIPLIER float64 // Multiplier for Ray interval length in cascades.
	PROBE_SPACING_MULTIPLIER       float64

	CascadeInfo []CascadeInfo
}

func NewCascadeCalculator(width, height int) *CascadeCalculator {
	cc := &CascadeCalculator{
		width:  width,
		height: height,
	}

	// the scene is just fixed to -1,-1 to 1,1
	sceneWidth := 2.
	sceneHeight := 2.
	sceneDiagonal := math.Sqrt(sceneWidth*sceneWidth + sceneHeight*sceneHeight)

	cc.dirCountC0 = 4 // Initial number of directions in the first cascade.
	cc.dirCountMultiplier = 4
	cc.RAY_INTERVAL_LENGTH_MULTIPLIER = 4.

	cc.PROBE_SPACING_MULTIPLIER = 2.

	cc.cellSizeC0 = sceneWidth / float64(width)
	cc.lengthC0 = cc.cellSizeC0 * 0.5
	//lengthC0 := cc.cellSizeC0 * 1.0

	// determine the number of cascades based on the scene width and the length of the first cascade.
	iFloat := math.Log(sceneDiagonal/cc.lengthC0) / math.Log(cc.RAY_INTERVAL_LENGTH_MULTIPLIER)
	cc.NCascades = int(math.Ceil(iFloat)) + 1 // add one, because we iterate to NCascades - 1 in the loop.

	cc.CascadeInfo = make([]CascadeInfo, cc.NCascades+1) // add one additional cascade into array to prevent out of bounds.
	cc.CascadeInfo[0].Cascade = 0
	cc.CascadeInfo[0].isLast = false
	cc.CascadeInfo[0].DirCount = cc.dirCountC0
	cc.CascadeInfo[0].spacing = cc.cellSizeC0
	cc.CascadeInfo[0].N = width
	cc.CascadeInfo[0].M = height
	for i := 1; i < cc.NCascades+1; i++ {
		cc.CascadeInfo[i].Cascade = i
		cc.CascadeInfo[i].isLast = false
		cc.CascadeInfo[i].DirCount = cc.CascadeInfo[i-1].DirCount * cc.dirCountMultiplier
		cc.CascadeInfo[i].spacing = cc.CascadeInfo[i-1].spacing * cc.PROBE_SPACING_MULTIPLIER
		cc.CascadeInfo[i].N = int(math.Ceil(float64(cc.CascadeInfo[i-1].N)/cc.PROBE_SPACING_MULTIPLIER)) + 1
		cc.CascadeInfo[i].M = int(math.Ceil(float64(cc.CascadeInfo[i-1].M)/cc.PROBE_SPACING_MULTIPLIER)) + 1
	}
	cc.CascadeInfo[cc.NCascades].isLast = true

	for i := 0; i < cc.NCascades+1; i++ {
		cc.CascadeInfo[i].angleStart = 0.5 / float64(cc.CascadeInfo[i].DirCount) * math.Pi * 2.
		cc.CascadeInfo[i].deltaAngle = 1. / float64(cc.CascadeInfo[i].DirCount) * math.Pi * 2.

		if i == 0 {
			cc.CascadeInfo[i].TStart = 0.
			cc.CascadeInfo[i].p0 = primitives.Vec2{
				X: 0.5*cc.CascadeInfo[i].spacing - 1.0,
				Y: 0.5*cc.CascadeInfo[i].spacing - 1.0,
			}
		} else {
			cc.CascadeInfo[i].TStart = cc.lengthC0 * math.Pow(cc.RAY_INTERVAL_LENGTH_MULTIPLIER, float64(i)-1.0)
			cc.CascadeInfo[i].p0 = primitives.Vec2{
				X: cc.CascadeInfo[i-1].p0.X - 0.25*cc.CascadeInfo[i].spacing,
				Y: cc.CascadeInfo[i-1].p0.Y - 0.25*cc.CascadeInfo[i].spacing,
			}
		}
		cc.CascadeInfo[i].TEnd = cc.lengthC0 * math.Pow(cc.RAY_INTERVAL_LENGTH_MULTIPLIER, float64(i))
	}

	return cc
}

func (cc *CascadeCalculator) Info() {
	n := 0
	for c := range cc.NCascades {
		ci := cc.CascadeInfo[c]
		fmt.Printf("Cascade %d: dirCount=%6d N=%4d M=%4d total=%8d tStart=%.5f tEnd=%.5f\n", c, ci.DirCount, ci.N, ci.M, ci.Total(), ci.TStart, ci.TEnd)
		n += ci.Total()
	}
	fmt.Println("Total number of probes:", n)
}

type CascadeProbe struct {
	Ray  primitives.Ray
	Tmax float64 // Maximum distance for the Ray to travel.
}

func (ci *CascadeInfo) GetProbe(i int, j int, index int) CascadeProbe {
	angle := ci.angleStart + ci.deltaAngle*float64(index)
	dir := primitives.NewVec2fromAngle(angle)
	ray := primitives.Ray{P: primitives.Vec2{
		ci.p0.X + float64(i)*ci.spacing + dir.X*ci.TStart,
		ci.p0.Y + float64(j)*ci.spacing + dir.Y*ci.TStart},
		Dir: dir}

	return CascadeProbe{
		Ray:  ray,
		Tmax: ci.TEnd - ci.TStart, // Maximum distance for the Ray to travel.
	}
}

func (cc *CascadeCalculator) GetProbeCenter(cascade int, i int, j int) primitives.Vec2 {
	ci := cc.CascadeInfo[cascade]
	p := primitives.Vec2{
		ci.p0.X + float64(i)*ci.spacing,
		ci.p0.Y + float64(j)*ci.spacing,
	}
	return p
}
