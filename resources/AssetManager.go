package resources

import (
	"github.com/veandco/go-sdl2/sdl"
	"sync"
)

type AssetManager struct {
	spriteSheets  map[string]*SpriteSheet
	fontRenderers map[string]*FontRenderer
}

var assetManagerLoadOnce sync.Once
var singleAssetManager *AssetManager

func getAssetManager() *AssetManager {
	if singleAssetManager == nil {
		assetManagerLoadOnce.Do(func() {
			singleAssetManager = &AssetManager{
				spriteSheets:  make(map[string]*SpriteSheet),
				fontRenderers: make(map[string]*FontRenderer),
			}
		})
	}
	return singleAssetManager
}

func GetSpriteSheet(name string, renderer *sdl.Renderer) *SpriteSheet {
	am := getAssetManager()

	if _, ok := am.spriteSheets[name]; !ok {
		sheet, err := NewSpriteSheet(name, renderer)
		if err != nil {
			return nil
		}
		am.spriteSheets[name] = &sheet
		return &sheet
	}
	return am.spriteSheets[name]
}

func GetFontRenderer(name string, fontSize int, renderer *sdl.Renderer) *FontRenderer {
	am := getAssetManager()

	if _, ok := am.fontRenderers[name]; !ok {
		fr, err := NewFontRenderer(name, fontSize, renderer)
		if err != nil {
			return nil
		}
		am.fontRenderers[name] = fr
		return fr
	}
	return am.fontRenderers[name]
}
