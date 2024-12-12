package engine

import (
	"sync/atomic"
)

var eidCounter = atomic.Uint64{}

type BasicEntity struct {
	Id       uint64
	Parent   *BasicEntity
	Children []*BasicEntity
}

func NewBasicEntity() BasicEntity {

	newId := eidCounter.Add(1)
	return BasicEntity{
		Id:       newId,
		Parent:   nil,
		Children: make([]*BasicEntity, 0),
	}
}

func (b BasicEntity) GetId() uint64 {
	return b.Id
}
