package hide

import (
	"github.com/superloach/hide/easy"
	"github.com/superloach/hide/level"
	"github.com/superloach/hide/state"
)

type Hide struct {
	CurScene string
	State    *state.State
	Image    *easy.Imager
	Menu     *easy.Menu
	Texter   *easy.Texter
	Level    *level.Level
}

func MkHide() (*Hide, error) {
	h := Hide{}

	h.CurScene = "init"
	h.State = state.Load()

	return &h, nil
}
