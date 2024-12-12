package resources

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"image"
	"image/draw"
	_ "image/png"
	"os"
)

type Texture struct {
	TextureID   uint32
	Width       int32
	Height      int32
	TextureUnit int32
}

func NewTexture(fileName string, unit int32) (*Texture, error) {
	t := &Texture{}
	t.TextureUnit = unit

	imgFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("texture %q not found on disk: %v", fileName, err)
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return nil, fmt.Errorf("unsupported stride")
	}

	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	t.Width = int32(rgba.Rect.Size().X)
	t.Height = int32(rgba.Rect.Size().Y)

	gl.GenTextures(1, &t.TextureID)
	gl.ActiveTexture(uint32(gl.TEXTURE0 + t.TextureUnit))
	gl.BindTexture(gl.TEXTURE_2D, t.TextureID)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return t, nil
}

func (t *Texture) Use() {
	gl.ActiveTexture(uint32(gl.TEXTURE0 + t.TextureUnit))
	gl.BindTexture(gl.TEXTURE_2D, t.TextureID)
}

func (t *Texture) Stop() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}
