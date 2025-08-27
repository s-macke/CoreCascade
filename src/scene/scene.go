package scene

import "CoreCascade/primitives"

type Scene interface {
	Intersect(r primitives.Ray, tmax float64) (visibility float64, c primitives.Color)
}
