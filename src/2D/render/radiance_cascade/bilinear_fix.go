package radiance_cascade

import (
	"vector"
)

func GetBilinearFixProbe(ciNear *CascadeInfo, ciFar *CascadeInfo, iNear int, jNear int, iFar int, jFar int, index int) CascadeProbe {

	angle := ciNear.angleStart + ciNear.deltaAngle*float64(index)
	dir := vector.NewVec2fromAngle(angle)
	tstart := ciNear.TStart
	tend := ciNear.TEnd

	pNear := vector.Vec2{
		X: ciNear.p0.X + float64(iNear)*ciNear.spacing + dir.X*tstart,
		Y: ciNear.p0.Y + float64(jNear)*ciNear.spacing + dir.Y*tstart,
	}
	pFar := vector.Vec2{
		X: ciFar.p0.X + float64(iFar)*ciFar.spacing + dir.X*tend,
		Y: ciFar.p0.Y + float64(jFar)*ciFar.spacing + dir.Y*tend,
	}
	pFar.Sub(pNear)
	length := pFar.Normalize()

	ray := vector.Ray2D{
		P:   pNear,
		Dir: pFar,
	}

	return CascadeProbe{
		Ray:  ray,
		Tmax: length, // Maximum distance for the Ray to travel.
	}
}

func (rc *RadianceCascade) CascadeMergeBilinearFix(cNear *Cascade, cFar *Cascade, x int, y int, k int) {
	ciNear := cNear.info
	ciFar := cFar.info

	factor := float64(ciNear.DirCount) / float64(ciFar.DirCount) // integration factor
	nDirMerge := ciFar.DirCount / ciNear.DirCount                // number of directions to merge
	bilinearWeights := vector.Vec2{X: 0.25 + 0.5*float64(x&1), Y: 0.25 + 0.5*float64(y&1)}
	merged := NewCascadeRadianceResult()
	if ciNear.isLast {
		cNear.radiance[x][y][k] = rc.Radiance(ciNear.GetProbe(x, y, k))
	} else {
		sNear0 := rc.Radiance(GetBilinearFixProbe(&ciNear, &ciFar, x, y, (x>>1)+0, (y>>1)+0, k))
		sNear1 := rc.Radiance(GetBilinearFixProbe(&ciNear, &ciFar, x, y, (x>>1)+1, (y>>1)+0, k))
		sNear2 := rc.Radiance(GetBilinearFixProbe(&ciNear, &ciFar, x, y, (x>>1)+0, (y>>1)+1, k))
		sNear3 := rc.Radiance(GetBilinearFixProbe(&ciNear, &ciFar, x, y, (x>>1)+1, (y>>1)+1, k))
		for kk := 0; kk < nDirMerge; kk++ { // merge the four directions
			d := 4*k + kk // direction index in the next cascade
			si0 := cFar.radiance[(x>>1)+0][(y>>1)+0][d]
			si1 := cFar.radiance[(x>>1)+1][(y>>1)+0][d]
			si2 := cFar.radiance[(x>>1)+0][(y>>1)+1][d]
			si3 := cFar.radiance[(x>>1)+1][(y>>1)+1][d]
			sNear0Temp := sNear0
			sNear1Temp := sNear1
			sNear2Temp := sNear2
			sNear3Temp := sNear3

			sNear0Temp.mergeIntervals(&si0)
			sNear1Temp.mergeIntervals(&si1)
			sNear2Temp.mergeIntervals(&si2)
			sNear3Temp.mergeIntervals(&si3)
			sBiLinear := BiLinear(bilinearWeights, sNear0Temp, sNear1Temp, sNear2Temp, sNear3Temp)
			merged.Add(&sBiLinear)
		}
		merged.Mul(factor)
		cNear.radiance[x][y][k] = merged
	}
}
