package archetypes

import "boxes/components"

type SpriteDrawer struct {
	*components.LocationComponent
	*components.SizeComponent
	*components.SpriteComponent
}

type SpriteDrawable interface {
	GetSpriteDrawer() SpriteDrawer
}
