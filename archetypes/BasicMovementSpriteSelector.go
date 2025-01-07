package archetypes

import "boxes/components"

type BasicMovementSpriteSelector struct {
	*components.BasicMovementComponent
	*components.HeadingComponent
	*components.SpriteComponent
}

type BasicMovementSpriteSelectable interface {
	GetBasicMovementSpriteSelector() BasicMovementSpriteSelector
}
