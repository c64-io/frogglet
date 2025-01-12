package systems

import (
	"boxes/archetypes"
	"boxes/engine"
	"boxes/utils"
	"reflect"
)

type FlyEatenSystem struct {
	targets map[uint64]archetypes.FlyEatenTarget
	queue   *engine.EntityQueue
}

func NewFlyEatenSystem() *FlyEatenSystem {
	return &FlyEatenSystem{
		targets: make(map[uint64]archetypes.FlyEatenTarget),
	}
}

func (k *FlyEatenSystem) Update(deltaT float32) {
	for id, target := range k.targets {
		if target.CollidesWith&utils.ColliderTypePlayer != 0 {
			k.queue.EnqueueRemove(id)
		}
	}
}

func (k *FlyEatenSystem) RemoveEntity(entityId uint64) {
	delete(k.targets, entityId)

}

func (k *FlyEatenSystem) AddEntity(entity engine.Identifier) {
	k.targets[entity.GetId()] = entity.(archetypes.FlyEatenTargetable).GetFlyEatenTarget()
}

func (k *FlyEatenSystem) GetTargetType() reflect.Type {
	return reflect.TypeFor[archetypes.FlyEatenTargetable]()
}

func (k *FlyEatenSystem) SetEntityQueue(queue *engine.EntityQueue) {
	k.queue = queue
}
