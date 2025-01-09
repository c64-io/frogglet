package systems

import (
	"boxes/archetypes"
	"boxes/engine"
	"boxes/entities"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
	"reflect"
)

type FlyControlSystem struct {
	targets  map[uint64]archetypes.FlyControlTarget
	renderer *sdl.Renderer
	queue    *engine.EntityQueue
}

func NewFlyControlSystem(r *sdl.Renderer) *FlyControlSystem {
	return &FlyControlSystem{
		targets:  make(map[uint64]archetypes.FlyControlTarget),
		renderer: r,
	}
}

func (k *FlyControlSystem) Update(deltaT float32) {
	if len(k.targets) == 0 {
		k.queue.EnqueueAdd(entities.NewFly(float32(rand.Intn(800)), float32(rand.Int31n(600))))
	}

	// Do nothing
}

func (k *FlyControlSystem) RemoveEntity(entity engine.Identifier) {
	delete(k.targets, entity.GetId())

}

func (k *FlyControlSystem) AddEntity(entity engine.Identifier) {
	k.targets[entity.GetId()] = entity.(archetypes.FlyControlTargetable).GetFlyControlTarget()
}

func (k *FlyControlSystem) GetTargetType() reflect.Type {
	return reflect.TypeFor[archetypes.FlyControlTargetable]()
}

func (k *FlyControlSystem) SetEntityQueue(queue *engine.EntityQueue) {
	k.queue = queue
}
