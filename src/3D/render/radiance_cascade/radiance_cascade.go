package radiance_cascade

import (
	"CoreCascade3D/scene"
	"color"
	"fmt"
	math "github.com/chewxy/math32"
	"linear_image"
	"vector"
)

type Cascade struct {
	info     CascadeInfo
	radiance [][][][][]CascadeRadianceResult // x, y, direction (directionLast)
}

func NewCascade(ci CascadeInfo) *Cascade {
	cr := &Cascade{
		info: ci,
	}
	cr.radiance = make([][][][][]CascadeRadianceResult, ci.N)
	for i := 0; i < ci.N; i++ {
		cr.radiance[i] = make([][][][]CascadeRadianceResult, ci.M)
		for j := 0; j < ci.M; j++ {
			cr.radiance[i][j] = make([][][]CascadeRadianceResult, ci.O)
			for k := 0; k < ci.O; k++ {
				cr.radiance[i][j][k] = make([][]CascadeRadianceResult, ci.DirCount)
				for l := 0; l < ci.DirCount; l++ {
					cr.radiance[i][j][k][l] = make([]CascadeRadianceResult, ci.DirCount)
					for m := 0; m < ci.DirCount; m++ {
						cr.radiance[i][j][k][l][m] = NewCascadeRadianceResult()
					}
				}
			}
		}
	}
	return cr
}

type RadianceCascade struct {
	width, height, depth int
	scene                scene.Scene
	image                *linear_image.SampledImage
	layers               []*linear_image.SampledImage
	cc                   *CascadeCalculator
	cascades             []*Cascade
}

func NewRadianceCascade(s scene.Scene, image *linear_image.SampledImage) *RadianceCascade {
	depth := 40
	var layers []*linear_image.SampledImage
	for layer := 0; layer < depth; layer++ {
		layers = append(layers, linear_image.NewSampledImage(image.Width, image.Height))
	}
	rc := &RadianceCascade{
		width:  image.Width,
		height: image.Height,
		depth:  depth,
		scene:  s,
		cc:     NewCascadeCalculator(image.Width, image.Height, depth),
		image:  image,
		layers: layers,
	}
	rc.cc.Info()
	rc.cascades = make([]*Cascade, rc.cc.NCascades+1)

	for c := 0; c < rc.cc.NCascades+1; c++ {
		ci := rc.cc.CascadeInfo[c]
		rc.cascades[c] = NewCascade(ci)
	}

	return rc
}

func TriLinear(ratio vector.Vec3, s0, s1, s2, s3, s4, s5, s6, s7 CascadeRadianceResult) CascadeRadianceResult {
	w0 := (1. - ratio.X) * (1. - ratio.Y) * (1. - ratio.Z)
	w1 := ratio.X * (1. - ratio.Y) * (1. - ratio.Z)
	w2 := (1. - ratio.X) * ratio.Y * (1. - ratio.Z)
	w3 := ratio.X * ratio.Y * (1. - ratio.Z)
	w4 := (1. - ratio.X) * (1. - ratio.Y) * ratio.Z
	w5 := ratio.X * (1. - ratio.Y) * ratio.Z
	w6 := (1. - ratio.X) * ratio.Y * ratio.Z
	w7 := ratio.X * ratio.Y * ratio.Z

	return CascadeRadianceResult{
		color: color.Color{
			R: s0.color.R*w0 + s1.color.R*w1 + s2.color.R*w2 + s3.color.R*w3 + s4.color.R*w4 + s5.color.R*w5 + s6.color.R*w6 + s7.color.R*w7,
			G: s0.color.G*w0 + s1.color.G*w1 + s2.color.G*w2 + s3.color.G*w3 + s4.color.G*w4 + s5.color.G*w5 + s6.color.G*w6 + s7.color.G*w7,
			B: s0.color.B*w0 + s1.color.B*w1 + s2.color.B*w2 + s3.color.B*w3 + s4.color.B*w4 + s5.color.B*w5 + s6.color.B*w6 + s7.color.B*w7,
		},
		visibility: s0.visibility*w0 + s1.visibility*w1 + s2.visibility*w2 + s3.visibility*w3 + s4.visibility*w4 + s5.visibility*w5 + s6.visibility*w6 + s7.visibility*w7,
	}
}

// IndexToSceneUVW from (-1, -1, 0) to (1, 1, 0.1)
func (rc *RadianceCascade) IndexToSceneUVW(x, y, z int) vector.Vec3 {
	return vector.Vec3{
		X: (float32(x)/float32(rc.width))*2. - 1.,
		Y: (float32(y)/float32(rc.height))*2. - 1.,
		Z: float32(z) / float32(rc.depth) * 0.1,
	}
}

