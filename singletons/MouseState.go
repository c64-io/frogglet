package singletons

import (
	"sync"
)

type MouseState struct {
	MouseMovementX float32
	MouseMovementY float32
}

var mouseStateSingleton *MouseState
var mouseStateOnce sync.Once

func GetMouseState() *MouseState {
	if mouseStateSingleton == nil {
		mouseStateOnce.Do(func() {
			mouseStateSingleton = &MouseState{}
		})
	}
	return mouseStateSingleton
}
