package systems

import (
	"boxes/archetypes"
	"boxes/engine"
	"boxes/singletons"
	"boxes/utils"
	"reflect"
)

type StepMovementSystem struct {
	targets map[uint64]archetypes.SteppedMover
	*singletons.KeyboardState
}

func NewStepMovementSystem() *StepMovementSystem {
	return &StepMovementSystem{
		targets:       make(map[uint64]archetypes.SteppedMover),
		KeyboardState: singletons.GetKeyboardState(),
	}
}

func startStep(frog archetypes.SteppedMover) {
	frog.IsMoving = true
	frog.Heading = frog.EnqueuedMove
	frog.StepDirection = frog.EnqueuedMove
	frog.StepStartX = frog.X
	frog.StepStartY = frog.Y
	frog.EnqueuedMove = utils.FacingNone

}

func stopStep(frog archetypes.SteppedMover) {
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

func (s *StepMovementSystem) RemoveEntity(entity engine.Identifier) {
	delete(s.targets, entity.GetId())
}

func (s *StepMovementSystem) AddEntity(entity engine.Identifier) {
	s.targets[entity.GetId()] = entity.(archetypes.SteppedMoveable).GetSteppedMover()
}

func (s *StepMovementSystem) GetTargetTypes() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf((*archetypes.SteppedMoveable)(nil)).Elem(),
	}
}
