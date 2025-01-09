package systems

import (
	"boxes/archetypes"
	"boxes/engine"
	"boxes/resources"
	"boxes/singletons"
	"cmp"
	"github.com/veandco/go-sdl2/sdl"
	"reflect"
	"slices"
)

type SpriteLayer struct {
	LayerNumber uint8
	Sprites     map[uint64]archetypes.SpriteDrawTarget
}

type DrawSpriteSystem struct {
	layers      []*SpriteLayer
	renderer    *sdl.Renderer
	spriteSheet *resources.SpriteSheet
	keys        *singletons.KeyboardState
	drawBorder  bool
}

func NewDrawSpriteSystem(r *sdl.Renderer) *DrawSpriteSystem {
	return &DrawSpriteSystem{
		layers:      make([]*SpriteLayer, 0),
		renderer:    r,
		spriteSheet: resources.GetSpriteSheet("assets/sprites.yaml", r),
		keys:        singletons.GetKeyboardState(),
		drawBorder:  false,
	}
}

func (k *DrawSpriteSystem) CreateOrGetLayer(layer uint8) *SpriteLayer {
	newLayer := &SpriteLayer{
		LayerNumber: layer,
		Sprites:     make(map[uint64]archetypes.SpriteDrawTarget),
	}

	if len(k.layers) == 0 {
		k.layers = append(k.layers, newLayer)
		return newLayer
	}

	index, found := slices.BinarySearchFunc(k.layers, newLayer, func(slA, slB *SpriteLayer) int {
		return cmp.Compare(slA.LayerNumber, slB.LayerNumber)
	})
	if !found {
		k.layers = slices.Insert(k.layers, index, newLayer)
	}
	return k.layers[index]
}

func (k *DrawSpriteSystem) Update(deltaT float32) {

	if k.keys.TestKeyReleased {
		k.drawBorder = !k.drawBorder
	}
	for _, target := range k.layers {
		for _, s := range target.Sprites {
			k.spriteSheet.DrawSprite(s.SpriteName, int32(s.X), int32(s.Y), s.Width, s.Height, k.renderer)
		}
	}

	if k.drawBorder {
		for _, target := range k.layers {
			for _, s := range target.Sprites {
				k.renderer.SetDrawColor(255, 0, 0, 255)
				k.renderer.DrawRect(&sdl.Rect{X: int32(s.X), Y: int32(s.Y), W: s.Width, H: s.Height})
			}
		}
	}

	// Do nothing
}

func (k *DrawSpriteSystem) RemoveEntity(entity engine.Identifier) {
	for _, layer := range k.layers {
		delete(layer.Sprites, entity.GetId())
	}

}

func (k *DrawSpriteSystem) AddEntity(entity engine.Identifier) {
	spriteDrawer := entity.(archetypes.SpriteDrawTargetable).GetSpriteDrawTarget()

	layer := k.CreateOrGetLayer(spriteDrawer.Layer)
	layer.Sprites[entity.GetId()] = spriteDrawer

}

func (k *DrawSpriteSystem) GetTargetType() reflect.Type {
	return reflect.TypeFor[archetypes.SpriteDrawTargetable]()
}
