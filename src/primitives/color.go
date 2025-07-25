package primitives

import (
	"image/color"
	"math"
)

type Color struct {
	R, G, B float64
}

var Black = Color{R: 0, G: 0, B: 0}

func NewSRGBColor(r float64, g float64, b float64) (c Color) {
	c = Color{r, g, b}
	c.R = math.Pow(c.R, 2.2)
	c.G = math.Pow(c.G, 2.2)
	c.B = math.Pow(c.B, 2.2)
	return c
}

func (c *Color) Add(other Color) {
	c.R += other.R
	c.G += other.G
	c.B += other.B
}

func (c *Color) Sub(other Color) {
	c.R -= other.R
	c.G -= other.G
	c.B -= other.B
}

func (c *Color) Abs() {
	c.R = math.Abs(c.R)
	c.G = math.Abs(c.G)
	c.B = math.Abs(c.B)
}

func (c *Color) Div(f float64) {
	invA := 1. / f
	c.R *= invA
	c.G *= invA
	c.B *= invA
}

func (c *Color) Mul(f float64) {
	c.R *= f
	c.G *= f
	c.B *= f
}

func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func (c *Color) ToRGBA() color.RGBA {
	return color.RGBA{
		R: uint8(clamp(c.R*255, 0, 255)),
		G: uint8(clamp(c.G*255, 0, 255)),
		B: uint8(clamp(c.B*255, 0, 255)),
		A: 255,
	}
}

// Intensity calculates the magnitude (or "energy") of the color.
func (c *Color) Intensity() float64 {
	return math.Sqrt(c.R*c.R + c.G*c.G + c.B*c.B)
}

/*
// https://knarkowicz.wordpress.com/2016/01/06/aces-filmic-tone-mapping-curve/
// https://github.com/ampas/aces-core/

	func (c *Color) ToSimpleAcesSRGBA() color.RGBA {
		const a = 2.51
		const b = 0.03
		const cc = 2.43
		const d = 0.59
		const e = 0.14

		return color.RGBA{
			R: uint8(clamp((c.R*(a*c.R+b))/(c.R*(cc*c.R+d)+e), 0., 1.) * 255.),
			G: uint8(clamp((c.G*(a*c.G+b))/(c.G*(cc*c.G+d)+e), 0., 1.) * 255.),
			B: uint8(clamp((c.B*(a*c.B+b))/(c.B*(cc*c.B+d)+e), 0., 1.) * 255.),
			A: 255,
		}
	}

// https://github.com/TheRealMJP/BakingLab/blob/master/BakingLab/ACES.hlsl
// ToAcesSRGBA applies the ACES tone mapping to the color and converts it to sRGBA.

	func (c *Color) ToAcesSRGBA() color.RGBA {
		// ACES input matrix (m1 in GLSL)
		m1 := [3][3]float64{
			{0.59719, 0.35458, 0.04823},
			{0.07600, 0.90834, 0.01566},
			{0.02840, 0.13383, 0.83777},
		}

		// ACES output matrix (m2 in GLSL)
		m2 := [3][3]float64{
			{1.60475, -0.53108, -0.07367},
			{-0.10208, 1.10813, -0.00605},
			{-0.00327, -0.07276, 1.07602},
		}

		// Transform color to ACEScg (v in GLSL)
		v := &Color{
			R: c.R*m1[0][0] + c.G*m1[0][1] + c.B*m1[0][2],
			G: c.R*m1[1][0] + c.G*m1[1][1] + c.B*m1[1][2],
			B: c.R*m1[2][0] + c.G*m1[2][1] + c.B*m1[2][2],
		}

		// Apply the RRT and ODT fit
		a := &Color{
			R: v.R*(v.R+0.0245786) - 0.000090537,
			G: v.G*(v.G+0.0245786) - 0.000090537,
			B: v.B*(v.B+0.0245786) - 0.000090537,
		}

		b := &Color{
			R: v.R*(0.983729*v.R+0.4329510) + 0.238081,
			G: v.G*(0.983729*v.G+0.4329510) + 0.238081,
			B: v.B*(0.983729*v.B+0.4329510) + 0.238081,
		}

		// Divide a by b
		div := &Color{
			R: a.R / b.R,
			G: a.G / b.G,
			B: a.B / b.B,
		}

		// Transform back to sRGB
		aces := &Color{
			R: div.R*m2[0][0] + div.G*m2[0][1] + div.B*m2[0][2],
			G: div.R*m2[1][0] + div.G*m2[1][1] + div.B*m2[1][2],
			B: div.R*m2[2][0] + div.G*m2[2][1] + div.B*m2[2][2],
		}

		// Clamp the values between 0 and 1
		clamped := &Color{
			R: math.Max(0.0, math.Min(1.0, aces.R)),
			G: math.Max(0.0, math.Min(1.0, aces.G)),
			B: math.Max(0.0, math.Min(1.0, aces.B)),
		}

		// Gamma correction (1.0 / 2.2)
		gammaCorrected := &Color{
			R: math.Pow(clamped.R, 1.0/2.2),
			G: math.Pow(clamped.G, 1.0/2.2),
			B: math.Pow(clamped.B, 1.0/2.2),
		}

		return color.RGBA{
			R: uint8(gammaCorrected.R * 255.0),
			G: uint8(gammaCorrected.G * 255.0),
			B: uint8(gammaCorrected.B * 255.0),
			A: 255,
		}
	}
*/

func linearToSRGB(x float64) float64 {
	if x <= 0.0031308 {
		return 12.92 * x
	}
	return 1.055*math.Pow(x, 1/2.4) - 0.055
}

func (c *Color) ToSRGBA() color.RGBA {
	newcol := Color{
		R: math.Pow(c.R, 1.0/2.2),
		G: math.Pow(c.G, 1.0/2.2),
		B: math.Pow(c.B, 1.0/2.2),
	}

	return color.RGBA{
		R: uint8(clamp(newcol.R*255, 0, 255)),
		G: uint8(clamp(newcol.G*255, 0, 255)),
		B: uint8(clamp(newcol.B*255, 0, 255)),
		A: 255,
	}
}
