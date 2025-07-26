package vanilla_rc

import "CoreCascade/primitives"

type CascadeProbeResult struct {
	color      primitives.Color
	visibility float64 // 1.0 if the interval hit nothing, and 0.0 it it did.
}

func NewCascadeProbeResult() CascadeProbeResult {
	return CascadeProbeResult{
		color:      primitives.Black,
		visibility: 1.0, // Initially, it is assumed that the interval is fully visible.
	}
}

func (c *CascadeProbeResult) mergeIntervals(far *CascadeProbeResult) {
	c.color.R += far.color.R * c.visibility
	c.color.G += far.color.G * c.visibility
	c.color.B += far.color.B * c.visibility
	c.visibility *= far.visibility
}

func (c *CascadeProbeResult) Mul(factor float64) {
	c.color.R *= factor
	c.color.G *= factor
	c.color.B *= factor
	c.visibility *= factor
}

func (c *CascadeProbeResult) Add(b *CascadeProbeResult) {
	c.color.R += b.color.R
	c.color.G += b.color.G
	c.color.B += b.color.B
	c.visibility += b.visibility
}
