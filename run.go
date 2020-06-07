package hide

import (
	"github.com/hajimehoshi/ebiten"
)

func (h *Hide) Run() error {
	ebiten.SetWindowSize(Width*Scale, Height*Scale)
	ebiten.SetWindowTitle("HIDE:dev")
	return ebiten.RunGame(h)
}
