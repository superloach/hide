package scene;

import "image/color";

import "github.com/hajimehoshi/ebiten";
import "github.com/hajimehoshi/ebiten/ebitenutil";
import "github.com/hajimehoshi/ebiten/inpututil";

// text example scene
type TextScene struct {
	Scene;
	Game GameI;
	Text string;
	Next ebiten.Key;
	Back ebiten.Key;
}

func MakeTextScene(game GameI, text string) (*TextScene, int) {
	s := TextScene{};
	s.Game = game;
	s.Text = text;
	s.Next = ebiten.KeyRight;
	s.Back = ebiten.KeyLeft;
	id := s.Game.AddScene(&s);
	return &s, id;
}

func (s *TextScene) Keys() {
	if inpututil.IsKeyJustPressed(s.Next) {
		s.Game.RelScene(+1);
	} else if inpututil.IsKeyJustPressed(s.Back) {
		s.Game.BackScene();
	}
}

func (s *TextScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black);
	ebitenutil.DebugPrint(screen, s.Text);
}

func (s *TextScene) Update(screen *ebiten.Image) {
	s.Keys();
	s.Draw(screen);
}
