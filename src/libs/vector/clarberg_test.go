package vector

import (
	"fmt"
	"testing"
)

func PrintArrow(u float32, v float32, r0 float32, r1 float32, color string) {
	//dir := EqualAreaSquareToSphere(FoldFixup(u, v))
	dir := ClarbergToSphere(u, v)
	fmt.Printf("AddArrow(%f, %f, %f, %f, %f, %f, \"%s\");\n", dir.X*r0, dir.Y*r0, dir.Z*r0, dir.X*r1, dir.Y*r1, dir.Z*r1, color)
}

func PlotCascade(L, I, J int) {
	u, v := TileCenterUV(Tile{L: L, I: I, J: J})
	color := "blue"
	if L > 1 {
		color = "red"
	}
	PrintArrow(u, v, float32(L-1), float32(L), color)
}

func TestClarberg(t *testing.T) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			//fmt.Println(i, j)
			fmt.Println(TileCenterUV(Tile{L: 1, I: i, J: j}))
			//PlotCascade(1, i, j)
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			//fmt.Println(i, j)
			fmt.Println(TileCenterUV(Tile{L: 2, I: i, J: j}))
			//PlotCascade(2, i, j)
		}
	}
}
