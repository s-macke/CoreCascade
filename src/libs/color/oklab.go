package color

import "math"

type Oklab struct {
	L, A, B float64
}

func NewRainbowOklabToLinear(t float64) Color {
	// t in [0, 1]
	angle := t * 2. * math.Pi
	lab := Oklab{
		L: 1.,
		A: 0.4 * math.Cos(angle),
		B: 0.4 * math.Sin(angle),
	}
	return lab.ToLinear()
}

func FromLinearToOkLab(c Color) Oklab {
	l := 0.4122214708*c.R + 0.5363325363*c.G + 0.0514459929*c.B
	m := 0.2119034982*c.R + 0.6806995451*c.G + 0.1073969566*c.B
	s := 0.0883024619*c.R + 0.2817188376*c.G + 0.6299787005*c.B

	l_ := math.Pow(l, 1./3.)
	m_ := math.Pow(m, 1./3.)
	s_ := math.Pow(s, 1./3.)

	return Oklab{
		L: 0.2104542553*l_ + 0.7936177850*m_ - 0.0040720468*s_,
		A: 1.9779984951*l_ - 2.4285922050*m_ + 0.4505937099*s_,
		B: 0.0259040371*l_ + 0.7827717662*m_ - 0.8086757660*s_,
	}
}

// https://blog.pkh.me/p/43-the-current-technology-is-not-ready-for-proper-blending.html
func (c *Oklab) ToLinear() Color {
	l := c.L
	a := c.A
	b := c.B
	l_ := l + 0.3963377774*a + 0.2158037573*b
	m_ := l - 0.1055613458*a - 0.0638541728*b
	s_ := l - 0.0894841775*a - 1.2914855480*b

	l = l_ * l_ * l_
	m := m_ * m_ * m_
	s := s_ * s_ * s_

	return Color{
		R: +4.0767416621*l - 3.3077115913*m + 0.2309699292*s,
		G: -1.2684380046*l + 2.6097574011*m - 0.3413193965*s,
		B: -0.0041960863*l - 0.7034186147*m + 1.7076147010*s,
	}
}
