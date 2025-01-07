package archetypes

import "boxes/components"

type DrawRect struct {
	*components.LocationComponent
	*components.SizeComponent
	*components.ColorComponent
}

type DrawRectTarget interface {
	GetDrawRect() DrawRect
}
