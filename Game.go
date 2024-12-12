package main

import (
	"boxes/engine"
	"boxes/systems"
	"boxes/utils"
	"github.com/veandco/go-sdl2/sdl"
	"runtime"
)

type Game struct {
	Engine  engine.Scene
	Window  *sdl.Window
	Surface *sdl.Surface
	Timer   *utils.Timer
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
	//g.Engine.AddSystem(systems.NewKeyboardSystem(g.Window))
	//g.Engine.AddSystem(systems.NewMouseSystem(g.Window))
}

func (g *Game) loadEntities() {

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

		g.Engine.Update(g.Timer.Elapsed)

		g.Window.UpdateSurface()

	}

}

func (g *Game) setupSdl() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
}

func (g *Game) setupWindow() {
	var err error
	g.Window, err = sdl.CreateWindow("Frogglet", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)

	if err != nil {
		panic(err)
	}

	g.Surface, err = g.Window.GetSurface()

	if err != nil {
		panic(err)
	}
}

func (g *Game) shutdown() {
	g.Window.Destroy()
	sdl.Quit()
}
