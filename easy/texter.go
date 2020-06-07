package easy

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/superloach/hide/state"
)

var LetterSets [][]string = [][]string{
	{
		"abcdefghi",
		"jklmnopqr",
		"stuvwxyz ",
	},
	{
		"ABCDEFGHI",
		"JKLMNOPQR",
		"STUVWXYZ ",
	},
}

type Texter struct {
	Prompt    string
	Text      string
	SelectX   int
	SelectY   int
	LetterSet int
}

func (t *Texter) Display(screen *ebiten.Image, s *state.State) error {
	letters := LetterSets[t.LetterSet]
	disp := t.Prompt + ": " + t.Text
	if s.Frame > 30 {
		disp += "_"
	}
	disp += "\n\n"
	for y, row := range letters {
		for x, r := range row {
			d := string(r)
			if d == " " {
				d = "begin"
			}
			if x == t.SelectX && y == t.SelectY {
				disp += "[" + d + "]"
			} else {
				disp += " " + d + " "
			}
		}
		disp += "\n"
	}

	return ebitenutil.DebugPrint(screen, disp)
}

func (t *Texter) DoKeys(s *state.State) error {
	letters := LetterSets[t.LetterSet]
	if RepeatKey(s.Keys.Up) {
		t.SelectY--
	}

	if RepeatKey(s.Keys.Down) {
		t.SelectY++
	}

	if RepeatKey(s.Keys.Left) {
		t.SelectX--
	}

	if RepeatKey(s.Keys.Right) {
		t.SelectX++
	}

	t.SelectY += len(letters)
	t.SelectY %= len(letters)

	t.SelectX += len(letters[t.SelectY])
	t.SelectX %= len(letters[t.SelectY])

	if RepeatKey(s.Keys.Confirm) {
		t.Text += string([]byte{letters[t.SelectY][t.SelectX]})
	}

	if RepeatKey(s.Keys.Back) && len(t.Text) > 0 {
		t.Text = t.Text[:len(t.Text)-1]
	}

	return nil
}
