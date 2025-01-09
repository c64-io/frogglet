package systems

import (
	"boxes/singletons"
	"github.com/veandco/go-sdl2/sdl"
)

type Keyboard struct {
	*singletons.KeyboardState
}

func NewKeyboardSystem() *Keyboard {
	ks := &Keyboard{
		KeyboardState: singletons.GetKeyboardState(),
	}
	return ks
}

func (k *Keyboard) Update(deltaT float32) {

	k.PreviousUp = k.Up
	k.PreviousDown = k.Down
	k.PreviousLeft = k.Left
	k.PreviousRight = k.Right
	k.PreviousQuit = k.Quit
	k.PreviousTestKey = k.TestKey

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.KeyboardEvent:
			switch t.Keysym.Sym {
			case sdl.K_UP:
				k.Up = t.State == sdl.PRESSED
				break
			case sdl.K_DOWN:
				k.Down = t.State == sdl.PRESSED
				break
			case sdl.K_LEFT:
				k.Left = t.State == sdl.PRESSED
				break
			case sdl.K_RIGHT:
				k.Right = t.State == sdl.PRESSED
				break
			case sdl.K_q:
				k.Quit = t.State == sdl.PRESSED
				break
			case sdl.K_t:
				k.TestKey = t.State == sdl.PRESSED
				break

			}
		}
	}

	k.UpPressed = k.Up && !k.PreviousUp
	k.DownPressed = k.Down && !k.PreviousDown
	k.LeftPressed = k.Left && !k.PreviousLeft
	k.RightPressed = k.Right && !k.PreviousRight
	k.QuitPressed = k.Quit && !k.PreviousQuit
	k.TestKeyPressed = k.TestKey && !k.PreviousTestKey

	k.UpReleased = !k.Up && k.PreviousUp
	k.DownReleased = !k.Down && k.PreviousDown
	k.LeftReleased = !k.Left && k.PreviousLeft
	k.RightReleased = !k.Right && k.PreviousRight
	k.QuitReleased = !k.Quit && k.PreviousQuit
	k.TestKeyReleased = !k.TestKey && k.PreviousTestKey

}
