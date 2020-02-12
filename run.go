package hide

import (
	"github.com/hajimehoshi/ebiten"
)

func (h *Hide) Run() error {
	//	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("HIDE:dev")
	return ebiten.RunGame(h)
}
