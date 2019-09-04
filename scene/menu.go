package scene

import "image/color"

import "github.com/hajimehoshi/ebiten"
import "github.com/hajimehoshi/ebiten/ebitenutil"
import "github.com/hajimehoshi/ebiten/inpututil"

type MenuOption struct {
	Name  string
	Scene int
}

// option menu scene
type MenuScene struct {
	Game    GameI
	Title   string
	Options []MenuOption
	Cursor  int
	Next    ebiten.Key
	Back    ebiten.Key
	Confirm ebiten.Key
	Exit    ebiten.Key
}

func MakeMenuScene(game GameI, title string) (*MenuScene, int) {
	s := MenuScene{}
	s.Game = game
	s.Title = title
	s.Options = make([]MenuOption, 0)
	s.Cursor = 0
	s.Next = ebiten.KeyDown
	s.Back = ebiten.KeyUp
	s.Confirm = ebiten.KeyRight
	s.Exit = ebiten.KeyLeft
	id := s.Game.AddScene(&s)
	return &s, id
}

func (s *MenuScene) AddOption(name string, scene int) {
	opt := MenuOption{name, scene}
	s.Options = append(s.Options, opt)
}

func (s *MenuScene) Keys() {
	if inpututil.IsKeyJustPressed(s.Next) {
		s.Cursor += 1
		s.Cursor %= len(s.Options)
	} else if inpututil.IsKeyJustPressed(s.Back) {
		s.Cursor -= 1
		s.Cursor %= len(s.Options)
	} else if inpututil.IsKeyJustPressed(s.Confirm) {
		opt := s.Options[s.Cursor]
		s.Game.GoScene(opt.Scene)
	} else if inpututil.IsKeyJustPressed(s.Exit) {
		s.Game.BackScene()
	}
}

func (s *MenuScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	ebitenutil.DebugPrint(screen, s.Title)
	shift := ""
	for idx, opt := range s.Options {
		pre := "   "
		if idx == s.Cursor {
			pre = " * "
		}
		shift += "\n"
		ebitenutil.DebugPrint(screen, shift+pre+opt.Name)
	}
	ebitenutil.DebugPrint(screen, shift+"\n\n-> to select")
	ebitenutil.DebugPrint(screen, shift+"\n\n\n<- to exit menu")
}

func (s *MenuScene) Update(screen *ebiten.Image) {
	s.Keys()
	s.Draw(screen)
}
