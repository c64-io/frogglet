package systems

import (
	"boxes/archetypes"
	"boxes/engine"
	"boxes/singletons"
	"boxes/utils"
	"reflect"
)

type StepMovementSystem struct {
	targets map[uint64]archetypes.SteppedMoveTarget
	*singletons.KeyboardState
}

func NewStepMovementSystem() *StepMovementSystem {
	return &StepMovementSystem{
		targets:       make(map[uint64]archetypes.SteppedMoveTarget),
		KeyboardState: singletons.GetKeyboardState(),
	}
}

func startStep(frog archetypes.SteppedMoveTarget) {
	frog.IsMoving = true
	frog.Heading = frog.EnqueuedMove
	frog.StepDirection = frog.EnqueuedMove
	frog.StepStartX = frog.X
	frog.StepStartY = frog.Y
	frog.EnqueuedMove = utils.FacingNone

}

func stopStep(frog archetypes.SteppedMoveTarget) {
	frog.IsMoving = false
}

func (s *StepMovementSystem) Update(deltaT float32) {
	for _, target := range s.targets {
		if target.IsMoving {
			switch target.StepDirection {
			case utils.FacingUp:
				target.Y -= target.StepSpeed * deltaT
				if target.Y < target.StepStartY-target.StepDistance {
					target.Y = target.StepStartY - target.StepDistance
					stopStep(target)
				}
				break
			case utils.FacingDown:
				target.Y += target.StepSpeed * deltaT
				if target.Y > target.StepStartY+target.StepDistance {
					target.Y = target.StepStartY + target.StepDistance
					stopStep(target)
				}
				break
			case utils.FacingLeft:
				target.X -= target.StepSpeed * deltaT
				if target.X < target.StepStartX-target.StepDistance {
					target.X = target.StepStartX - target.StepDistance
					stopStep(target)
				}
				break
			case utils.FacingRight:
				target.X += target.StepSpeed * deltaT
				if target.X > target.StepStartX+target.StepDistance {
					target.X = target.StepStartX + target.StepDistance
					stopStep(target)
				}
			}

		} else {

			if target.EnqueuedMove != utils.FacingNone {
				startStep(target)
			}

		}

	}
}

func (s *StepMovementSystem) RemoveEntity(entityId uint64) {
	delete(s.targets, entityId)
}

func (s *StepMovementSystem) AddEntity(entity engine.Identifier) {
	s.targets[entity.GetId()] = entity.(archetypes.SteppedMoveTargetable).GetSteppedMoveTarget()
}

func (s *StepMovementSystem) GetTargetType() reflect.Type {
	return reflect.TypeFor[archetypes.SteppedMoveTargetable]()
}
