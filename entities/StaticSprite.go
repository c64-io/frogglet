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

func NewStaticSprite(x, y float32, width, height int32, spriteName string, layer uint8) *StaticSprite {
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

func (s *StaticSprite) GetSpriteDrawer() archetypes.SpriteDrawer {
	return archetypes.SpriteDrawer{
		LocationComponent: &s.LocationComponent,
		SpriteComponent:   &s.SpriteComponent,
		SizeComponent:     &s.SizeComponent,
	}
}
