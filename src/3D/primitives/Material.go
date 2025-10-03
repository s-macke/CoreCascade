package primitives

import (
	"color"
)

type Material struct {
	Emissive   color.Color
	Absorption float32
	Diffuse    color.Color // Diffuse color for SRGB materials,
}

func (m *Material) Merge(other *Material) {
	m.Emissive.Add(other.Emissive)
	m.Absorption += other.Absorption
	m.Diffuse.Add(other.Diffuse)
}

var VoidMaterial = Material{
	Emissive:   color.Black,
	Diffuse:    color.Black,
	Absorption: 0,
}

func NewEmissiveMaterial(r, g, b float32) Material {
	return Material{
		Emissive:   color.Color{R: r, G: g, B: b},
		Absorption: 0,
		Diffuse:    color.Black,
	}
}

func NewEmissiveSRGBMaterial(r, g, b float32) Material {
	return Material{
		Emissive:   color.NewSRGBColor(r, g, b),
		Diffuse:    color.Black,
		Absorption: 0,
	}
}

func NewBlackMaterial() Material {
	return Material{
		Emissive:   color.Black,
		Diffuse:    color.Black,
		Absorption: 1.,
	}
}

func NewAbsorbiveMaterial(value float32, r, g, b float32) Material {
	return Material{
		Emissive:   color.Black,
		Absorption: value,
		Diffuse:    color.Color{R: r, G: g, B: b},
	}
}
