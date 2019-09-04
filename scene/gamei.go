package scene

import "github.com/hajimehoshi/ebiten"

type GameI interface {
	Width() int
	Height() int
	Scene() Scene
	GoScene(int)
	RelScene(int)
	LastScene() int
	BackScene()
	AddScene(Scene) int
	Update(*ebiten.Image)
	Run()
}
