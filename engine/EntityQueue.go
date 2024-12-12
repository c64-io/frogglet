package engine

type EntityQueue struct {
	HasAdditions bool
	HasRemovals  bool
	Additions    []Identifier
	Removals     []Identifier
}

func (e *EntityQueue) EnqueueAdd(entity Identifier) {
	e.Additions = append(e.Additions, entity)
	e.HasAdditions = true
}

func (e *EntityQueue) EnqueueRemove(entity Identifier) {
	e.Removals = append(e.Removals, entity)
	e.HasRemovals = true
}

func NewEntityQueue() *EntityQueue {
	return &EntityQueue{
		Additions: make([]Identifier, 0),
		Removals:  make([]Identifier, 0),
	}
}
