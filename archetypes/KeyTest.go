package archetypes

import "boxes/components"

type KeyTest struct {
	*components.TargetKeyComponent
}

type KeyTestTarget interface {
	GetKeyTestTarget() KeyTest
}
