package scene

import "CoreCascade/primitives"

type Scene interface {
	GetMaterial(p primitives.Vec2) primitives.Material
	Intersect(r primitives.Ray, tmax float64) (visibility float64, c primitives.Color)
}
