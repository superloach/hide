package scene

import "bytes"
import "image"
import "image/color"

import "hide/hiderr"

import "github.com/hajimehoshi/ebiten"
import "github.com/hajimehoshi/ebiten/ebitenutil"
import "github.com/hajimehoshi/ebiten/inpututil"

// image viewer scene
type ImageScene struct {
	Game    GameI
	Image   *ebiten.Image
	Next    ebiten.Key
	Back    ebiten.Key
	Caption string
}

func MakeImageScene(game GameI, imageBytes []byte) (*ImageScene, int) {
	s := ImageScene{}
	s.Game = game
	s.Next = ebiten.KeyRight
	s.Back = ebiten.KeyLeft

	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	hiderr.Do(err)
	s.Image, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	id := s.Game.AddScene(&s)

	return &s, id
}

func (s *ImageScene) Keys() {
	if inpututil.IsKeyJustPressed(s.Next) {
		s.Game.RelScene(+1)
	} else if inpututil.IsKeyJustPressed(s.Back) {
		s.Game.BackScene()
	}
}

func (s *ImageScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	opt := &ebiten.DrawImageOptions{}
	w, h := s.Image.Size()
	opt.GeoM.Translate(float64(s.Game.Width())/2, float64(s.Game.Height())/2)
	opt.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	screen.DrawImage(s.Image, opt)

	ebitenutil.DebugPrint(screen, s.Caption)
}

func (s *ImageScene) Update(screen *ebiten.Image) {
	s.Keys()
	s.Draw(screen)
}
