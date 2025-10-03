package scene

import (
	"CoreCascade2D/primitives"
	"color"
	"vector"
)

type Scene interface {
	GetMaterial(p vector.Vec2) primitives.Material
	Trace(r vector.Ray2D, tmax float32) (visibility float32, c color.Color)
}
