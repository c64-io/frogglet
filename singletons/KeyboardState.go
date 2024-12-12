package singletons

import (
	"sync"
)

type KeyboardState struct {
	Up       bool
	Down     bool
	Left     bool
	Right    bool
	Quit     bool
	ShowWire bool
	TestKey  bool
}

var keyboardStateSingleton *KeyboardState
var keyboardStateOnce sync.Once

func GetKeyboardState() *KeyboardState {
	if keyboardStateSingleton == nil {
		keyboardStateOnce.Do(func() {
			keyboardStateSingleton = &KeyboardState{}
		})
	}
	return keyboardStateSingleton
}
