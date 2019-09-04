package game

import "errors"

import "hide/hiderr"
import "hide/scene"

import "github.com/hajimehoshi/ebiten"
import "github.com/hajimehoshi/ebiten/inpututil"

type Game struct {
	name       string
	width      int
	height     int
	scale      float64
	scenes     []scene.Scene
	curScene   int
	sceneStack []int
	QuitKey    ebiten.Key
	Fullscreen bool
}

func MakeGame(name string, width int, height int) *Game {
	g := Game{}
	g.name = name
	g.width = width
	g.height = height
	g.scale = ebiten.DeviceScaleFactor()
	g.scenes = make([]scene.Scene, 0)
	g.curScene = 0
	g.sceneStack = make([]int, 0)
	g.QuitKey = ebiten.KeyEscape
	g.Fullscreen = false
	return &g
}

func (g *Game) Width() int  { return g.width }
func (g *Game) Height() int { return g.height }

func (g *Game) Scene() scene.Scene {
	if len(g.scenes) == 0 {
		s, _ := scene.MakeTextScene(g, "no scenes :(")
		g.AddScene(s)
	}
	return g.scenes[g.curScene]
}

func (g *Game) GoScene(id int) {
	if id >= len(g.scenes) {
		hiderr.Do(errors.New("index out of bounds"))
	}
	g.curScene = id
	g.sceneStack = append(g.sceneStack, id)
}

func (g *Game) RelScene(rel int) {
	g.GoScene((g.curScene + rel) % len(g.scenes))
}

func (g *Game) LastScene() int {
	last := len(g.sceneStack) - 2
	if last < 0 {
		last = 0
	}
	return g.sceneStack[last]
}

func (g *Game) BackScene() {
	g.curScene = g.LastScene()
	g.sceneStack = g.sceneStack[:g.curScene+1]
}

func (g *Game) AddScene(scene scene.Scene) int {
	g.scenes = append(g.scenes, scene)
	return len(g.scenes) - 1
}

func (g *Game) Update(screen *ebiten.Image) {
	if inpututil.IsKeyJustPressed(g.QuitKey) {
		hiderr.Quit()
	}

	s := g.Scene()
	s.Update(screen)
}

func (g *Game) Run() {
	g.GoScene(g.curScene)

	if g.Fullscreen {
		ebiten.SetFullscreen(true)

		w, h := ebiten.ScreenSizeInFullscreen()

		if w > 0 && h > 0 {
			g.width = int(float64(w) * g.scale)
			g.height = int(float64(h) * g.scale)
		}
	}

	tmp := func(screen *ebiten.Image) error { g.Update(screen); return nil }
	err := ebiten.Run(tmp, g.width, g.height, 1/g.scale, g.name)
	hiderr.Do(err)
}
