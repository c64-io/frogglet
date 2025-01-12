package systems

import (
	"boxes/archetypes"
	"boxes/engine"
	"github.com/veandco/go-sdl2/sdl"
	"reflect"
)

type DrawRectSystem struct {
	targets  map[uint64]archetypes.DrawRectTarget
	renderer *sdl.Renderer
}

func NewDrawRectSystem(r *sdl.Renderer) *DrawRectSystem {
	return &DrawRectSystem{
		targets:  make(map[uint64]archetypes.DrawRectTarget),
		renderer: r,
	}
}

func (k *DrawRectSystem) Update(deltaT float32) {
	for _, target := range k.targets {
		k.renderer.SetDrawColor(target.Color.R, target.Color.G, target.Color.B, target.Color.A)
		k.renderer.DrawRect(&sdl.Rect{
			X: int32(target.X),
			Y: int32(target.Y),
			W: int32(target.Width),
			H: int32(target.Height),
		})
	}

	// Do nothing
}

func (k *DrawRectSystem) RemoveEntity(entityId uint64) {
	delete(k.targets, entityId)

}

func (k *DrawRectSystem) AddEntity(entity engine.Identifier) {
	k.targets[entity.GetId()] = entity.(archetypes.DrawRectTargetable).GetDrawRectTarget()
}

func (k *DrawRectSystem) GetTargetType() reflect.Type {
	return reflect.TypeFor[archetypes.DrawRectTargetable]()
}
