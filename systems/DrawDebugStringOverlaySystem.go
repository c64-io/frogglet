package systems

import (
	"boxes/archetypes"
	"boxes/engine"
	"boxes/resources"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"reflect"
	"strings"
)

const DEBUG_STRING_FONT_SIZE = 14

type DrawDebugStringOverlaySystem struct {
	targets      map[uint64]archetypes.DebugStringOverlay
	fontRenderer *resources.FontRenderer
}

func NewDrawDebugStringOverlaySystem(r *sdl.Renderer) *DrawDebugStringOverlaySystem {
	return &DrawDebugStringOverlaySystem{
		targets:      make(map[uint64]archetypes.DebugStringOverlay),
		fontRenderer: resources.GetFontRenderer("assets/Good Old DOS.ttf", DEBUG_STRING_FONT_SIZE, r),
	}
}

func (k *DrawDebugStringOverlaySystem) Update(deltaT float32) {

	y := int32(30)

	for _, target := range k.targets {
		for _, line := range target.Text {
			s := line

			if target.LocationComponent != nil {
				s = strings.ReplaceAll(s, "{x}", fmt.Sprintf("%v", target.X))
				s = strings.ReplaceAll(s, "{y}", fmt.Sprintf("%v", target.Y))
			}
			if target.SizeComponent != nil {
				s = strings.ReplaceAll(s, "{w}", fmt.Sprintf("%v", target.SizeComponent.Width))
				s = strings.ReplaceAll(s, "{h}", fmt.Sprintf("%v", target.SizeComponent.Height))
			}
			if target.SpriteComponent != nil {
				s = strings.ReplaceAll(s, "{sprite}", target.SpriteComponent.SpriteName)
			}

			k.fontRenderer.RenderText(s, 5, y, target.Color)
			y += DEBUG_STRING_FONT_SIZE + 2
		}
	}
}

func (k *DrawDebugStringOverlaySystem) RemoveEntity(entity engine.Identifier) {
	delete(k.targets, entity.GetId())

}

func (k *DrawDebugStringOverlaySystem) AddEntity(entity engine.Identifier) {
	k.targets[entity.GetId()] = entity.(archetypes.DebugStringOverlayable).GetDebugStringOverlay()
}

func (k *DrawDebugStringOverlaySystem) GetTargetTypes() []reflect.Type {
	return []reflect.Type{
		reflect.TypeFor[archetypes.DebugStringOverlayable](),
	}
}
