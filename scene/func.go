package scene

import "github.com/hajimehoshi/ebiten"

// function runner scene
type FuncScene struct {
	Scene
	Game GameIface
	Func func(Scene)
}

func MakeFuncScene(game GameIface, funct func(Scene)) (*FuncScene, int) {
	s := FuncScene{}
	s.Func = funct
	id := s.Game.AddScene(&s)
	return &s, id
}

func (s *FuncScene) Update(screen *ebiten.Image) {
	s.Func(s)
	s.Game.RelScene(+1)
}
