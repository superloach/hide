package easy

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/superloach/hide/state"
)

type Menu struct {
	Opts    []*ebiten.Image
	Chosen  bool
	Select  int
}

func (m *Menu) Display(screen *ebiten.Image, s *state.State) error {
	if len(m.Opts) > 0 {
		err := screen.DrawImage(m.Opts[m.Select], nil)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Menu) DoKeys(s *state.State) error {
	if RepeatKey(s.Keys.Up) {
		m.Select--
	}

	if RepeatKey(s.Keys.Down) {
		m.Select++
	}

	m.Select += len(m.Opts)
	m.Select %= len(m.Opts)

	if RepeatKey(s.Keys.Confirm) {
		m.Chosen = true
	}

	return nil
}