func (rc *RadianceCascade) MergeOnImage() {
	rc.image.Clear()
	cascade0 := rc.cascades[0]
	dz := float32(0.1) / float32(rc.depth)
	for y := 0; y < rc.height; y++ {
		for x := 0; x < rc.width; x++ {
			visibility := float32(1.0)
			col := color.Black
			for z := rc.depth - 1; z >= 0; z-- {
				uvw := rc.IndexToSceneUVW(x, y, z)
				mat := rc.scene.GetMaterial(uvw)

				fluence := color.Black
				for k := 0; k < cascade0.info.DirCount; k++ {
					for l := 0; l < cascade0.info.DirCount; l++ {
						fluence.Add(cascade0.radiance[x][y][z][k][l].color)
					}
				}
				fluence.Div(float32(cascade0.info.DirCount * cascade0.info.DirCount)) // average the colors

				col.R += fluence.R * mat.Diffuse.R * visibility
				col.G += fluence.G * mat.Diffuse.G * visibility
				col.B += fluence.B * mat.Diffuse.B * visibility
				visibility *= math.Exp(-mat.Absorption * dz)
			}
			rc.image.SetColor(x, y, col)
		}
	}
}

func (rc *RadianceCascade) Radiance(probe CascadeProbe) CascadeRadianceResult {
	visibility, col := rc.scene.Trace(probe.Ray, probe.Tmax)
	return CascadeRadianceResult{
		color:      col,
		visibility: visibility,
	}
}

func (rc *RadianceCascade) CascadeMerge(cNear *Cascade, cFar *Cascade, x int, y int, z int, k int, l int) {
	ciNear := cNear.info
	ciFar := cFar.info

	//factor := float32(ciNear.DirCount*ciNear.DirCount) / float32(ciFar.DirCount*ciFar.DirCount) // integration factor
	factor := float32(ciNear.DirCount) / float32(ciFar.DirCount) // integration factor
	nDirMerge := ciFar.DirCount / ciNear.DirCount                // number of directions to merge
	trilinearWeights := vector.Vec3{
		X: 0.25 + 0.5*float32(x&1),
		Y: 0.25 + 0.5*float32(y&1),
		Z: 0.25 + 0.5*float32(z&1),
	}
	cNear.radiance[x][y][z][k][l] = rc.Radiance(ciNear.GetProbe(x, y, z, k, l))
	if !ciFar.isLast { // is not last cascade then merging needed
		merged := NewCascadeRadianceResult()

		for ll := 0; ll < nDirMerge; ll++ { // merge the four directions
			for kk := 0; kk < nDirMerge; kk++ { // merge the four directions
				d := 2*k + kk
				e := 2*l + ll
				si0 := cFar.radiance[(x>>1)+0][(y>>1)+0][(z>>1)+0][d][e]
				si1 := cFar.radiance[(x>>1)+1][(y>>1)+0][(z>>1)+0][d][e]
				si2 := cFar.radiance[(x>>1)+0][(y>>1)+1][(z>>1)+0][d][e]
				si3 := cFar.radiance[(x>>1)+1][(y>>1)+1][(z>>1)+0][d][e]
				si4 := cFar.radiance[(x>>1)+0][(y>>1)+0][(z>>1)+1][d][e]
				si5 := cFar.radiance[(x>>1)+1][(y>>1)+0][(z>>1)+1][d][e]
				si6 := cFar.radiance[(x>>1)+0][(y>>1)+1][(z>>1)+1][d][e]
				si7 := cFar.radiance[(x>>1)+1][(y>>1)+1][(z>>1)+1][d][e]
				sTriLinear := TriLinear(trilinearWeights, si0, si1, si2, si3, si4, si5, si6, si7)
				merged.Add(&sTriLinear)
			}
		}
		merged.Mul(factor)
		cNear.radiance[x][y][z][k][l].mergeIntervals(&merged) // merge the radiance from the near cascade
	}
}

func (rc *RadianceCascade) Render() {
	for c := rc.cc.NCascades - 1; c >= 0; c-- {
		cNear := rc.cascades[c]
		cFar := rc.cascades[c+1]
		fmt.Println("Calculating merge from cascade", cFar.info.Cascade, "to", cNear.info.Cascade)
		for x := 0; x < cNear.info.N; x++ {
			for y := 0; y < cNear.info.M; y++ {
				for z := 0; z < cNear.info.O; z++ {
					for k := 0; k < cNear.info.DirCount; k++ {
						for l := 0; l < cNear.info.DirCount; l++ {
							rc.CascadeMerge(cNear, cFar, x, y, z, k, l)
						}
					}
				}
			}
		}
	}
	rc.MergeOnImage()
}
