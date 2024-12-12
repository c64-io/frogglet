package systems

import (
	"boxes/archetypes"
	"boxes/engine"
	"reflect"
)

type KeyToColorSystem struct {
	targets map[uint64]archetypes.KeyToColor
}

func NewKeyToColorSystem() *KeyToColorSystem {
	ks := &KeyToColorSystem{
		targets: make(map[uint64]archetypes.KeyToColor),
	}
	return ks
}

func (k *KeyToColorSystem) Update(deltaT float32) {

	// Do nothing
}

func (k *KeyToColorSystem) RemoveEntity(entity engine.Identifier) {
	target := entity.(archetypes.KeyTestTarget).GetKeyTestTarget()

}

func (k *KeyToColorSystem) AddEntity(entity engine.Identifier) {
	target := entity.(archetypes.KeyTestTarget).GetKeyTestTarget()

	if k.targets[target.TargetKey] == nil {
		k.targets[target.TargetKey] = make(map[uint64]archetypes.KeyTest)
	}

	k.targets[target.TargetKey][entity.GetId()] = target

}

func (k *Keyboard) GetTargetTypes() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf((*archetypes.KeyTestTarget)(nil)).Elem(),
	}
}
