package state

import "github.com/hajimehoshi/ebiten"

var BlankState *State

func init() {
	BlankState = &State{
		Keys: &Keys{
			Up:      ebiten.KeyUp,
			Down:    ebiten.KeyDown,
			Left:    ebiten.KeyLeft,
			Right:   ebiten.KeyRight,
			Confirm: ebiten.KeyZ,
			Back:    ebiten.KeyX,
			Pause:   ebiten.KeyC,
		},
	}
}
