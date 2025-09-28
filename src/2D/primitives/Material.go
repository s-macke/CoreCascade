package primitives

import (
	"color"
	"vector"
)

type Material struct {
	Emissive          color.Color
	DirectionEnabled  bool
	EmissiveDirection vector.Vec2
	EmissiveAngle     float64

	Absorption float64
	Diffuse    color.Color // Diffuse color for SRGB materials,
}

func (m *Material) Merge(other *Material) {
	m.Emissive.Add(other.Emissive)
	m.DirectionEnabled = other.DirectionEnabled
	m.EmissiveDirection = other.EmissiveDirection
	m.EmissiveAngle = other.EmissiveAngle
	m.Absorption += other.Absorption
	m.Diffuse.Add(other.Diffuse)
}

var VoidMaterial = Material{
	Emissive:   color.Black,
	Diffuse:    color.Black,
	Absorption: 0,
}

func (m *Material) Emission(t vector.Vec2) color.Color {
	if !m.DirectionEnabled {
		return m.Emissive
	}
	dot := t.X*m.EmissiveDirection.X + t.Y*m.EmissiveDirection.Y
	if dot > m.EmissiveAngle {
		return m.Emissive
	} else {
		return color.Black
	}
}

func NewEmissiveMaterial(r, g, b float64) Material {
	return Material{
		DirectionEnabled: false,
		Emissive:         color.Color{r, g, b},
		Absorption:       0,
		Diffuse:          color.Black,
	}
}

func NewEmissiveSRGBMaterial(r, g, b float64) Material {
	return Material{
		DirectionEnabled: false,
		Emissive:         color.NewSRGBColor(r, g, b),
		Diffuse:          color.Black,
		Absorption:       0,
	}
}

func NewBlackMaterial() Material {
	return Material{
		DirectionEnabled: false,
		Emissive:         color.Black,
		Diffuse:          color.Black,
		Absorption:       1.,
	}
}

func NewAbsorbiveMaterial(value float64, r, g, b float64) Material {
	return Material{
		DirectionEnabled: false,
		Emissive:         color.Black,
		Absorption:       value,
		Diffuse:          color.Color{R: r, G: g, B: b},
	}
}
