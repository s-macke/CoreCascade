package radiance_cascade

import (
	"CoreCascade2D/scene"
	"color"
	"linear_image"
	"vector"
)

type Cascade struct {
	info     CascadeInfo
	radiance [][][]CascadeRadianceResult // x, y, direction (directionLast)
}

func NewCascade(ci CascadeInfo) *Cascade {
	cr := &Cascade{
		info: ci,
	}
	cr.radiance = make([][][]CascadeRadianceResult, ci.N)
	for i := 0; i < ci.N; i++ {
		cr.radiance[i] = make([][]CascadeRadianceResult, ci.M)
		for j := 0; j < ci.M; j++ {
			cr.radiance[i][j] = make([]CascadeRadianceResult, ci.DirCount)
			for k := 0; k < ci.DirCount; k++ {
				cr.radiance[i][j][k] = NewCascadeRadianceResult()
			}
		}
	}
	return cr
}

func BiLinear(ratio vector.Vec2, s0, s1, s2, s3 CascadeRadianceResult) CascadeRadianceResult {
	w1 := (1. - ratio.X) * (1. - ratio.Y)
	w2 := ratio.X * (1. - ratio.Y)
	w3 := (1. - ratio.X) * ratio.Y
	w4 := ratio.X * ratio.Y
	return CascadeRadianceResult{
		color: color.Color{
			R: s0.color.R*w1 + s1.color.R*w2 + s2.color.R*w3 + s3.color.R*w4,
			G: s0.color.G*w1 + s1.color.G*w2 + s2.color.G*w3 + s3.color.G*w4,
			B: s0.color.B*w1 + s1.color.B*w2 + s2.color.B*w3 + s3.color.B*w4,
		},
		visibility: s0.visibility*w1 + s1.visibility*w2 + s2.visibility*w3 + s3.visibility*w4,
	}
}

type RadianceCascade struct {
	width, height int
	bilinearFix   bool
	scene         scene.Scene
	s             *linear_image.SampledImage
	cc            *CascadeCalculator
	cascades      []*Cascade
}

func NewRadianceCascade(scene scene.Scene, s *linear_image.SampledImage, bilinearFix bool) *RadianceCascade {
	rc := &RadianceCascade{
		width:       s.Width,
		height:      s.Height,
		bilinearFix: bilinearFix,
		s:           s,
		cc:          NewCascadeCalculator(s.Width, s.Height),
		scene:       scene,
	}
	rc.cc.Info()
	rc.cascades = make([]*Cascade, rc.cc.NCascades+1)

	for c := 0; c < rc.cc.NCascades+1; c++ {
		ci := rc.cc.CascadeInfo[c]
		rc.cascades[c] = NewCascade(ci)
	}

	return rc
}

func (rc *RadianceCascade) MergeOnImage() {
	// final merge of c0 to determine the pixel color
	// in some way it is similar to merge into cascade "-1"
	cascade0 := rc.cascades[0]
	for y := 0; y < rc.height; y++ {
		for x := 0; x < rc.width; x++ {
			merged := color.Black
			for k := 0; k < cascade0.info.DirCount; k++ {
				merged.Add(cascade0.radiance[x][y][k].color)
			}
			merged.Div(float64(cascade0.info.DirCount)) // average the colors
			rc.s.SetColor(x, y, merged)
		}
	}
}

func (rc *RadianceCascade) Radiance(probe CascadeProbe) CascadeRadianceResult {
	visibility, color := rc.scene.Trace(probe.Ray, probe.Tmax)

	return CascadeRadianceResult{
		color:      color,
		visibility: visibility,
	}
}

func (rc *RadianceCascade) CascadeMerge(cNear *Cascade, cFar *Cascade, x int, y int, k int) {
	ciNear := cNear.info
	ciFar := cFar.info

	factor := float64(ciNear.DirCount) / float64(ciFar.DirCount) // integration factor
	nDirMerge := ciFar.DirCount / ciNear.DirCount                // number of directions to merge
	bilinearWeights := vector.Vec2{X: 0.25 + 0.5*float64(x&1), Y: 0.25 + 0.5*float64(y&1)}
	cNear.radiance[x][y][k] = rc.Radiance(ciNear.GetProbe(x, y, k))
	if !ciFar.isLast { // is not last cascade then merging needed
		merged := NewCascadeRadianceResult()
		for kk := 0; kk < nDirMerge; kk++ { // merge the four directions
			d := 4*k + kk // direction index in the next cascade
			si0 := cFar.radiance[(x>>1)+0][(y>>1)+0][d]
			si1 := cFar.radiance[(x>>1)+1][(y>>1)+0][d]
			si2 := cFar.radiance[(x>>1)+0][(y>>1)+1][d]
			si3 := cFar.radiance[(x>>1)+1][(y>>1)+1][d]
			sBiLinear := BiLinear(bilinearWeights, si0, si1, si2, si3)
			merged.Add(&sBiLinear)
		}
		merged.Mul(factor)
		cNear.radiance[x][y][k].mergeIntervals(&merged) // merge the radiance from the near cascade
	}
}

func (rc *RadianceCascade) Render() {
	rc.s.Clear()

	for c := rc.cc.NCascades - 1; c >= 0; c-- {
		cNear := rc.cascades[c]
		cFar := rc.cascades[c+1]
		for x := 0; x < cNear.info.N; x++ {
			for y := 0; y < cNear.info.M; y++ {
				for k := 0; k < cNear.info.DirCount; k++ {
					if rc.bilinearFix {
						rc.CascadeMergeBilinearFix(cNear, cFar, x, y, k)
					} else {
						rc.CascadeMerge(cNear, cFar, x, y, k)
					}
				}
			}
		}
	}
	rc.MergeOnImage()
}
