package main

func main() {
	scene := NewScene()
	img := RenderPathTracing(scene)
	//img := RenderPathTracingParallel(scene)
	//img := RenderCascade(scene)
	StoreImage(img, "output.png")

	PlotCascade()
	PlotCascade2()
	PlotCascade3()
	/*
		for x := -20.0; x <= 20.0; x += 0.5 {
			for y := -20.0; y <= 20.0; y += 0.5 {
				p := Vec2{X: x, Y: y}
				fmt.Println(x, y, scene.sd(p))
			}
			fmt.Println()
		}
	*/
}
