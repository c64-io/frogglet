package archetypes

import "boxes/components"

type DebugStringOverlayTarget struct {
	*components.LocationComponent
	*components.SizeComponent
	*components.SpriteComponent
	*components.DebugStringComponent
}

type DebugStringOverlayTargetable interface {
	GetDebugStringOverlayTarget() DebugStringOverlayTarget
}
