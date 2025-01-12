package systems

import (
	"boxes/archetypes"
	"boxes/engine"
	"boxes/singletons"
	"boxes/utils"
	"reflect"
)

type BasicSpriteSelectionSystem struct {
	targets map[uint64]archetypes.BasicSpriteSelectionTarget
	*singletons.KeyboardState
	moveQueueLockout float32
}

func NewBasicSpriteSelectionSystem() *BasicSpriteSelectionSystem {
	return &BasicSpriteSelectionSystem{
		targets:       make(map[uint64]archetypes.BasicSpriteSelectionTarget),
		KeyboardState: singletons.GetKeyboardState(),
	}
}

func (s *BasicSpriteSelectionSystem) Update(deltaT float32) {
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

func (s *BasicSpriteSelectionSystem) RemoveEntity(entityId uint64) {
	delete(s.targets, entityId)
}

func (s *BasicSpriteSelectionSystem) AddEntity(entity engine.Identifier) {
	s.targets[entity.GetId()] = entity.(archetypes.BasicSpriteSelectionTargetable).GetBasicSpriteSelectionTarget()
}

func (s *BasicSpriteSelectionSystem) GetTargetType() reflect.Type {
	return reflect.TypeFor[archetypes.BasicSpriteSelectionTargetable]()
}
