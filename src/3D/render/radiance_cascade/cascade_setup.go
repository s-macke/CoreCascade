package radiance_cascade

import (
	"fmt"
	math "github.com/chewxy/math32"
	"vector"
)

type CascadeInfo struct {
	Cascade    int  // the cascade number, starting from 0.
	isLast     bool // true if this is the last cascade
	DirCount   int
	AngleStart float32     // Start of the phase function for this cascade.
	DeltaAngle float32     // Angular resolution for this cascade.
	TStart     float32     // Start of the Ray.
	TEnd       float32     // End of the Ray.
	spacing    float32     // Spacing of the probes in this cascade.
	p0         vector.Vec3 // Position of the first probe in this cascade.
	N, M, O    int         // N and M are used for the grid size in the cascade.
}

func (ci *CascadeInfo) Total() int {
	return ci.DirCount * ci.DirCount * ci.N * ci.M * ci.O
}

type CascadeCalculator struct {
	width, height, depth int
	NCascades            int

	dirCountC0                     int     // Number of directions in cascade 0 (angular resolution).
	dirCountMultiplier             int     // Each next cascade has 4 times more directions.
	cellSizeC0                     float32 // Size of a cell in cascade 0, used for spacing.
	lengthC0                       float32 // Length of the first Ray in cascade 0.
	RAY_INTERVAL_LENGTH_MULTIPLIER float32 // Multiplier for Ray interval length in cascades.
	PROBE_SPACING_MULTIPLIER       float32

	CascadeInfo []CascadeInfo
}

func NewCascadeCalculator(width, height, depth int) *CascadeCalculator {
	cc := &CascadeCalculator{
		width:  width,
		height: height,
		depth:  depth,
	}
	// the scene is just fixed to -1,-1 to 1,1
	sceneWidth := float32(2.)
	sceneHeight := float32(2.)
	sceneDepth := float32(0.1)
	sceneDiagonal := math.Sqrt(sceneWidth*sceneWidth + sceneHeight*sceneHeight + sceneDepth*sceneDepth)
	cc.dirCountC0 = 4         // Initial number of directions in the first cascade. (real number is squared)
	cc.dirCountMultiplier = 2 // (real number is squared)
	cc.RAY_INTERVAL_LENGTH_MULTIPLIER = 4.
	cc.PROBE_SPACING_MULTIPLIER = 2.

	cc.cellSizeC0 = sceneWidth / float32(width) // the same in all directions
	cc.lengthC0 = cc.cellSizeC0 * 1.0

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
	cc.CascadeInfo[0].O = depth

	for i := 1; i < cc.NCascades+1; i++ {
		cc.CascadeInfo[i].Cascade = i
		cc.CascadeInfo[i].isLast = false
		cc.CascadeInfo[i].DirCount = cc.CascadeInfo[i-1].DirCount * cc.dirCountMultiplier
		cc.CascadeInfo[i].spacing = cc.CascadeInfo[i-1].spacing * cc.PROBE_SPACING_MULTIPLIER
		cc.CascadeInfo[i].N = int(math.Ceil(float32(cc.CascadeInfo[i-1].N)/cc.PROBE_SPACING_MULTIPLIER)) + 1
		cc.CascadeInfo[i].M = int(math.Ceil(float32(cc.CascadeInfo[i-1].M)/cc.PROBE_SPACING_MULTIPLIER)) + 1
		cc.CascadeInfo[i].O = int(math.Ceil(float32(cc.CascadeInfo[i-1].O)/cc.PROBE_SPACING_MULTIPLIER)) + 1
	}
	cc.CascadeInfo[cc.NCascades].isLast = true

	for i := 0; i < cc.NCascades+1; i++ {
		cc.CascadeInfo[i].AngleStart = 0.5 / float32(cc.CascadeInfo[i].DirCount)
		cc.CascadeInfo[i].DeltaAngle = 1. / float32(cc.CascadeInfo[i].DirCount)

		if i == 0 {
			cc.CascadeInfo[i].TStart = 0.
			cc.CascadeInfo[i].p0 = vector.Vec3{
				X: 0.5*cc.CascadeInfo[i].spacing - 1.0,
				Y: 0.5*cc.CascadeInfo[i].spacing - 1.0,
				Z: 0.5*cc.CascadeInfo[i].spacing - 1.0,
			}
		} else {
			cc.CascadeInfo[i].TStart = cc.lengthC0 * math.Pow(cc.RAY_INTERVAL_LENGTH_MULTIPLIER, float32(i)-1.0)
			cc.CascadeInfo[i].p0 = vector.Vec3{
				X: cc.CascadeInfo[i-1].p0.X - 0.25*cc.CascadeInfo[i].spacing,
				Y: cc.CascadeInfo[i-1].p0.Y - 0.25*cc.CascadeInfo[i].spacing,
				Z: cc.CascadeInfo[i-1].p0.Z - 0.25*cc.CascadeInfo[i].spacing,
			}
		}
		cc.CascadeInfo[i].TEnd = cc.lengthC0 * math.Pow(cc.RAY_INTERVAL_LENGTH_MULTIPLIER, float32(i))
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
	Ray  vector.Ray3D
	Tmax float32 // Maximum distance for the Ray to travel.
}

func (ci *CascadeInfo) GetProbe(i int, j int, k int, angleI int, angleJ int) CascadeProbe {
	u := ci.AngleStart + ci.DeltaAngle*float32(angleI)
	v := ci.AngleStart + ci.DeltaAngle*float32(angleJ)
	dir := vector.ClarbergToSphere(u, v)

	tstart := ci.TStart
	tend := ci.TEnd

	ray := vector.Ray3D{
		P: vector.Vec3{
			X: ci.p0.X + float32(i)*ci.spacing + dir.X*tstart,
			Y: ci.p0.Y + float32(j)*ci.spacing + dir.Y*tstart,
			Z: ci.p0.Z + float32(k)*ci.spacing + dir.Z*tstart},
		Dir: dir,
	}

	return CascadeProbe{
		Ray:  ray,
		Tmax: tend - tstart, // Maximum distance for the Ray to travel.
	}
}
