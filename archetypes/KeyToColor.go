package archetypes

import "boxes/components"

type KeyToColor struct {
	*components.ColorComponent
	*components.TargetKeyComponent
}

type KeyToColorTarget interface {
	GetKeyToColorTarget() KeyToColor
}
