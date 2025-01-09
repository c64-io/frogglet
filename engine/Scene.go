package engine

import (
	"fmt"
	"reflect"
)

type Scene struct {
	SystemMap   map[reflect.Type]EntityTargeter
	Systems     []Updater
	EntityQueue *EntityQueue
}

func NewScene() Scene {
	return Scene{
		SystemMap:   make(map[reflect.Type]EntityTargeter),
		Systems:     make([]Updater, 0),
		EntityQueue: NewEntityQueue(),
	}
}

func (e *Scene) AddSystem(system Updater) {
	e.Systems = append(e.Systems, system)

	if entityTargeter, ok := system.(EntityTargeter); ok {
		e.SystemMap[entityTargeter.GetTargetType()] = entityTargeter
	}
	if eqUser, ok := system.(EntityQueueUser); ok {
		eqUser.SetEntityQueue(e.EntityQueue)
	}
	if sysInit, ok := system.(Initializer); ok {
		sysInit.Init()
	}
}

func (e *Scene) AddEntity(entity Identifier) {
	for systemType, system := range e.SystemMap {
		if reflect.TypeOf(entity).Implements(systemType) {
			system.AddEntity(entity)
		}
	}
}

func (e *Scene) RemoveEntity(entity Identifier) {
	for _, system := range e.SystemMap {
		system.RemoveEntity(entity)
	}
}

func (e *Scene) Update(deltaT float32) {
	for _, system := range e.Systems {
		system.Update(deltaT)
	}

	if e.EntityQueue.HasAdditions {
		for _, entity := range e.EntityQueue.Additions {
			e.AddEntity(entity)
		}
		e.EntityQueue.HasAdditions = false
		e.EntityQueue.Additions = make([]Identifier, 0)
	}
	if e.EntityQueue.HasRemovals {
		for _, entity := range e.EntityQueue.Removals {
			fmt.Printf("Removing Entity %v\n", entity.GetId())
			e.RemoveEntity(entity)
		}
		e.EntityQueue.HasRemovals = false
		e.EntityQueue.Removals = make([]Identifier, 0)
	}
}
