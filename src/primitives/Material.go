package primitives

type Material struct {
	Emissive   Color
	Absorption float64
	Diffuse    Color // Diffuse color for SRGB materials,
}

func (m *Material) Merge(other *Material) {
	m.Emissive.Add(other.Emissive)
	m.Absorption += other.Absorption
	m.Diffuse.Add(other.Diffuse)
}

var VoidMaterial = Material{
	Emissive:   Black,
	Diffuse:    Black,
	Absorption: 0,
}

func NewEmissiveMaterial(r, g, b float64) Material {
	return Material{
		Emissive:   Color{r, g, b},
		Absorption: 0,
		Diffuse:    Black,
	}
}

func NewEmissiveSRGBMaterial(r, g, b float64) Material {
	return Material{
		Emissive:   NewSRGBColor(r, g, b),
		Diffuse:    Black,
		Absorption: 0,
	}
}

func NewBlackMaterial() Material {
	return Material{
		Emissive:   Black,
		Diffuse:    Black,
		Absorption: 1.,
	}
}

func NewAbsorbiveMaterial(value float64, r, g, b float64) Material {
	return Material{
		Emissive:   Black,
		Absorption: value,
		Diffuse:    Color{r, g, b},
	}
}
