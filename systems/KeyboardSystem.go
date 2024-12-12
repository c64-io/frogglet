package systems

import (
	"boxes/archetypes"
	"boxes/engine"
	"github.com/veandco/go-sdl2/sdl"
	"reflect"
)

type Keyboard struct {
	targets map[sdl.Keycode]map[uint64]archetypes.KeyTest
}

func NewKeyboardSystem() *Keyboard {
	ks := &Keyboard{
		targets: make(map[sdl.Keycode]map[uint64]archetypes.KeyTest),
	}
	return ks
}

func (k *Keyboard) Update(deltaT float32) {

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.KeyboardEvent:
			if targets, ok := k.targets[t.Keysym.Sym]; ok {
				for _, target := range targets {
					target.IsPressed = t.Type == sdl.KEYDOWN
				}
			}
		}
	}
	// Do nothing
}

func (k *Keyboard) RemoveEntity(entity engine.Identifier) {

	target := entity.(archetypes.KeyTestTarget).GetKeyTestTarget()

	if keyTarget, ok := k.targets[target.TargetKey]; ok {
		delete(keyTarget, entity.GetId())
	}
}

func (k *Keyboard) AddEntity(entity engine.Identifier) {
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
