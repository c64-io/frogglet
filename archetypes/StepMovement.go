package archetypes

import "boxes/components"

type SteppedMover struct {
	*components.BasicMovementComponent
	*components.SteppedMovementComponent
	*components.QueuedMovementComponent
	*components.HeadingComponent
	*components.SpriteComponent
	*components.LocationComponent
}

type SteppedMoveable interface {
	GetSteppedMover() SteppedMover
}
