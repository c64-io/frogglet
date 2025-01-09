package archetypes

import "boxes/components"

type BasicSpriteSelectionTarget struct {
	*components.BasicMovementComponent
	*components.HeadingComponent
	*components.SpriteComponent
}

type BasicSpriteSelectionTargetable interface {
	GetBasicSpriteSelectionTarget() BasicSpriteSelectionTarget
}
