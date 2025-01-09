package engine

import "reflect"

type EntityTargeter interface {
	GetTargetType() reflect.Type
	RemoveEntity(entity Identifier)
	AddEntity(entity Identifier)
}
