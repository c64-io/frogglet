package archetypes

import "boxes/components"

type CollisionBoxHighlightTarget struct {
	*components.ColliderComponent
	*components.ColorComponent
}

type CollisionBoxHighlightTargetable interface {
	GetCollisionBoxHighlightTarget() CollisionBoxHighlightTarget
}
