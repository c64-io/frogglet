package engine

import (
	"reflect"
)

type ISystem interface {
	GetTargetTypes() []reflect.Type
	Update(deltaT float32)
	RemoveEntity(entity Identifier)
	AddEntity(entity Identifier)
}
