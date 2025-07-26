package main

import (
	"CoreCascade/primitives"
	"CoreCascade/scene"
)

type CascadeIntervalResult struct {
	color      primitives.Color
	visibility float64 // 1.0 if the interval hit nothing, and 0.0 it it did.
}

func (c *CascadeIntervalResult) mergeIntervals(far *CascadeIntervalResult) {
	c.color.R += far.color.R * c.visibility
	c.color.G += far.color.G * c.visibility
	c.color.B += far.color.B * c.visibility
	c.visibility *= far.visibility
}

func (c *CascadeIntervalResult) Mul(factor float64) {
	c.color.R *= factor
	c.color.G *= factor
	c.color.B *= factor
	c.visibility *= factor
}

func (c *CascadeIntervalResult) Add(b *CascadeIntervalResult) {
	c.color.R += b.color.R
	c.color.G += b.color.G
	c.color.B += b.color.B
	c.visibility += b.visibility
}

type CascadeResult struct {
	cascade [][][]CascadeIntervalResult // x, y, direction (directionLast)
}

func NewCascadeResult(c CascadeInfo) *CascadeResult {
	cr := &CascadeResult{}
	cr.cascade = make([][][]CascadeIntervalResult, c.N+1)
	for i := 0; i < c.N+1; i++ {
		cr.cascade[i] = make([][]CascadeIntervalResult, c.M+1)
		for j := 0; j < c.M+1; j++ {
			cr.cascade[i][j] = make([]CascadeIntervalResult, c.dirCount)
			for k := 0; k < c.dirCount; k++ {
				cr.cascade[i][j][k] = CascadeIntervalResult{
					color:      primitives.Color{R: 0, G: 0, B: 0},
					visibility: 0.0,
				}
			}
		}
	}
	return cr
}

func BiLinear(ratio primitives.Vec2, s0, s1, s2, s3 CascadeIntervalResult) CascadeIntervalResult {
	w1 := (1. - ratio.X) * (1. - ratio.Y)
	w2 := ratio.X * (1. - ratio.Y)
	w3 := (1. - ratio.X) * ratio.Y
	w4 := ratio.X * ratio.Y
	return CascadeIntervalResult{
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
	cascadeResult []*CascadeResult
}

func NewRadianceCascadeVanilla(scene *scene.Scene, s *primitives.SampledImage) *RadianceCascadeVanilla {
	rc := &RadianceCascadeVanilla{
		width:  s.Width,
		height: s.Height,
		s:      s,
		cc:     NewCascadeCalculator(s.Width, s.Height),
		scene:  scene,
	}
	rc.cascadeResult = make([]*CascadeResult, rc.cc.NCascades+1)

	for c := 0; c < rc.cc.NCascades+1; c++ {
		ci := rc.cc.cascadeInfo[c]
		rc.cascadeResult[c] = NewCascadeResult(ci)
	}

	return rc
}

func (rc *RadianceCascadeVanilla) MergeOnImage() {
	// final merge of c0 to determine the pixel color
	c0 := rc.cc.cascadeInfo[0]
	c0R := rc.cascadeResult[0]
	for y := 0; y < rc.height; y++ {
		for x := 0; x < rc.width; x++ {
			col := primitives.Black
			for k := 0; k < c0.dirCount; k++ {
				col.Add(c0R.cascade[x][y][k].color)
			}
			col.Div(float64(c0.dirCount)) // average the colors
			//col.Mul(2. * math.Pi)
			rc.s.SetColor(x, y, col)
		}
	}
}

func (rc *RadianceCascadeVanilla) Radiance(cascade int, i int, j int, index int) CascadeIntervalResult {
	probe := rc.cc.GetProbe(cascade, i, j, index)
	hit, color := rc.scene.Intersect(probe.ray, probe.tmax)

	// 1. it hit nothing, 0. if hit
	visibility := 1.
	if hit {
		visibility = 0.
	}

	return CascadeIntervalResult{
		color:      color,
		visibility: visibility,
	}
}

func (rc *RadianceCascadeVanilla) Cascade(c int) {
	ciNear := rc.cc.cascadeInfo[c]
	ciRNear := rc.cascadeResult[c]

	ciFar := rc.cc.cascadeInfo[c+1]
	ciRFar := rc.cascadeResult[c+1]

	factor := float64(ciNear.dirCount) / float64(ciFar.dirCount) // integration factor
	nDirMerge := ciFar.dirCount / ciNear.dirCount                // number of directions to merge
	for x := 0; x < ciNear.N; x++ {
		for y := 0; y < ciNear.M; y++ {
			for k := 0; k < ciNear.dirCount; k++ {
				s := CascadeIntervalResult{}
				siNearTemp := rc.Radiance(c, x, y, k)
				for kk := 0; kk < nDirMerge; kk++ {
					siNear := siNearTemp // copy radiance
					if c != rc.cc.NCascades-1 {
						d := 4*k + kk // direction index in the next cascade
						si0 := ciRFar.cascade[(x>>1)+0][(y>>1)+0][d]
						si1 := ciRFar.cascade[(x>>1)+1][(y>>1)+0][d]
						si2 := ciRFar.cascade[(x>>1)+0][(y>>1)+1][d]
						si3 := ciRFar.cascade[(x>>1)+1][(y>>1)+1][d]
						/*
							rnear := rc.cc.GetProbeCenter(c, x, y)
							rfar0 := rc.cc.GetProbeCenter(c+1, x>>1, y>>1)
							rfar1 := rc.cc.GetProbeCenter(c+1, (x>>1)+1, y>>1)
							fmt.Println("total ", rfar1.X-rfar0.X, rnear.X-rfar0.X, rnear.X-rfar1.X)
							if rfar0.X > rnear.X {
								panic("Bad cascade")
							}
							if rfar0.Y > rnear.Y {
								panic("Bad cascade")
							}
							if rfar1.X < rnear.X {
								panic("Bad cascade")
							}
							if rfar1.Y > rnear.Y {
								panic("Bad cascade")
							}
						*/
						var sBiLinear CascadeIntervalResult
						sBiLinear = BiLinear(primitives.Vec2{X: float64(0.25) + 0.5*float64(x&1), Y: float64(0.25) + 0.5*float64(y&1)}, si0, si1, si2, si3)

						siNear.mergeIntervals(&sBiLinear)
					}
					s.Add(&siNear)
				}
				s.Mul(factor)
				ciRNear.cascade[x][y][k] = s
			}
		}
	}

}

func (rc *RadianceCascadeVanilla) Render() {
	rc.s.Clear()

	for c := rc.cc.NCascades - 1; c >= 0; c-- {
		rc.Cascade(c)
	}
	/*
		{
			// TODO, debug
			c0 := cc.cascadeInfo[1]
			c0R := cascadeResult[1]
			for x := 0; x < 100; x++ {
				for k := 0; k < c0.dirCount; k++ {
					probe0 := cc.GetProbe(1, x, 0, k)
					probe1 := cc.GetProbe(1, x, c0.M, k)
					col0 := c0R.cascade[x][0][k].color
					col1 := c0R.cascade[x][c0.M][k].color
					if col0.R != col1.R || col0.G != col1.G || col0.B != col1.B {
						fmt.Println(x, k, probe0.ray.p.Y, probe1.ray.p.Y, col0, col1)
					}
				}
			}
		}
	*/
	rc.MergeOnImage()
}
