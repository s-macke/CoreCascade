package color

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

func linearToSRGB(x float64) float64 {
	if x <= 0.0031308 {
		return 12.92 * x
	}
	return 1.055*math.Pow(x, 1/2.4) - 0.055
}

// only Gamma correction
func (c *Color) ToSRGBAOnlyGamma() color.RGBA {
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

// https://64.github.io/tonemapping/
func (c *Color) ToSRGBAReinhard() color.RGBA {
	newcol := Color{
		R: math.Pow(c.R/(c.R+1.), 1.0/2.2),
		G: math.Pow(c.G/(c.G+1.), 1.0/2.2),
		B: math.Pow(c.B/(c.B+1.), 1.0/2.2),
	}

	return color.RGBA{
		R: uint8(newcol.R * 255),
		G: uint8(newcol.G * 255),
		B: uint8(newcol.B * 255),
		A: 255,
	}
}

func (c *Color) ToSRGBAFilmic() color.RGBA {
	newcol := Color{
		R: 1. - 1./math.Pow(1.+c.R, 2.2),
		G: 1. - 1./math.Pow(1.+c.G, 2.2),
		B: 1. - 1./math.Pow(1.+c.B, 2.2),
	}

	return color.RGBA{
		R: uint8(newcol.R * 255),
		G: uint8(newcol.G * 255),
		B: uint8(newcol.B * 255),
		A: 255,
	}
}
