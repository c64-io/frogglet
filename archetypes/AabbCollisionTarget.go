package archetypes

import "boxes/components"

type AabbCollisionTarget struct {
	*components.LocationComponent
	*components.SizeComponent
	*components.ColliderComponent
}

type AabbCollisionTargetable interface {
	GetAabbCollisionTarget() AabbCollisionTarget
}
