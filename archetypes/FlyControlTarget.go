package archetypes

import "boxes/components"

type FlyControlTarget struct {
	*components.SpriteComponent
	*components.LocationComponent
}

type FlyControlTargetable interface {
	GetFlyControlTarget() FlyControlTarget
}
