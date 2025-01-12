package systems

import (
	"boxes/archetypes"
	"boxes/engine"
	"boxes/utils"
	"reflect"
)

type AabbCollisionSystem struct {
	targets map[uint64]archetypes.AabbCollisionTarget
}

func NewAabbCollisionSystem() *AabbCollisionSystem {
	return &AabbCollisionSystem{
		targets: make(map[uint64]archetypes.AabbCollisionTarget),
	}
}

func (s *AabbCollisionSystem) Update(deltaT float32) {

	for _, target := range s.targets {
		target.CollidesWith = utils.ColliderTypeNone
	}

	for id, target := range s.targets {
		for otherId, otherTarget := range s.targets {
			if id == otherId {
				continue
			}
			if target.X < otherTarget.X+otherTarget.Width &&
				target.X+target.Width > otherTarget.X &&
				target.Y < otherTarget.Y+otherTarget.Height &&
				target.Y+target.Height > otherTarget.Y {
				target.CollidesWith |= otherTarget.ColliderType
			}
		}

	}
}

func (s *AabbCollisionSystem) RemoveEntity(entityId uint64) {
	delete(s.targets, entityId)
}

func (s *AabbCollisionSystem) AddEntity(entity engine.Identifier) {
	s.targets[entity.GetId()] = entity.(archetypes.AabbCollisionTargetable).GetAabbCollisionTarget()
}

func (s *AabbCollisionSystem) GetTargetType() reflect.Type {
	return reflect.TypeFor[archetypes.AabbCollisionTargetable]()
}
