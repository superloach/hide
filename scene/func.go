package scene

import "github.com/hajimehoshi/ebiten"

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
