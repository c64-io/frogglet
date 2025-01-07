package resources

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type FontRenderer struct {
	font     *ttf.Font
	renderer *sdl.Renderer
}

func NewFontRenderer(fontPath string, fontSize int, renderer *sdl.Renderer) (*FontRenderer, error) {

	font, err := ttf.OpenFont(fontPath, fontSize)
	if err != nil {
		return nil, err
	}

	return &FontRenderer{
		font:     font,
		renderer: renderer,
	}, nil
}

func (f *FontRenderer) RenderText(text string, x, y int32, color sdl.Color) {
	surface, err := f.font.RenderUTF8Blended(text, color)
	if err != nil {
		return
	}
	defer surface.Free()

	texture, err := f.renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return
	}

	texture.SetBlendMode(sdl.BLENDMODE_BLEND)
	defer texture.Destroy()

	_, _, w, h, err := texture.Query()
	if err != nil {
		return
	}

	dst := sdl.Rect{X: x, Y: y, W: w, H: h}
	f.renderer.Copy(texture, nil, &dst)
}
