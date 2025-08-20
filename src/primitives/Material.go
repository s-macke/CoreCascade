package primitives

type Material struct {
	Emissive   Color
	Absorption float64
}

func NewEmissiveMaterial(r, g, b float64) Material {
	return Material{
		Emissive:   Color{r, g, b},
		Absorption: 0,
	}
}

func NewEmissiveSRGBMaterial(r, g, b float64) Material {
	return Material{
		Emissive:   NewSRGBColor(r, g, b),
		Absorption: 0,
	}
}

func NewBlackMaterial() Material {
	return Material{
		Emissive:   Black,
		Absorption: 0,
	}
}
