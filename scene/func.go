package scene

import "image/color"

import "github.com/hajimehoshi/ebiten"
import "github.com/hajimehoshi/ebiten/ebitenutil"
import "github.com/hajimehoshi/ebiten/inpututil"

// function runner scene
type FuncScene struct {
	Scene
	Game GameI
	Func func(GameI)
}

func MakeFuncScene(game GameI, funct func(GameI)) (*FuncScene, int) {
	s := FuncScene{}
	s.Func = funct
	id := s.Game.AddScene(&s)
	return &s, id
}

func (s *FuncScene) Update(screen *ebiten.Image) {
	s.Func(s.Game)
	s.Game.RelScene(+1)
}
