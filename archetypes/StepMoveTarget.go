package archetypes

import "boxes/components"

type SteppedMoveTarget struct {
	*components.BasicMovementComponent
	*components.SteppedMovementComponent
	*components.QueuedMovementComponent
	*components.HeadingComponent
	*components.SpriteComponent
	*components.LocationComponent
}

type SteppedMoveTargetable interface {
	GetSteppedMoveTarget() SteppedMoveTarget
}
