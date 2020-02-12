package util

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var Letters []string = []string{
	"abcdefghi",
	"jklmnopqr",
	"stuvwxyz ",
}

type TextEntry struct {
	Prompt  string
	Text    string
	SelectX int
	SelectY int
}

func (t *TextEntry) Display(screen *ebiten.Image) error {
	disp := t.Prompt
	disp += ": "
	disp += t.Text
	disp += "\n\n"
	for y, row := range Letters {
		for x, r := range row {
			if x == t.SelectX && y == t.SelectY {
				disp += "["
				disp += string(r)
				disp += "]"
			} else {
				disp += " "
				disp += string(r)
				disp += " "
			}
		}
		disp += "\n"
	}

	return ebitenutil.DebugPrint(screen, disp)
}
