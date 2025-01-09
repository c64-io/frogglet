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

type DebugStringOverlaySystem struct {
	targets      map[uint64]archetypes.DebugStringOverlayTarget
	fontRenderer *resources.FontRenderer
}

func NewDebugStringOverlaySystem(r *sdl.Renderer) *DebugStringOverlaySystem {
	return &DebugStringOverlaySystem{
		targets:      make(map[uint64]archetypes.DebugStringOverlayTarget),
		fontRenderer: resources.GetFontRenderer("assets/Good Old DOS.ttf", DEBUG_STRING_FONT_SIZE, r),
	}
}

func (k *DebugStringOverlaySystem) Update(deltaT float32) {

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

func (k *DebugStringOverlaySystem) RemoveEntity(entity engine.Identifier) {
	delete(k.targets, entity.GetId())

}

func (k *DebugStringOverlaySystem) AddEntity(entity engine.Identifier) {
	k.targets[entity.GetId()] = entity.(archetypes.DebugStringOverlayTargetable).GetDebugStringOverlayTarget()
}

func (k *DebugStringOverlaySystem) GetTargetType() reflect.Type {
	return reflect.TypeFor[archetypes.DebugStringOverlayTargetable]()
}
