package engine

type EntityQueue struct {
	HasAdditions bool
	HasRemovals  bool
	Additions    []Identifier
	Removals     []uint64
}

func (e *EntityQueue) EnqueueAdd(entity Identifier) {
	e.Additions = append(e.Additions, entity)
	e.HasAdditions = true
}

func (e *EntityQueue) EnqueueRemove(entityId uint64) {
	e.Removals = append(e.Removals, entityId)
	e.HasRemovals = true
}

func NewEntityQueue() *EntityQueue {
	return &EntityQueue{
		Additions: make([]Identifier, 0),
		Removals:  make([]uint64, 0),
	}
}
