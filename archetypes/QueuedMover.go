package archetypes

import "boxes/components"

type QueuedMover struct {
	*components.QueuedMovementComponent
}

type QueuedMoveable interface {
	GetQueuedMover() QueuedMover
}
