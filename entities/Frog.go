package entities

import (
	"boxes/archetypes"
	"boxes/components"
	"boxes/engine"
	"boxes/utils"
	"github.com/veandco/go-sdl2/sdl"
)

const FROG_JUMP_SPEED = 200.0
const FROG_JUMP_DISTANCE = 40.0
const MOVE_QUEUE_LOCKOUT = FROG_JUMP_DISTANCE / FROG_JUMP_SPEED * 2

type TestPlayer struct {
	engine.BasicEntity
	components.LocationComponent
	components.SpriteComponent
	components.SizeComponent
	components.HeadingComponent
	components.DebugStringComponent
	components.QueuedMovementComponent
	components.BasicMovementComponent
	components.SteppedMovementComponent
}

func NewTestPlayer(x, y float32) *TestPlayer {
	return &TestPlayer{
		BasicEntity: engine.NewBasicEntity(),
		LocationComponent: components.LocationComponent{
			X: x,
			Y: y,
		},
		SpriteComponent: components.SpriteComponent{
			SpriteName: "FrogUp",
			Layer:      99,
		},
		SizeComponent: components.SizeComponent{
			Width:  44,
			Height: 52,
		},
		HeadingComponent: components.HeadingComponent{
			Heading: utils.FacingDown,
		},
		DebugStringComponent: components.DebugStringComponent{
			Text: []string{
				"Player x,y: {x},{y}",
				"Player w,h: {w},{h}",
				"Player sprite: {sprite}",
			},
			Color: sdl.Color{R: 255, G: 0, B: 0, A: 150},
		},
		QueuedMovementComponent: components.QueuedMovementComponent{
			QueueCooldown: MOVE_QUEUE_LOCKOUT,
		},
		SteppedMovementComponent: components.SteppedMovementComponent{
			StepSpeed:    FROG_JUMP_SPEED,
			StepDistance: FROG_JUMP_DISTANCE,
		},
	}
}

func (t *TestPlayer) GetQueuedMover() archetypes.QueuedMover {
	return archetypes.QueuedMover{
		QueuedMovementComponent: &t.QueuedMovementComponent,
	}
}

func (t *TestPlayer) GetSteppedMover() archetypes.SteppedMover {
	return archetypes.SteppedMover{
		BasicMovementComponent:   &t.BasicMovementComponent,
		SteppedMovementComponent: &t.SteppedMovementComponent,
		QueuedMovementComponent:  &t.QueuedMovementComponent,
		HeadingComponent:         &t.HeadingComponent,
		LocationComponent:        &t.LocationComponent,
		SpriteComponent:          &t.SpriteComponent,
	}
}

func (t *TestPlayer) GetBasicMovementSpriteSelector() archetypes.BasicMovementSpriteSelector {
	return archetypes.BasicMovementSpriteSelector{
		BasicMovementComponent: &t.BasicMovementComponent,
		HeadingComponent:       &t.HeadingComponent,
		SpriteComponent:        &t.SpriteComponent,
	}
}

func (t *TestPlayer) GetSpriteDrawer() archetypes.SpriteDrawer {
	return archetypes.SpriteDrawer{
		LocationComponent: &t.LocationComponent,
		SpriteComponent:   &t.SpriteComponent,
		SizeComponent:     &t.SizeComponent,
	}
}

func (t *TestPlayer) GetDebugStringOverlay() archetypes.DebugStringOverlay {
	return archetypes.DebugStringOverlay{
		LocationComponent:    &t.LocationComponent,
		SizeComponent:        &t.SizeComponent,
		SpriteComponent:      &t.SpriteComponent,
		DebugStringComponent: &t.DebugStringComponent,
	}
}
