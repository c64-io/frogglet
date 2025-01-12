package systems

import (
	"boxes/archetypes"
	"boxes/engine"
	"boxes/entities"
	"math/rand"
	"reflect"
)

type FlySpawnSystem struct {
	targets map[uint64]archetypes.FlySpawnTarget
	queue   *engine.EntityQueue
}

func NewFlySpawnSystem() *FlySpawnSystem {
	return &FlySpawnSystem{
		targets: make(map[uint64]archetypes.FlySpawnTarget),
	}
}

func (k *FlySpawnSystem) Update(deltaT float32) {
	if len(k.targets) == 0 {
		k.queue.EnqueueAdd(entities.NewFly(float32(rand.Intn(800)), float32(rand.Int31n(600))))
	}

	// Do nothing
}

func (k *FlySpawnSystem) RemoveEntity(entityId uint64) {
	delete(k.targets, entityId)

}

func (k *FlySpawnSystem) AddEntity(entity engine.Identifier) {
	k.targets[entity.GetId()] = entity.(archetypes.FlySpawnTargetable).GetFlySpawnTarget()
}

func (k *FlySpawnSystem) GetTargetType() reflect.Type {
	return reflect.TypeFor[archetypes.FlySpawnTargetable]()
}

func (k *FlySpawnSystem) SetEntityQueue(queue *engine.EntityQueue) {
	k.queue = queue
}
