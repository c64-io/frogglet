package components

import "github.com/veandco/go-sdl2/sdl"

type TargetKeyComponent struct {
	TargetKey sdl.Keycode
	IsPressed bool
}
