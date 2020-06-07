package hide

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/superloach/hide/easy"
)

func (h *Hide) Loading() {
	h.State.Frame = 0
	h.CurScene = "loading"
}

func (h *Hide) Update(screen *ebiten.Image) error {
	h.State.Frame++
	h.State.Frame %= 60

	switch h.CurScene {
	case "loading":
		n := h.State.Frame / 15
		text := "loading"
		for i := 0; i < n; i++ {
			text += "."
		}
		ebitenutil.DebugPrint(screen, text)
	case "init":
		h.Loading()
		go func() {
			img, err := easy.LoadImage("splash")
			if err != nil {
				panic(err)
			}

			ks := h.State.Keys
			caption := "\n controls:\n"
			caption += "      up - [" + ks.Up.String() + "]\n"
			caption += "    down - [" + ks.Down.String() + "]\n"
			caption += "    left - [" + ks.Left.String() + "]\n"
			caption += "   right - [" + ks.Right.String() + "]\n"
			caption += " confirm - [" + ks.Confirm.String() + "]\n"
			caption += "    back - [" + ks.Back.String() + "]\n"
			caption += "   pause - [" + ks.Pause.String() + "]\n\n"
			caption += "(press any key)"

			h.Image = &easy.Imager{
				Image:   img,
				Caption: caption,
			}

			h.CurScene = "splash"

			time.Sleep(time.Second * 5)

			h.Image.Next = true
		}()
	case "splash":
		err := h.Image.Display(screen, h.State)
		if err != nil {
			return err
		}

		err = h.Image.DoKeys(h.State)
		if err != nil {
			return err
		}

		if h.Image.Next {
			h.CurScene = "init_menu"
		}
	case "init_menu":
		h.Loading()

		go func() {
			opts := make([]*ebiten.Image, 0)
			for _, opt := range []string{"menu_load", "menu_new", "menu_extras", "menu_quit"} {
				img, err := easy.LoadImage(opt)
				if err != nil {
					panic(err)
				}
				opts = append(opts, img)
			}

			h.Menu = &easy.Menu{}
			h.Menu.Opts = opts

			h.CurScene = "menu"
		}()
	case "menu":
		err := h.Menu.DoKeys(h.State)
		if err != nil {
			return err
		}

		err = h.Menu.Display(screen, h.State)
		if err != nil {
			return err
		}

		if h.Menu.Chosen {
			switch h.Menu.Select {
			case 0:
				if !h.State.Exists() {
					h.CurScene = "no_save"
				} else {
					h.CurScene = h.State.Scene
				}
			case 1:
				h.CurScene = "init_new"
			case 2:
				h.CurScene = "init_extras"
			case 3:
				return easy.Exit
			default:
			}
		}
	case "init_no_save":
		h.Loading()

		go func() {
			h.Image.Image = nil
			h.Image.Caption = "no save\n\n(press any key)"
			h.Image.Next = false

			h.CurScene = "no_save"
		}()
	case "no_save":
		err := h.Image.Display(screen, h.State)
		if err != nil {
			return err
		}

		err = h.Image.DoKeys(h.State)
		if err != nil {
			return err
		}

		if h.Image.Next {
			h.CurScene = "init_menu"
		}
	case "init_new":
		h.Loading()
		go func() {
			h.Texter = &easy.Texter{
				Prompt: "enter your name",
			}

			h.CurScene = "new"
		}()
	case "new":
		err := h.Texter.DoKeys(h.State)
		if err != nil {
			return err
		}

		err = h.Texter.Display(screen, h.State)
		if err != nil {
			return err
		}

		text := h.Texter.Text
		if len(text) > 0 && text[len(text)-1] == ' ' {
			if len(text) == 1 {
				h.Texter.Text = ""
			} else {
				h.CurScene = "init_begin"
			}
		}
	case "init_begin":
		h.Loading()
		go func() {
			text := h.Texter.Text

			h.State.Reset()
			h.State.Name = text[:len(text)-1]

			h.CurScene = "begin"
			h.State.Scene = h.CurScene

			h.State.Save()
		}()
	case "begin":
		text := "begin STUB\nname: " + h.State.Name + "\nscene: " + h.State.Scene

		err := ebitenutil.DebugPrint(screen, text)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown scene %s", h.CurScene)
	}

	return nil
}
