package main

type CascadeIntervalResult struct {
	color      Color
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
					color:      Color{R: 0, G: 0, B: 0},
					visibility: 0.0,
				}
			}
		}
	}
	return cr
}

func BiLinear(ratio Vec2, s0, s1, s2, s3 CascadeIntervalResult) CascadeIntervalResult {
	w1 := (1. - ratio.X) * (1. - ratio.Y)
	w2 := ratio.X * (1. - ratio.Y)
	w3 := (1. - ratio.X) * ratio.Y
	w4 := ratio.X * ratio.Y
	return CascadeIntervalResult{
		color: Color{
			R: s0.color.R*w1 + s1.color.R*w2 + s2.color.R*w3 + s3.color.R*w4,
			G: s0.color.G*w1 + s1.color.G*w2 + s2.color.G*w3 + s3.color.G*w4,
			B: s0.color.B*w1 + s1.color.B*w2 + s2.color.B*w3 + s3.color.B*w4,
		},
		visibility: s0.visibility*w1 + s1.visibility*w2 + s2.visibility*w3 + s3.visibility*w4,
	}
}

func Radiance(scene *Scene, cc *CascadeCalculator, cascade int, i int, j int, index int) CascadeIntervalResult {
	probe := cc.GetProbe(cascade, i, j, index)
	visibility, color := scene.Intersect(probe.ray, probe.tmax)
	return CascadeIntervalResult{
		color:      color,
		visibility: 1. - visibility, // 1. it hit nothing, 0. if hit
	}
}

func RenderCascade(scene *Scene) *SampledImage {
	const WIDTH, HEIGHT = 800, 800
	s := NewSampledImage(WIDTH, HEIGHT)
	cc := NewCascadeCalculator(WIDTH, HEIGHT)
	cascadeResult := make([]*CascadeResult, cc.NCascades+1)

	for c := 0; c < cc.NCascades+1; c++ {
		ci := cc.cascadeInfo[c]
		cascadeResult[c] = NewCascadeResult(ci)
	}

	// merge
	for c := cc.NCascades - 1; c >= 0; c-- {
		ciNear := cc.cascadeInfo[c]
		ciRNear := cascadeResult[c]

		ciFar := cc.cascadeInfo[c+1]
		ciRFar := cascadeResult[c+1]

		factor := float64(ciNear.dirCount) / float64(ciFar.dirCount) // integration factor
		nDirMerge := ciFar.dirCount / ciNear.dirCount                // number of directions to merge
		for x := 0; x < ciNear.N; x++ {
			for y := 0; y < ciNear.M; y++ {
				for k := 0; k < ciNear.dirCount; k++ {
					s := CascadeIntervalResult{}
					siNearTemp := Radiance(scene, cc, c, x, y, k)
					for kk := 0; kk < nDirMerge; kk++ {
						siNear := siNearTemp // copy radiance
						if c != cc.NCascades-1 {
							d := 4*k + kk // direction index in the next cascade
							si0 := ciRFar.cascade[(x>>1)+0][(y>>1)+0][d]
							si1 := ciRFar.cascade[(x>>1)+1][(y>>1)+0][d]
							si2 := ciRFar.cascade[(x>>1)+0][(y>>1)+1][d]
							si3 := ciRFar.cascade[(x>>1)+1][(y>>1)+1][d]
							var sBiLinear CascadeIntervalResult
							sBiLinear = BiLinear(Vec2{X: float64(0.33333) + 0.333333*float64(x&1), Y: float64(0.33333) + 0.333333*float64(y&1)}, si0, si1, si2, si3)

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
	// final merge of c0 to determine the pixel color
	c0 := cc.cascadeInfo[0]
	c0R := cascadeResult[0]
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			col := Color{R: 0, G: 0, B: 0}
			for k := 0; k < c0.dirCount; k++ {
				col.Add(c0R.cascade[x][y][k].color)
			}
			col.Mul(1.0 / float64(c0.dirCount)) // average the colors
			//col.Mul(2. * math.Pi)
			s.SetColor(x, y, col)
		}
	}
	return s
}
