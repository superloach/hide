package scene

import "github.com/hajimehoshi/ebiten"

// function runner scene
type JumpScene struct {
	Scene
	Game GameI
	To int
}

func MakeJumpScene(game GameI, to int) (*JumpScene, int) {
	s := JumpScene{}
	s.To = to
	id := s.Game.AddScene(&s)
	return &s, id
}

func (s *JumpScene) Update(screen *ebiten.Image) {
	s.Game.GoScene(s.To)
}
