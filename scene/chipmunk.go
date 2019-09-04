package scene

import "fmt"
import "image/color"

import "github.com/hajimehoshi/ebiten"
import "github.com/hajimehoshi/ebiten/ebitenutil"
import "github.com/hajimehoshi/ebiten/inpututil"
import "github.com/jakecoffman/cp"

// chipmunk physics scene
type ChipmunkScene struct {
	Scene
	Game    GameI
	Space   *cp.Space
	Sprites []*ebiten.Image
}

func MakeChipmunkScene(game GameI, text string) (*ChipmunkScene, int) {
	s := ChipmunkScene{}
	s.Game = game
	id := s.Game.AddScene(&s)
	return &s, id
}

func (s *ChipmunkScene) Step() {
	if s.Space == nil {
		s.Space = cp.NewSpace()
		s.Space.Iterations = 1
	}
	space.Step(1.0 / float64(ebiten.MaxTPS()))
}

func (s *ChipmunkScene) Keys() {}

func (s *ChipmunkScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	ebitenutil.DebugPrint(screen, s.Text)

	
	space.EachBody(func(body *cp.Body) {
		op.GeoM.Reset()
		op.GeoM.Translate(body.Position().X, body.Position().Y)
		screen.DrawImage(dot, op)
	})

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (s *ChipmunkScene) Update(screen *ebiten.Image) {
	s.Step()
	s.Keys()
	if !ebiten.IsDrawingSkipped() {
		s.Draw(screen)
	}
}

func AddRect(sprite *ebiten.Image, x, y float64) (*cp.Body, *cp.Shape) {
	body := cp.NewBody(1.0, cp.INFINITY)
	body.SetPosition(cp.Vector{x, y})

	shape := cp.NewCircle(body, 0.95, cp.Vector{})
	shape.SetElasticity(0)
	shape.SetFriction(0)

	return body, shape
}
