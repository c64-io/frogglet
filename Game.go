package main

import (
	"boxes/engine"
	"boxes/entities"
	"boxes/systems"
	"boxes/utils"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"runtime"
)

type Game struct {
	Engine   engine.Scene
	Window   *sdl.Window
	Renderer *sdl.Renderer
	Timer    *utils.Timer
}

func NewGame() *Game {

	game := &Game{
		Engine: engine.NewScene(),
		Timer:  utils.NewTimer(),
	}

	return game
}

func (g *Game) loadSystems() {
	g.Engine.AddSystem(systems.NewKeyboardSystem())

	g.Engine.AddSystem(systems.NewQueueMovementSystem())
	g.Engine.AddSystem(systems.NewStepMovementSystem())
	g.Engine.AddSystem(systems.NewFlyEatenSystem())
	g.Engine.AddSystem(systems.NewFlySpawnSystem())

	g.Engine.AddSystem(systems.NewAabbCollisionSystem())
	g.Engine.AddSystem(systems.NewCollisionBoxHighlightSystem())

	g.Engine.AddSystem(systems.NewBasicSpriteSelectionSystem())
	g.Engine.AddSystem(systems.NewDrawSpriteSystem(g.Renderer))
	//g.Engine.AddSystem(systems.NewDrawRectSystem(g.Renderer))
	g.Engine.AddSystem(systems.NewDrawKeyboardOverlaySystem(g.Renderer))
	g.Engine.AddSystem(systems.NewDebugStringOverlaySystem(g.Renderer))

}

func (g *Game) loadEntities() {

	g.Engine.AddEntity(entities.NewTestPlayer(0, 0))

}

func (g *Game) Run() {
	defer g.shutdown()
	runtime.LockOSThread()
	g.setupSdl()
	g.setupWindow()
	g.loadSystems()
	g.loadEntities()

	g.Timer.Tick()
	for {
		g.Timer.Tick()
		g.Renderer.SetDrawColor(0, 0, 0, 0)
		g.Renderer.Clear()

		g.Engine.Update(g.Timer.Elapsed)

		g.Renderer.Present()

	}

}

func (g *Game) setupSdl() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	if err := ttf.Init(); err != nil {
		panic(err)
	}
}

func (g *Game) setupWindow() {
	var err error

	g.Window, err = sdl.CreateWindow("Frogglet", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)

	if err != nil {
		panic(err)
	}

	g.Renderer, err = sdl.CreateRenderer(g.Window, -1, sdl.RENDERER_ACCELERATED)
	g.Renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	if err != nil {
		panic(err)
	}

}

func (g *Game) shutdown() {
	g.Window.Destroy()
	sdl.Quit()
}
