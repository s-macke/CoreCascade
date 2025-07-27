package radiance_cascade

import "CoreCascade/primitives"

type CascadeRadianceResult struct {
	color      primitives.Color
	visibility float64 // 1.0 if the interval hit nothing, and 0.0 it it did.
}

func NewCascadeRadianceResult() CascadeRadianceResult {
	return CascadeRadianceResult{
		color:      primitives.Black,
		visibility: 1.0, // Initially, it is assumed that the interval is fully visible.
	}
}

func (c *CascadeRadianceResult) mergeIntervals(far *CascadeRadianceResult) {
	c.color.R += far.color.R * c.visibility
	c.color.G += far.color.G * c.visibility
	c.color.B += far.color.B * c.visibility
	c.visibility *= far.visibility
}

func (c *CascadeRadianceResult) Mul(factor float64) {
	c.color.R *= factor
	c.color.G *= factor
	c.color.B *= factor
	c.visibility *= factor
}

func (c *CascadeRadianceResult) Add(b *CascadeRadianceResult) {
	c.color.R += b.color.R
	c.color.G += b.color.G
	c.color.B += b.color.B
	c.visibility += b.visibility
}
