package systems

import (
	"boxes/archetypes"
	"boxes/engine"
	"boxes/singletons"
	"boxes/utils"
	"reflect"
)

type QueueMovementSystem struct {
	targets map[uint64]archetypes.QueuedMoveTarget
	*singletons.KeyboardState
}

func NewQueueMovementSystem() *QueueMovementSystem {
	return &QueueMovementSystem{
		targets:       make(map[uint64]archetypes.QueuedMoveTarget),
		KeyboardState: singletons.GetKeyboardState(),
	}
}

func (s *QueueMovementSystem) Update(deltaT float32) {
	for _, target := range s.targets {
		if target.MovementQueueLockoutCountdown > 0 {
			target.MovementQueueLockoutCountdown -= deltaT
			target.EnqueuedMove = utils.FacingNone
			continue
		}
		if s.Up && !s.Down {
			target.EnqueuedMove = utils.FacingUp
			target.MovementQueueLockoutCountdown = target.QueueCooldown
		}
		if s.Down && !s.Up {
			target.EnqueuedMove = utils.FacingDown
			target.MovementQueueLockoutCountdown = target.QueueCooldown
		}
		if s.Left && !s.Right {
			target.EnqueuedMove = utils.FacingLeft
			target.MovementQueueLockoutCountdown = target.QueueCooldown
		}
		if s.Right && !s.Left {
			target.EnqueuedMove = utils.FacingRight
			target.MovementQueueLockoutCountdown = target.QueueCooldown
		}
	}

}

func (s *QueueMovementSystem) RemoveEntity(entity engine.Identifier) {
	delete(s.targets, entity.GetId())
}

func (s *QueueMovementSystem) AddEntity(entity engine.Identifier) {
	s.targets[entity.GetId()] = entity.(archetypes.QueuedMoveTargetable).GetQueuedMoveTarget()
}

func (s *QueueMovementSystem) GetTargetType() reflect.Type {
	return reflect.TypeFor[archetypes.QueuedMoveTargetable]()
}
