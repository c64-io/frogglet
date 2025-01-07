package archetypes

import "boxes/components"

type FlyControl struct {
	*components.SpriteComponent
	*components.LocationComponent
}

type FlyControllable interface {
	GetFlyControl() FlyControl
}
