package hide

import "github.com/hajimehoshi/ebiten"

func MkHide() *Hide {
	h := Hide{}

	h.Scene = "init"
	h.Frame = 0
	h.Keys = map[string]ebiten.Key{
		"up":      ebiten.KeyUp,
		"down":    ebiten.KeyDown,
		"left":    ebiten.KeyLeft,
		"right":   ebiten.KeyRight,
		"confirm": ebiten.KeyZ,
		"back":    ebiten.KeyX,
		"pause":   ebiten.KeyC,
	}

	return &h
}
