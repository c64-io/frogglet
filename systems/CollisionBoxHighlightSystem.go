package systems

import (
	"boxes/archetypes"
	"boxes/engine"
	"boxes/utils"
	"github.com/veandco/go-sdl2/sdl"
	"reflect"
)

type CollisionBoxHighlightSystem struct {
	targets map[uint64]archetypes.CollisionBoxHighlightTarget
}

func NewCollisionBoxHighlightSystem() *CollisionBoxHighlightSystem {
	return &CollisionBoxHighlightSystem{
		targets: make(map[uint64]archetypes.CollisionBoxHighlightTarget),
	}
}

func (c *CollisionBoxHighlightSystem) Update(deltaT float32) {
	for _, target := range c.targets {
		if target.CollidesWith == utils.ColliderTypeNone {
			target.Color = sdl.Color{R: 255, G: 255, B: 255, A: 255}
		} else if target.CollidesWith&utils.ColliderTypeEnemy != 0 {
			target.Color = sdl.Color{R: 255, G: 0, B: 0, A: 255}
		} else if target.CollidesWith&utils.ColliderTypePlayer != 0 {
			target.Color = sdl.Color{R: 0, G: 255, B: 0, A: 255}
		} else if target.CollidesWith&utils.ColliderTypeProjectile != 0 {
			target.Color = sdl.Color{R: 255, G: 255, B: 0, A: 255}
		} else if target.CollidesWith&utils.ColliderTypeWall != 0 {
			target.Color = sdl.Color{R: 0, G: 0, B: 255, A: 255}
		} else if target.CollidesWith&utils.ColliderTypePickup != 0 {
			target.Color = sdl.Color{R: 255, G: 0, B: 255, A: 255}
		}
	}
}

func (c *CollisionBoxHighlightSystem) GetTargetType() reflect.Type {
	return reflect.TypeFor[archetypes.CollisionBoxHighlightTargetable]()
}

func (c *CollisionBoxHighlightSystem) RemoveEntity(entityId uint64) {
	delete(c.targets, entityId)
}

func (c *CollisionBoxHighlightSystem) AddEntity(entity engine.Identifier) {
	targetable := entity.(archetypes.CollisionBoxHighlightTargetable)
	c.targets[entity.GetId()] = targetable.GetCollisionBoxHighlightTarget()
}
