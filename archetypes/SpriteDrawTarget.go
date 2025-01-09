package archetypes

import "boxes/components"

type SpriteDrawTarget struct {
	*components.LocationComponent
	*components.SizeComponent
	*components.SpriteComponent
}

type SpriteDrawTargetable interface {
	GetSpriteDrawTarget() SpriteDrawTarget
}
