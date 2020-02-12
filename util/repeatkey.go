package util

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

var (
	RepeatDelay    int = 30
	RepeatInterval int = 5
)

func RepeatKey(key ebiten.Key) bool {
	d := inpututil.KeyPressDuration(key)
	return d == 1 || (d >= RepeatDelay && (d-RepeatDelay)%RepeatInterval == 0)
}
