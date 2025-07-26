package vanilla_rc

import (
	"CoreCascade/primitives"
	"CoreCascade/scene"
)

type Cascade struct {
	info    CascadeInfo
	cascade [][][]CascadeProbeResult // x, y, direction (directionLast)
}

func NewCascade(ci CascadeInfo) *Cascade {
	cr := &Cascade{
		info: ci,
	}
	cr.cascade = make([][][]CascadeProbeResult, ci.N)
	for i := 0; i < ci.N; i++ {
		cr.cascade[i] = make([][]CascadeProbeResult, ci.M)
		for j := 0; j < ci.M; j++ {
			cr.cascade[i][j] = make([]CascadeProbeResult, ci.DirCount)
			for k := 0; k < ci.DirCount; k++ {
				cr.cascade[i][j][k] = NewCascadeProbeResult()
			}
		}
	}
	return cr
}

func BiLinear(ratio primitives.Vec2, s0, s1, s2, s3 CascadeProbeResult) CascadeProbeResult {
	w1 := (1. - ratio.X) * (1. - ratio.Y)
	w2 := ratio.X * (1. - ratio.Y)
	w3 := (1. - ratio.X) * ratio.Y
	w4 := ratio.X * ratio.Y
	return CascadeProbeResult{
		color: primitives.Color{
			R: s0.color.R*w1 + s1.color.R*w2 + s2.color.R*w3 + s3.color.R*w4,
			G: s0.color.G*w1 + s1.color.G*w2 + s2.color.G*w3 + s3.color.G*w4,
			B: s0.color.B*w1 + s1.color.B*w2 + s2.color.B*w3 + s3.color.B*w4,
		},
		visibility: s0.visibility*w1 + s1.visibility*w2 + s2.visibility*w3 + s3.visibility*w4,
	}
}

type RadianceCascadeVanilla struct {
	width, height int
	scene         *scene.Scene
	s             *primitives.SampledImage
	cc            *CascadeCalculator
	cascadeResult []*Cascade
}

func NewRadianceCascadeVanilla(scene *scene.Scene, s *primitives.SampledImage) *RadianceCascadeVanilla {
	rc := &RadianceCascadeVanilla{
		width:  s.Width,
		height: s.Height,
		s:      s,
		cc:     NewCascadeCalculator(s.Width, s.Height),
		scene:  scene,
	}
	rc.cc.Info()
	rc.cascadeResult = make([]*Cascade, rc.cc.NCascades+1)

	for c := 0; c < rc.cc.NCascades+1; c++ {
		ci := rc.cc.CascadeInfo[c]
		rc.cascadeResult[c] = NewCascade(ci)
	}

	return rc
}

func (rc *RadianceCascadeVanilla) MergeOnImage() {
	// final merge of c0 to determine the pixel color
	cascade0 := rc.cascadeResult[0]
	for y := 0; y < rc.height; y++ {
		for x := 0; x < rc.width; x++ {
			col := primitives.Black
			for k := 0; k < cascade0.info.DirCount; k++ {
				col.Add(cascade0.cascade[x][y][k].color)
			}
			col.Div(float64(cascade0.info.DirCount)) // average the colors
			rc.s.SetColor(x, y, col)
		}
	}
}

func (rc *RadianceCascadeVanilla) Radiance(ci *CascadeInfo, i int, j int, index int) CascadeProbeResult {
	probe := ci.GetProbe(i, j, index)
	hit, color := rc.scene.Intersect(probe.Ray, probe.Tmax)

	// 1. it hit nothing, 0. if hit
	visibility := 1.
	if hit {
		visibility = 0.
	}

	return CascadeProbeResult{
		color:      color,
		visibility: visibility,
	}
}

func (rc *RadianceCascadeVanilla) CascadeMerge(cNear *Cascade, cFar *Cascade) {
	ciNear := cNear.info
	ciFar := cFar.info

	factor := float64(ciNear.DirCount) / float64(ciFar.DirCount) // integration factor
	nDirMerge := ciFar.DirCount / ciNear.DirCount                // number of directions to merge
	for x := 0; x < ciNear.N; x++ {
		for y := 0; y < ciNear.M; y++ {
			bilinearWeights := primitives.Vec2{X: 0.25 + 0.5*float64(x&1), Y: 0.25 + 0.5*float64(y&1)}
			for k := 0; k < ciNear.DirCount; k++ {
				cNear.cascade[x][y][k] = rc.Radiance(&ciNear, x, y, k)
				if !ciFar.isLast { // is not last cascade then merging needed
					merged := NewCascadeProbeResult()
					for kk := 0; kk < nDirMerge; kk++ { // merge the four directions
						d := 4*k + kk // direction index in the next cascade
						si0 := cFar.cascade[(x>>1)+0][(y>>1)+0][d]
						si1 := cFar.cascade[(x>>1)+1][(y>>1)+0][d]
						si2 := cFar.cascade[(x>>1)+0][(y>>1)+1][d]
						si3 := cFar.cascade[(x>>1)+1][(y>>1)+1][d]
						sBiLinear := BiLinear(bilinearWeights, si0, si1, si2, si3)
						merged.Add(&sBiLinear)
					}
					merged.Mul(factor)
					cNear.cascade[x][y][k].mergeIntervals(&merged) // merge the radiance from the near cascade
				}
			}
		}
	}
}

func (rc *RadianceCascadeVanilla) Render() {
	rc.s.Clear()

	for c := rc.cc.NCascades - 1; c >= 0; c-- {
		cNear := rc.cascadeResult[c]
		cFar := rc.cascadeResult[c+1]
		rc.CascadeMerge(cNear, cFar)
	}
	/*
		{
			// TODO, debug
			c0 := cc.CascadeInfo[1]
			c0R := cascadeResult[1]
			for x := 0; x < 100; x++ {
				for k := 0; k < c0.DirCount; k++ {
					probe0 := cc.GetProbe(1, x, 0, k)
					probe1 := cc.GetProbe(1, x, c0.M, k)
					col0 := c0R.cascade[x][0][k].color
					col1 := c0R.cascade[x][c0.M][k].color
					if col0.R != col1.R || col0.G != col1.G || col0.B != col1.B {
						fmt.Println(x, k, probe0.Ray.p.Y, probe1.Ray.p.Y, col0, col1)
					}
				}
			}
		}
	*/
	rc.MergeOnImage()
}
