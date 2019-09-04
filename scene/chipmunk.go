package scene

import "fmt"
import "image"
import "bytes"
import "image/color"

import "hide/hiderr"

import "github.com/hajimehoshi/ebiten"
import "github.com/hajimehoshi/ebiten/ebitenutil"
//import "github.com/hajimehoshi/ebiten/inpututil"
import "github.com/jakecoffman/cp"

type ChipEnt struct {
	Space  *cp.Space
	Sprite *ebiten.Image
	Body   *cp.Body
	Shape  *cp.Shape
}

// chipmunk physics scene
type ChipmunkScene struct {
	Scene
	Game    GameI
	Space   *cp.Space
	Gravity cp.Vector
	Damping float64
	Ents    []*ChipEnt
}

func MakeChipmunkScene(game GameI) (*ChipmunkScene, int) {
	s := ChipmunkScene{}
	s.Game = game
	s.Gravity = cp.Vector{0, 100}
	id := s.Game.AddScene(&s)
	return &s, id
}

func (s *ChipmunkScene) MakeRect(image *ebiten.Image, x, y float64, mass float64) {
	ce := ChipEnt{}
	ce.Sprite = image

	ow, oh := ce.Sprite.Size()
	w := float64(ow)
	h := float64(oh)

	moment := cp.MomentForBox(mass, w, h)
	body := cp.NewBody(mass, moment)
	body.SetPosition(cp.Vector{x - w / 2.0, y - h / 2.0})
	ce.Body = body

	shape := cp.NewBox(body, w, h, 0)
	ce.Shape = shape

	s.Ents = append(s.Ents, &ce)
}

func (s *ChipmunkScene) Step() {
	if s.Space == nil {
		s.Space = cp.NewSpace()
		s.Space.Iterations = 1
		s.Space.SetGravity(s.Gravity)
	}

	s.Space.Step(1.0 / float64(ebiten.MaxTPS()))
}

func (s *ChipmunkScene) Keys() {}

func (s *ChipmunkScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	op := &ebiten.DrawImageOptions{}

	for _, ent := range s.Ents {
		if ent.Space == nil {
			ent.Space = s.Space
			ent.Space.AddBody(ent.Body)
			ent.Space.AddShape(ent.Shape)
		}
		op.GeoM.Reset()
		op.GeoM.Translate(ent.Body.Position().X, ent.Body.Position().Y)
		screen.DrawImage(ent.Sprite, op)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (s *ChipmunkScene) Update(screen *ebiten.Image) {
	s.Step()
	s.Keys()
	if !ebiten.IsDrawingSkipped() {
		s.Draw(screen)
	}
}
