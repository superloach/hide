package hide

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/superloach/hide/util"
)

type Hide struct {
	Scene     string
	Frame     int
	Keys      map[string]ebiten.Key
	Menu      *util.Menu
	TextEntry *util.TextEntry
}
