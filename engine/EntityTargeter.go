package engine

import "reflect"

type EntityTargeter interface {
	GetTargetType() reflect.Type
	RemoveEntity(entityId uint64)
	AddEntity(entity Identifier)
}
