package archetypes

import "boxes/components"

type DebugStringOverlay struct {
	*components.LocationComponent
	*components.SizeComponent
	*components.SpriteComponent
	*components.DebugStringComponent
}

type DebugStringOverlayable interface {
	GetDebugStringOverlay() DebugStringOverlay
}
