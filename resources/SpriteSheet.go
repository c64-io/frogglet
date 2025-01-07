package resources

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"gopkg.in/yaml.v3"
	"os"
)

type SpriteSheetConfig struct {
	FileName string `yaml:"fileName"`
	Sprites  []struct {
		Name   string `yaml:"name"`
		X      int    `yaml:"x"`
		Y      int    `yaml:"y"`
		Width  int    `yaml:"width"`
		Height int    `yaml:"height"`
	} `yaml:"sprites"`
}

func NewSpriteSheetConfig(filePath string) (SpriteSheetConfig, error) {
	var config SpriteSheetConfig
	buf, err := os.ReadFile(filePath)

	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(buf, &config)

	return config, err
}

type Sprite struct {
	Name   string
	X      int32
	Y      int32
	Width  int32
	Height int32
}

type SpriteSheet struct {
	sprites map[string]Sprite
	texture *sdl.Texture
}

func NewSpriteSheet(fileName string, renderer *sdl.Renderer) (SpriteSheet, error) {
	var sheet SpriteSheet

	config, err := NewSpriteSheetConfig(fileName)

	if err != nil {
		return sheet, fmt.Errorf("error reading sprite sheet config: %v", err)
	}

	sheet.sprites = make(map[string]Sprite)
	surface, err := img.Load(config.FileName)

	if err != nil {
		return sheet, fmt.Errorf("error loading sprite sheet image: %v", err)
	}

	//create texture from surface
	sheet.texture, err = renderer.CreateTextureFromSurface(surface)

	sheet.texture.SetBlendMode(sdl.BLENDMODE_BLEND)
	//sheet.texture.SetScaleMode(sdl.ScaleModeNearest)

	if err != nil {
		return sheet, fmt.Errorf("error creating texture from surface: %v", err)
	}

	for _, sprite := range config.Sprites {
		sheet.sprites[sprite.Name] = Sprite{
			Name:   sprite.Name,
			X:      int32(sprite.X),
			Y:      int32(sprite.Y),
			Width:  int32(sprite.Width),
			Height: int32(sprite.Height),
		}
	}

	return sheet, nil

}

func (s *SpriteSheet) GetSprite(name string) (Sprite, bool) {
	sprite, ok := s.sprites[name]
	return sprite, ok
}

func (s *SpriteSheet) DrawSprite(name string, x, y, w, h int32, renderer *sdl.Renderer) {
	sprite, ok := s.sprites[name]

	if !ok {
		return
	}

	renderer.Copy(
		s.texture,
		&sdl.Rect{X: sprite.X, Y: sprite.Y, W: sprite.Width, H: sprite.Height},
		&sdl.Rect{X: x, Y: y, W: w, H: h},
	)
}
