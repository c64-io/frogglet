package entities

import (
	"boxes/archetypes"
	"boxes/components"
	"boxes/engine"
	"boxes/utils"
)

type Fly struct {
	engine.BasicEntity
	components.LocationComponent
	components.SizeComponent
	components.SpriteComponent
	components.ColliderComponent
	components.ColorComponent
}

func NewFly(x, y float32) *Fly {
	return &Fly{
		BasicEntity: engine.NewBasicEntity(),
		LocationComponent: components.LocationComponent{
			X: x,
			Y: y,
		},
		SpriteComponent: components.SpriteComponent{
			SpriteName: "Fly",
			Layer:      98,
		},
		SizeComponent: components.SizeComponent{
			Width:  15,
			Height: 15,
		},
		ColliderComponent: components.ColliderComponent{
			ColliderType: utils.ColliderTypeEnemy,
		},
	}
}

func (t *Fly) GetSpriteDrawTarget() archetypes.SpriteDrawTarget {
	return archetypes.SpriteDrawTarget{
		LocationComponent: &t.LocationComponent,
		SpriteComponent:   &t.SpriteComponent,
		SizeComponent:     &t.SizeComponent,
	}
}

func (t *Fly) GetFlySpawnTarget() archetypes.FlySpawnTarget {
	return archetypes.FlySpawnTarget{
		SpriteComponent:   &t.SpriteComponent,
		LocationComponent: &t.LocationComponent,
	}
}

func (t *Fly) GetAabbCollisionTarget() archetypes.AabbCollisionTarget {
	return archetypes.AabbCollisionTarget{
		LocationComponent: &t.LocationComponent,
		SizeComponent:     &t.SizeComponent,
		ColliderComponent: &t.ColliderComponent,
	}
}

func (t *Fly) GetCollisionBoxHighlightTarget() archetypes.CollisionBoxHighlightTarget {
	return archetypes.CollisionBoxHighlightTarget{
		ColliderComponent: &t.ColliderComponent,
		ColorComponent:    &t.ColorComponent,
	}
}

func (t *Fly) GetDrawRectTarget() archetypes.DrawRectTarget {
	return archetypes.DrawRectTarget{
		LocationComponent: &t.LocationComponent,
		SizeComponent:     &t.SizeComponent,
		ColorComponent:    &t.ColorComponent,
	}
}

func (t *Fly) GetFlyEatenTarget() archetypes.FlyEatenTarget {
	return archetypes.FlyEatenTarget{
		ColliderComponent: &t.ColliderComponent,
	}
}
