package singletons

import (
	"sync"
)

type KeyboardState struct {
	Up      bool
	Down    bool
	Left    bool
	Right   bool
	Quit    bool
	TestKey bool

	UpPressed      bool
	DownPressed    bool
	LeftPressed    bool
	RightPressed   bool
	QuitPressed    bool
	TestKeyPressed bool

	UpReleased      bool
	DownReleased    bool
	LeftReleased    bool
	RightReleased   bool
	QuitReleased    bool
	TestKeyReleased bool

	PreviousUp      bool
	PreviousDown    bool
	PreviousLeft    bool
	PreviousRight   bool
	PreviousQuit    bool
	PreviousTestKey bool
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
