package archetypes

import "boxes/components"

type DrawRectTarget struct {
	*components.LocationComponent
	*components.SizeComponent
	*components.ColorComponent
}

type DrawRectTargetable interface {
	GetDrawRectTarget() DrawRectTarget
}
