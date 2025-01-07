package entities

import (
	"boxes/archetypes"
	"boxes/components"
	"boxes/engine"
)

type Fly struct {
	engine.BasicEntity
	components.LocationComponent
	components.SizeComponent
	components.SpriteComponent
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
	}
}

func (t *Fly) GetSpriteDrawer() archetypes.SpriteDrawer {
	return archetypes.SpriteDrawer{
		LocationComponent: &t.LocationComponent,
		SpriteComponent:   &t.SpriteComponent,
		SizeComponent:     &t.SizeComponent,
	}
}

func (t *Fly) GetFlyControl() archetypes.FlyControl {
	return archetypes.FlyControl{
		SpriteComponent:   &t.SpriteComponent,
		LocationComponent: &t.LocationComponent,
	}
}
