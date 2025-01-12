package entities

import (
	"boxes/archetypes"
	"boxes/components"
	"boxes/engine"
)

type StaticSprite struct {
	engine.BasicEntity
	components.LocationComponent
	components.SpriteComponent
	components.SizeComponent
}

func NewStaticSprite(x, y, width, height float32, spriteName string, layer uint8) *StaticSprite {
	return &StaticSprite{
		BasicEntity: engine.NewBasicEntity(),
		LocationComponent: components.LocationComponent{
			X: x,
			Y: y,
		},
		SpriteComponent: components.SpriteComponent{
			SpriteName: spriteName,
			Layer:      layer,
		},
		SizeComponent: components.SizeComponent{
			Width:  width,
			Height: height,
		},
	}
}

func (s *StaticSprite) GetSpriteDrawTarget() archetypes.SpriteDrawTarget {
	return archetypes.SpriteDrawTarget{
		LocationComponent: &s.LocationComponent,
		SpriteComponent:   &s.SpriteComponent,
		SizeComponent:     &s.SizeComponent,
	}
}
