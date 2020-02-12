package util

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Menu struct {
	Title  string
	Opts   []string
	Chosen string
	Select int
}

func (m *Menu) Display(screen *ebiten.Image) error {
	disp := m.Title + "\n"
	for i, o := range m.Opts {
		disp += " "
		if i == m.Select {
			disp += ">"
		} else {
			disp += " "
		}
		disp += " " + o + "\n"
	}

	return ebitenutil.DebugPrint(screen, disp)
}

func (m *Menu) Keys(keys map[string]ebiten.Key) error {
	for _, k := range []string{"up", "down", "confirm"} {
		_, ok := keys[k]
		if !ok {
			return fmt.Errorf("missing key %s", k)
		}
	}

	if RepeatKey(keys["up"]) {
		m.Select--
	}

	if RepeatKey(keys["down"]) {
		m.Select++
	}

	m.Select += len(m.Opts)
	m.Select %= len(m.Opts)

	if RepeatKey(keys["confirm"]) {
		m.Chosen = m.Opts[m.Select]
	}

	return nil
}
