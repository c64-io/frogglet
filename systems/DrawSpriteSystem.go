package systems

import (
	"boxes/archetypes"
	"boxes/engine"
	"boxes/resources"
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
}

func NewDrawSpriteSystem(r *sdl.Renderer) *DrawSpriteSystem {
	return &DrawSpriteSystem{
		layers:      make([]*SpriteLayer, 0),
		renderer:    r,
		spriteSheet: resources.GetSpriteSheet("assets/sprites.yaml", r),
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
	for _, target := range k.layers {
		for _, s := range target.Sprites {
			k.spriteSheet.DrawSprite(s.SpriteName, int32(s.X), int32(s.Y), int32(s.Width), int32(s.Height), k.renderer)
		}
	}
}

func (k *DrawSpriteSystem) RemoveEntity(entityId uint64) {
	for _, layer := range k.layers {
		delete(layer.Sprites, entityId)
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
