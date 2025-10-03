package light_propagation_volumes

import (
	"CoreCascade2D/primitives"
	"CoreCascade2D/scene"
	"color"
	math "github.com/chewxy/math32"
	"linear_image"
	"vector"
)

type CircularHarmonics struct {
	b0, b1, b2 float32 // b0 = constant term, b1 = cos(theta), b2 = sin(theta)
}

type Cell struct {
	m       primitives.Material
	r, g, b CircularHarmonics // // emitted r,g,b light as circular harmonics
}

func ToImage(grid [][]Cell, image *linear_image.SampledImage) {
	b0_base := math.Sqrt(1. / (2. * math.Pi))
	for i := range grid {
		for j := range grid[i] {
			// simple average of the harmonics
			c := grid[i][j]
			image.SetColor(j, i, color.Color{
				R: c.r.b0 * b0_base,
				G: c.g.b0 * b0_base,
				B: c.b.b0 * b0_base,
			})
		}
	}
}

func ToGrid(scene scene.Scene, image *linear_image.SampledImage) [][]Cell {
	var grid [][]Cell
	grid = make([][]Cell, image.Height)
	for i := range grid {
		grid[i] = make([]Cell, image.Width)
		for j := range grid[i] {
			grid[i][j] = Cell{m: scene.GetMaterial(vector.Vec2{
				X: (float32(j)/float32(image.Width))*2 - 1,
				Y: (float32(i)/float32(image.Height))*2 - 1})}
		}
	}
	return grid
}

func PropagateFrom(n vector.Vec2, sa float32, src *Cell, dest *Cell) {
	absorption := math.Exp(-src.m.Absorption*2./800.*0.5) * math.Exp(-dest.m.Absorption*2./800.*0.5)
	dV0 := math.Sqrt(1. / (2. * math.Pi))
	dV1 := n.X * math.Sqrt(2./(2.*math.Pi)) // projected solid angle in the x direction
	dV2 := n.Y * math.Sqrt(2./(2.*math.Pi))

	Lr := max(0.0, src.r.b0*dV0+src.r.b1*dV1+src.r.b2*dV2) * sa * absorption
	Lg := max(0.0, src.g.b0*dV0+src.g.b1*dV1+src.g.b2*dV2) * sa * absorption
	Lb := max(0.0, src.b.b0*dV0+src.b.b1*dV1+src.b.b2*dV2) * sa * absorption

	dest.r.b0 += dV0 * Lr
	dest.r.b1 += dV1 * Lr
	dest.r.b2 += dV2 * Lr

	dest.g.b0 += dV0 * Lg
	dest.g.b1 += dV1 * Lg
	dest.g.b2 += dV2 * Lg

	dest.b.b0 += dV0 * Lb
	dest.b.b1 += dV1 * Lb
	dest.b.b2 += dV2 * Lb
}

func Propagate(grid [][]Cell, grid2 [][]Cell, width, height int) {

	// TODO: the comments are wrong
	const sa1 = 0.6435011087932845                                    // ~53.13°, projecting to the right side of our pixel
	const sa2 = 0.4636476090008066                                    // ~18.43°, projecting to the top/bottom side of our pixel
	dn1 := vector.Vec2{X: 0.8944271909999159, Y: 0.4472135954999579}  // normal to the center of the top side of our pixel
	dn0 := vector.Vec2{X: 0.8944271909999159, Y: -0.4472135954999579} // normal to the center of the bottom side of our pixel

	cellWidth := 2. / float32(width) // assuming the scene is from -1 to 1 in both x and y
	var src *Cell
	for j := 1; j < height-1; j++ {
		for i := 1; i < width-1; i++ {
			dst := &grid2[j][i]
			dst.r.b0 = 0.
			dst.r.b1 = 0.
			dst.r.b2 = 0.
			dst.g.b0 = 0.
			dst.g.b1 = 0.
			dst.g.b2 = 0.
			dst.b.b0 = 0.
			dst.b.b1 = 0.
			dst.b.b2 = 0.

			src = &grid[j][i-1]
			PropagateFrom(vector.Vec2{X: 1, Y: 0}, sa1, src, dst)
			PropagateFrom(vector.Vec2{X: dn0.X, Y: dn0.Y}, sa2, src, dst)
			PropagateFrom(vector.Vec2{X: dn1.X, Y: dn1.Y}, sa2, src, dst)

			src = &grid[j][i+1]
			PropagateFrom(vector.Vec2{X: -1, Y: 0}, sa1, src, dst)
			PropagateFrom(vector.Vec2{X: -dn0.X, Y: -dn0.Y}, sa2, src, dst)
			PropagateFrom(vector.Vec2{X: -dn1.X, Y: -dn1.Y}, sa2, src, dst)

			src = &grid[j-1][i]
			PropagateFrom(vector.Vec2{X: 0, Y: 1}, sa1, src, dst)
			PropagateFrom(vector.Vec2{X: dn0.Y, Y: dn0.X}, sa2, src, dst)
			PropagateFrom(vector.Vec2{X: dn1.Y, Y: dn1.X}, sa2, src, dst)

			src = &grid[j+1][i]
			PropagateFrom(vector.Vec2{X: 0, Y: -1}, sa1, src, dst)
			PropagateFrom(vector.Vec2{X: -dn0.Y, Y: -dn0.X}, sa2, src, dst)
			PropagateFrom(vector.Vec2{X: -dn1.Y, Y: -dn1.X}, sa2, src, dst)

			dst.r.b0 += dst.m.Emissive.R * cellWidth
			dst.g.b0 += dst.m.Emissive.G * cellWidth
			dst.b.b0 += dst.m.Emissive.B * cellWidth
		}
	}

	for j := 1; j < height-1; j++ {
		for i := 1; i < width-1; i++ {
			grid[j][i].r = grid2[j][i].r
			grid[j][i].g = grid2[j][i].g
			grid[j][i].b = grid2[j][i].b
		}
	}

}

func LightPropagationVolume(scene scene.Scene, image *linear_image.SampledImage) {
	grid := ToGrid(scene, image)
	grid2 := ToGrid(scene, image)

	//iterations := max(image.Width, image.Height) * 2.82 // light has the chance to bounce from one side to the other and back
	iterations := int(float32(max(image.Width, image.Height)) * 2.) // light has the chance to bounce from one side to the other

	for it := 0; it < iterations; it++ {
		Propagate(grid, grid2, image.Width, image.Height)
	}
	ToImage(grid, image)
}
