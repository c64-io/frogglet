package systems

import (
	"boxes/engine"
	"boxes/singletons"
	"github.com/go-gl/glfw/v3.3/glfw"
	"reflect"
)

type MouseSystem struct {
	mouseState *singletons.MouseState

	previousMouseX float64
	previousMouseY float64

	cumulativeMouseX float64
	cumulativeMouseY float64
}

func NewMouseSystem(window *glfw.Window) *MouseSystem {
	ms := &MouseSystem{
		mouseState: singletons.GetMouseState(),
	}
	window.SetCursorPosCallback(ms.HandleMouse)
	return ms
}

func (k *MouseSystem) Update(deltaT float32) {
	k.mouseState.MouseMovementX = float32(k.cumulativeMouseX)
	k.mouseState.MouseMovementY = float32(k.cumulativeMouseY)

	k.cumulativeMouseX = 0
	k.cumulativeMouseY = 0
}

func (k *MouseSystem) AddEntity(entity engine.Identifier) {
	panic("What the actual fuck are you doing?? You cant add entities to the mouse system!")
}

func (k *MouseSystem) RemoveEntity(entity engine.Identifier) {
	// Do nothing
}

func (k *MouseSystem) GetTargetTypes() []reflect.Type {
	return make([]reflect.Type, 0)
}

func (k *MouseSystem) HandleMouse(w *glfw.Window, xpos float64, ypos float64) {
	deltaX := xpos - k.previousMouseX
	deltaY := ypos - k.previousMouseY

	k.previousMouseX = xpos
	k.previousMouseY = ypos

	k.cumulativeMouseX += deltaX
	k.cumulativeMouseY += deltaY
}
