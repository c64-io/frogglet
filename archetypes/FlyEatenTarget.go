package archetypes

import "boxes/components"

type FlyEatenTarget struct {
	*components.ColliderComponent
}

type FlyEatenTargetable interface {
	GetFlyEatenTarget() FlyEatenTarget
}
