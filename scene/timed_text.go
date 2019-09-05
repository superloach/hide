package scene

import "time"

//import "fmt";
import "image/color"

import "github.com/hajimehoshi/ebiten"
import "github.com/hajimehoshi/ebiten/ebitenutil"

// text example scene
type TimedTextScene struct {
	Scene
	Game GameIface
	Text string
	Wait time.Duration
}

func MakeTimedTextScene(game GameIface, text string, wait time.Duration) (*TimedTextScene, int) {
	s := TimedTextScene{}
	s.Game = game
	s.Text = text
	s.Wait = wait
	id := s.Game.AddScene(&s)
	return &s, id
}

func (s *TimedTextScene) Step() {
	s.Wait -= time.Duration(float64(time.Second) / float64(ebiten.MaxTPS()))
	if s.Wait <= 0 {
		s.Game.RelScene(+1)
	}
}

func (s *TimedTextScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	left := (s.Wait + time.Second/2).Round(time.Second)
	ebitenutil.DebugPrint(screen, s.Text+"\n("+left.String()+")")
}

func (s *TimedTextScene) Update(screen *ebiten.Image) {
	s.Step()
	s.Draw(screen)
}
