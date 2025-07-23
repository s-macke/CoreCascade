package main

func main() {
	scene := NewScene(Scene2)

	//img := RenderPathTracing(scene)
	//img := RenderPathTracingParallel(scene, 100)
	img := NewRadianceCascadeVanilla(scene).Render()
	img.Store("output")
	/*
		truth := NewSampledImageFromFile("ring_shadow.raw")
		truth.Error(img)
		truth.StoreImage("diff.png")
	*/
	/*
		PlotSignedDistance()
		PlotCascade()
		PlotCascade2()
		PlotCascade3()
		PlotCascade4()
		PlotCascade5()
	*/
	PlotEnergy(img)

}
