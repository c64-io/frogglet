package archetypes

import "boxes/components"

type QueuedMoveTarget struct {
	*components.QueuedMovementComponent
}

type QueuedMoveTargetable interface {
	GetQueuedMoveTarget() QueuedMoveTarget
}
