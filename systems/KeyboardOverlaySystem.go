package systems

import (
	"boxes/singletons"
	"github.com/veandco/go-sdl2/sdl"
)

type DrawKeyboardOverlaySystem struct {
	*singletons.KeyboardState
	renderer *sdl.Renderer
}

func NewDrawKeyboardOverlaySystem(r *sdl.Renderer) *DrawKeyboardOverlaySystem {
	return &DrawKeyboardOverlaySystem{
		KeyboardState: singletons.GetKeyboardState(),
		renderer:      r,
	}
}

func drawKeyRect(r *sdl.Renderer, x, y int32, pressed, released, held bool) {
	if pressed {
		r.SetDrawColor(255, 15, 15, 150)
	} else if released {
		r.SetDrawColor(15, 255, 15, 150)
	} else if held {
		r.SetDrawColor(15, 15, 255, 150)
	} else {
		r.SetDrawColor(75, 75, 75, 150)
	}
	r.FillRect(&sdl.Rect{
		X: x,
		Y: y,
		W: 15,
		H: 15,
	})
}

func (k *DrawKeyboardOverlaySystem) Update(deltaT float32) {

	drawKeyRect(k.renderer, 0, 15, k.LeftPressed, k.LeftReleased, k.Left)
	drawKeyRect(k.renderer, 15, 15, k.DownPressed, k.DownReleased, k.Down)
	drawKeyRect(k.renderer, 30, 15, k.RightPressed, k.RightReleased, k.Right)
	drawKeyRect(k.renderer, 15, 0, k.UpPressed, k.UpReleased, k.Up)

	// Do nothing
}
