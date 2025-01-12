package archetypes

import "boxes/components"

type FlySpawnTarget struct {
	*components.SpriteComponent
	*components.LocationComponent
}

type FlySpawnTargetable interface {
	GetFlySpawnTarget() FlySpawnTarget
}
