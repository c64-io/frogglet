package resources

import (
	"sync"
)

type AssetManager struct {
	Textures map[string]*Texture
}

var assetManagerLoadOnce sync.Once
var singleAssetManager *AssetManager

func getAssetManager() *AssetManager {
	if singleAssetManager == nil {
		assetManagerLoadOnce.Do(func() {
			singleAssetManager = &AssetManager{
				Textures: make(map[string]*Texture),
			}
		})
	}
	return singleAssetManager
}

func GetTexture(name string) *Texture {
	am := getAssetManager()

	if _, ok := am.Textures[name]; !ok {
		tex, err := NewTexture(name, 0)
		if err != nil {
			return nil
		}
		am.Textures[name] = tex
		return tex
	}
	return am.Textures[name]
}
