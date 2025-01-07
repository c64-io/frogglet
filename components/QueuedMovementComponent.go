package components

import "boxes/utils"

type QueuedMovementComponent struct {
	EnqueuedMove                  utils.Facing
	MovementQueueLockoutCountdown float32
	QueueCooldown                 float32
}
