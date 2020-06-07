package level

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/superloach/hide/state"
)

type Level struct {
	Name       string
	Tiles      map[string]string
	TileImages map[string]*ebiten.Image
}

func (l *Level) Display(screen *ebiten.Image, s *state.State) error {
	return nil
}
