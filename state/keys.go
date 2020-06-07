package state

import "github.com/hajimehoshi/ebiten"

type Keys struct {
	Up      ebiten.Key
	Down    ebiten.Key
	Left    ebiten.Key
	Right   ebiten.Key
	Confirm ebiten.Key
	Back    ebiten.Key
	Pause   ebiten.Key
}

func (k *Keys) All() []ebiten.Key {
	return []ebiten.Key{
		k.Up,
		k.Down,
		k.Left,
		k.Right,
		k.Confirm,
		k.Back,
		k.Pause,
	}
}
