package easy

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/superloach/hide/state"
)

type Imager struct {
	Image   *ebiten.Image
	Caption string
	Next    bool
}

func (i *Imager) Display(screen *ebiten.Image, s *state.State) error {
	if i.Image != nil {
		opts := &ebiten.DrawImageOptions{}
		err := screen.DrawImage(i.Image, opts)
		if err != nil {
			return err
		}
	}
	return ebitenutil.DebugPrint(screen, i.Caption)
}

func (i *Imager) DoKeys(s *state.State) error {
	for _, k := range s.Keys.All() {
		if RepeatKey(k) {
			i.Next = true
		}
	}

	return nil
}
