package systems

import (
	"boxes/archetypes"
	"boxes/engine"
	"boxes/singletons"
	"boxes/utils"
	"reflect"
)

type BasicMovementSpriteSelectionSystem struct {
	targets map[uint64]archetypes.BasicMovementSpriteSelector
	*singletons.KeyboardState
	moveQueueLockout float32
}

func NewBasicMovementSpriteSelectionSystem() *BasicMovementSpriteSelectionSystem {
	return &BasicMovementSpriteSelectionSystem{
		targets:       make(map[uint64]archetypes.BasicMovementSpriteSelector),
		KeyboardState: singletons.GetKeyboardState(),
	}
}

func (s *BasicMovementSpriteSelectionSystem) Update(deltaT float32) {
	for _, target := range s.targets {

		if target.IsMoving {
			switch target.Heading {
			case utils.FacingUp:
				target.SpriteName = "FrogUpJump"
				break
			case utils.FacingDown:
				target.SpriteName = "FrogDownJump"
				break
			case utils.FacingLeft:
				target.SpriteName = "FrogLeftJump"
				break
			case utils.FacingRight:
				target.SpriteName = "FrogRightJump"
				break
			}
		} else {
			switch target.Heading {
			case utils.FacingUp:
				target.SpriteName = "FrogUp"
				break
			case utils.FacingDown:
				target.SpriteName = "FrogDown"
				break
			case utils.FacingLeft:
				target.SpriteName = "FrogLeft"
				break
			case utils.FacingRight:
				target.SpriteName = "FrogRight"
				break
			}
		}
	}
}

func (s *BasicMovementSpriteSelectionSystem) RemoveEntity(entity engine.Identifier) {
	delete(s.targets, entity.GetId())
}

func (s *BasicMovementSpriteSelectionSystem) AddEntity(entity engine.Identifier) {
	s.targets[entity.GetId()] = entity.(archetypes.BasicMovementSpriteSelectable).GetBasicMovementSpriteSelector()
}

func (s *BasicMovementSpriteSelectionSystem) GetTargetTypes() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf((*archetypes.BasicMovementSpriteSelectable)(nil)).Elem(),
	}
}
