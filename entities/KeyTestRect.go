package entities

import (
	"boxes/archetypes"
	"boxes/components"
	"boxes/engine"
)

type KeyTestRect struct {
	engine.BasicEntity
	components.LocationComponent
	components.SizeComponent
	components.TargetKeyComponent
	components.ColorComponent
}

func (k KeyTestRect) GetDrawKeyRectTarget() archetypes.KeyToColor {
	return archetypes.KeyToColor{
		ColorComponent:     &k.ColorComponent,
		TargetKeyComponent: &k.TargetKeyComponent,
	}
}

func (k KeyTestRect) GetKeyTestTarget() archetypes.KeyTest {
	return archetypes.KeyTest{
		TargetKeyComponent: &k.TargetKeyComponent,
	}
}
