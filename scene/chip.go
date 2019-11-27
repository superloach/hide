package scene

import "fmt"
import "image/color"

import "github.com/superloach/hide/hiderr"

import "github.com/hajimehoshi/ebiten"
import "github.com/hajimehoshi/ebiten/ebitenutil"
import "github.com/jakecoffman/cp"

type ChipEnt struct {
	Name    string
	Space   *cp.Space
	Sprite  *ebiten.Image
	Width   float64
	Height  float64
	Body    *cp.Body
	Shape   *cp.Shape
	Collide func(*ChipEnt)
}

// chipmunk physics scene
type ChipScene struct {
	Scene
	Game  GameIface
	Space *cp.Space
	Ents  []*ChipEnt
}

func MakeChipScene(game GameIface) (*ChipScene, int) {
	s := ChipScene{}
	s.Game = game
	s.Space = cp.NewSpace()
	s.Space.Iterations = 1
	s.Space.SetGravity(cp.Vector{0, 333})

	id := s.Game.AddScene(&s)
	return &s, id
}

func (s *ChipScene) MakeRect(image *ebiten.Image, x, y float64, mass float64) *ChipEnt {
	ce := ChipEnt{}
	ce.Sprite = image

	ow, oh := ce.Sprite.Size()
	ce.Width = float64(ow)
	ce.Height = float64(oh)

	moment := cp.MomentForBox(mass, ce.Width, ce.Height)
	body := cp.NewBody(mass, moment)
	body.SetPosition(cp.Vector{x, y})
	ce.Body = body

	shape := cp.NewBox(body, ce.Width, ce.Height, -1)
	shape.SetElasticity(0.5)
	ce.Shape = shape

	s.Ents = append(s.Ents, &ce)
	return &ce
}

func (s *ChipScene) MakeCirc(image *ebiten.Image, x, y float64, mass float64) *ChipEnt {
	ce := ChipEnt{}
	ce.Sprite = image

	ow, oh := ce.Sprite.Size()
	if ow != oh {
		hiderr.Msg("not a square sprite")
	}

	ce.Width = float64(ow)
	ce.Height = float64(oh)
	r := ce.Width / 2

	moment := cp.MomentForCircle(mass, 0, r, cp.Vector{})
	body := cp.NewBody(mass, moment)
	body.SetPosition(cp.Vector{x, y})
	ce.Body = body

	shape := cp.NewCircle(body, r, cp.Vector{0, 0})
	shape.SetElasticity(0.7)
	ce.Shape = shape

	s.Ents = append(s.Ents, &ce)
	return &ce
}

func (s *ChipScene) Step() {
	s.Space.Step(1.0 / float64(ebiten.MaxTPS()))
}

func (s *ChipScene) Keys() {}

func (s *ChipScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	op := &ebiten.DrawImageOptions{}

	for _, ent := range s.Ents {
		if ent.Space == nil {
			ent.Space = s.Space
			ent.Space.AddBody(ent.Body)
			ent.Space.AddShape(ent.Shape)
		}
		op.GeoM.Reset()
		op.GeoM.Translate(ent.Body.Position().X-ent.Width/2, ent.Body.Position().Y-ent.Height/2)
		screen.DrawImage(ent.Sprite, op)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (s *ChipScene) Update(screen *ebiten.Image) {
	s.Step()
	s.Keys()
	if !ebiten.IsDrawingSkipped() {
		s.Draw(screen)
	}
}
