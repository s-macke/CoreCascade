package main

func main() {
	//img := NewSampledImageFromFile("output.raw")
	//img.StoreImage("output3.png")
	scene := NewScene()
	//img := RenderPathTracing(scene)
	img := RenderPathTracingParallel(scene)
	//img := RenderCascade(scene)
	img.Store("output")

	/*
		PlotSignedDistance()
		PlotCascade()
		PlotCascade2()
		PlotCascade3()
	*/
}
