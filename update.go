package hide

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/superloach/hide/util"
)

func (h *Hide) Update(screen *ebiten.Image) error {
	h.Frame++
	h.Frame %= 60

	switch h.Scene {
	case "init":
		h.Scene = "loading"
		go func() {
			h.Menu = &util.Menu{
				Title: "HIDE:dev",
				Opts: []string{
					"new",
					"thingy",
					"exit",
				},
			}
			time.Sleep(time.Second)
			h.Scene = "menu1"
		}()
		return nil
	case "loading":
		n := h.Frame / 15
		text := "loading"
		for i := 0; i < n; i++ {
			text += "."
		}
		ebitenutil.DebugPrint(screen, text)
		return nil
	case "menu1":
		err := h.Menu.Keys(h.Keys)
		if err != nil {
			return err
		}

		err = h.Menu.Display(screen)
		if err != nil {
			return err
		}

		switch h.Menu.Chosen {
		case "":
			return nil
		case "exit":
			return util.Exit
		case "new":
			h.Scene = "init_new"
			return nil
		case "thingy":
			h.Scene = "thingy"
			return nil
		default:
			return nil
		}
	case "init_new":
		h.Scene = "loading"
		go func() {
			h.TextEntry = &util.TextEntry{
				Prompt: "enter your name: ",
			}
			time.Sleep(time.Second)
			h.Scene = "new"
		}()
		return nil
	case "new":
		//		err := h.TextEntry.Keys(h.Keys)
		//		if err != nil {
		//			return err
		//		}

		return h.TextEntry.Display(screen)
	default:
		return fmt.Errorf("unknown scene %s", h.Scene)
	}
}
