package scene;

import "time";
//import "fmt";
import "image/color";

import "github.com/hajimehoshi/ebiten";
import "github.com/hajimehoshi/ebiten/ebitenutil";

// text example scene
type TimedTextScene struct {
	Scene;
	Game GameI;
	Text string;
	Until time.Time;
	Wait time.Duration;
}

func MakeTimedTextScene(game GameI, text string, wait time.Duration) (*TimedTextScene, int) {
	s := TimedTextScene{};
	s.Game = game;
	s.Text = text;
	s.Wait = wait;
	id := s.Game.AddScene(&s);
	return &s, id;
}

func (s *TimedTextScene) CheckTime() {
	if s.Until.IsZero() {
		s.Until = time.Now().Add(s.Wait);
	} else if time.Now().After(s.Until) {
		s.Game.RelScene(+1);
		s.Until = time.Time{};
	}
}

func (s *TimedTextScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black);
	left := (s.Until.Sub(time.Now()) + (time.Second / 2)).Round(time.Second);
	ebitenutil.DebugPrint(screen, s.Text + "\n(" + left.String() + ")");
}

func (s *TimedTextScene) Update(screen *ebiten.Image) {
	s.CheckTime();
	s.Draw(screen);
}
