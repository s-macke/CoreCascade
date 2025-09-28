package scene

import (
	"CoreCascade3D/primitives"
	"color"
	"vector"
)

type Scene interface {
	GetMaterial(p vector.Vec3) primitives.Material
	Trace(r vector.Ray3D, tmax float64) (visibility float64, c color.Color)
}
